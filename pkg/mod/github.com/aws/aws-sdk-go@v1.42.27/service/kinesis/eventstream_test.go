// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.16
// +build go1.16

package kinesis

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamtest"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

var _ time.Time
var _ awserr.Error
var _ context.Context
var _ sync.WaitGroup
var _ strings.Reader

func TestSubscribeToShard_Read(t *testing.T) {
	expectEvents, eventMsgs := mockSubscribeToShardReadEvents()
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.SubscribeToShard(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}
	defer resp.GetStream().Close()
	// Trim off response output type pseudo event so only event messages remain.
	expectEvents = expectEvents[1:]

	var i int
	for event := range resp.GetStream().Events() {
		if event == nil {
			t.Errorf("%d, expect event, got nil", i)
		}
		if e, a := expectEvents[i], event; !reflect.DeepEqual(e, a) {
			t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
		}
		i++
	}

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestSubscribeToShard_ReadClose(t *testing.T) {
	_, eventMsgs := mockSubscribeToShardReadEvents()
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.SubscribeToShard(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	// Assert calling Err before close does not close the stream.
	resp.GetStream().Err()
	select {
	case _, ok := <-resp.GetStream().Events():
		if !ok {
			t.Fatalf("expect stream not to be closed, but was")
		}
	default:
	}

	resp.GetStream().Close()
	<-resp.GetStream().Events()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestSubscribeToShard_ReadUnknownEvent(t *testing.T) {
	expectEvents, eventMsgs := mockSubscribeToShardReadEvents()
	eventOffset := 1

	unknownEvent := eventstream.Message{
		Headers: eventstream.Headers{
			eventstreamtest.EventMessageTypeHeader,
			{
				Name:  eventstreamapi.EventTypeHeader,
				Value: eventstream.StringValue("UnknownEventName"),
			},
		},
		Payload: []byte("some unknown event"),
	}

	eventMsgs = append(eventMsgs[:eventOffset],
		append([]eventstream.Message{unknownEvent}, eventMsgs[eventOffset:]...)...)

	expectEvents = append(expectEvents[:eventOffset],
		append([]SubscribeToShardEventStreamEvent{
			&SubscribeToShardEventStreamUnknownEvent{
				Type:    "UnknownEventName",
				Message: unknownEvent,
			},
		},
			expectEvents[eventOffset:]...)...)

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.SubscribeToShard(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}
	defer resp.GetStream().Close()
	// Trim off response output type pseudo event so only event messages remain.
	expectEvents = expectEvents[1:]

	var i int
	for event := range resp.GetStream().Events() {
		if event == nil {
			t.Errorf("%d, expect event, got nil", i)
		}
		if e, a := expectEvents[i], event; !reflect.DeepEqual(e, a) {
			t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
		}
		i++
	}

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func BenchmarkSubscribeToShard_Read(b *testing.B) {
	_, eventMsgs := mockSubscribeToShardReadEvents()
	var buf bytes.Buffer
	encoder := eventstream.NewEncoder(&buf)
	for _, msg := range eventMsgs {
		if err := encoder.Encode(msg); err != nil {
			b.Fatalf("failed to encode message, %v", err)
		}
	}
	stream := &loopReader{source: bytes.NewReader(buf.Bytes())}

	sess := unit.Session
	svc := New(sess, &aws.Config{
		Endpoint:               aws.String("https://example.com"),
		DisableParamValidation: aws.Bool(true),
	})
	svc.Handlers.Send.Swap(corehandlers.SendHandler.Name,
		request.NamedHandler{Name: "mockSend",
			Fn: func(r *request.Request) {
				r.HTTPResponse = &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(stream),
				}
			},
		},
	)

	resp, err := svc.SubscribeToShard(nil)
	if err != nil {
		b.Fatalf("failed to create request, %v", err)
	}
	defer resp.GetStream().Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = resp.GetStream().Err(); err != nil {
			b.Fatalf("expect no error, got %v", err)
		}
		event := <-resp.GetStream().Events()
		if event == nil {
			b.Fatalf("expect event, got nil, %v, %d", resp.GetStream().Err(), i)
		}
	}
}

func mockSubscribeToShardReadEvents() (
	[]SubscribeToShardEventStreamEvent,
	[]eventstream.Message,
) {
	expectEvents := []SubscribeToShardEventStreamEvent{
		&SubscribeToShardOutput{},
		&SubscribeToShardEvent{
			ChildShards: []*ChildShard{
				{
					HashKeyRange: &HashKeyRange{
						EndingHashKey:   aws.String("string value goes here"),
						StartingHashKey: aws.String("string value goes here"),
					},
					ParentShards: []*string{
						aws.String("string value goes here"),
						aws.String("string value goes here"),
						aws.String("string value goes here"),
					},
					ShardId: aws.String("string value goes here"),
				},
				{
					HashKeyRange: &HashKeyRange{
						EndingHashKey:   aws.String("string value goes here"),
						StartingHashKey: aws.String("string value goes here"),
					},
					ParentShards: []*string{
						aws.String("string value goes here"),
						aws.String("string value goes here"),
						aws.String("string value goes here"),
					},
					ShardId: aws.String("string value goes here"),
				},
				{
					HashKeyRange: &HashKeyRange{
						EndingHashKey:   aws.String("string value goes here"),
						StartingHashKey: aws.String("string value goes here"),
					},
					ParentShards: []*string{
						aws.String("string value goes here"),
						aws.String("string value goes here"),
						aws.String("string value goes here"),
					},
					ShardId: aws.String("string value goes here"),
				},
			},
			ContinuationSequenceNumber: aws.String("string value goes here"),
			MillisBehindLatest:         aws.Int64(1234),
			Records: []*Record{
				{
					ApproximateArrivalTimestamp: aws.Time(time.Unix(1396594860, 0).UTC()),
					Data:                        []byte("blob value goes here"),
					EncryptionType:              aws.String("string value goes here"),
					PartitionKey:                aws.String("string value goes here"),
					SequenceNumber:              aws.String("string value goes here"),
				},
				{
					ApproximateArrivalTimestamp: aws.Time(time.Unix(1396594860, 0).UTC()),
					Data:                        []byte("blob value goes here"),
					EncryptionType:              aws.String("string value goes here"),
					PartitionKey:                aws.String("string value goes here"),
					SequenceNumber:              aws.String("string value goes here"),
				},
				{
					ApproximateArrivalTimestamp: aws.Time(time.Unix(1396594860, 0).UTC()),
					Data:                        []byte("blob value goes here"),
					EncryptionType:              aws.String("string value goes here"),
					PartitionKey:                aws.String("string value goes here"),
					SequenceNumber:              aws.String("string value goes here"),
				},
			},
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(jsonrpc.BuildHandler)
	payloadMarshaler := protocol.HandlerPayloadMarshal{
		Marshalers: marshalers,
	}
	_ = payloadMarshaler

	eventMsgs := []eventstream.Message{
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("initial-response"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[0]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("SubscribeToShardEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[1]),
		},
	}

	return expectEvents, eventMsgs
}
func TestSubscribeToShard_ReadException(t *testing.T) {
	expectEvents := []SubscribeToShardEventStreamEvent{
		&SubscribeToShardOutput{},
		&InternalFailureException{
			RespMetadata: protocol.ResponseMetadata{
				StatusCode: 200,
			},
			Message_: aws.String("string value goes here"),
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(jsonrpc.BuildHandler)
	payloadMarshaler := protocol.HandlerPayloadMarshal{
		Marshalers: marshalers,
	}

	eventMsgs := []eventstream.Message{
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("initial-response"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[0]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventExceptionTypeHeader,
				{
					Name:  eventstreamapi.ExceptionTypeHeader,
					Value: eventstream.StringValue("InternalFailureException"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[1]),
		},
	}

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.SubscribeToShard(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	defer resp.GetStream().Close()

	<-resp.GetStream().Events()

	err = resp.GetStream().Err()
	if err == nil {
		t.Fatalf("expect err, got none")
	}

	expectErr := &InternalFailureException{
		RespMetadata: protocol.ResponseMetadata{
			StatusCode: 200,
		},
		Message_: aws.String("string value goes here"),
	}
	aerr, ok := err.(awserr.Error)
	if !ok {
		t.Errorf("expect exception, got %T, %#v", err, err)
	}
	if e, a := expectErr.Code(), aerr.Code(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectErr.Message(), aerr.Message(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	if e, a := expectErr, aerr; !reflect.DeepEqual(e, a) {
		t.Errorf("expect error %+#v, got %+#v", e, a)
	}
}

var _ awserr.Error = (*InternalFailureException)(nil)
var _ awserr.Error = (*KMSAccessDeniedException)(nil)
var _ awserr.Error = (*KMSDisabledException)(nil)
var _ awserr.Error = (*KMSInvalidStateException)(nil)
var _ awserr.Error = (*KMSNotFoundException)(nil)
var _ awserr.Error = (*KMSOptInRequired)(nil)
var _ awserr.Error = (*KMSThrottlingException)(nil)
var _ awserr.Error = (*ResourceInUseException)(nil)
var _ awserr.Error = (*ResourceNotFoundException)(nil)

type loopReader struct {
	source *bytes.Reader
}

func (c *loopReader) Read(p []byte) (int, error) {
	if c.source.Len() == 0 {
		c.source.Seek(0, 0)
	}

	return c.source.Read(p)
}
