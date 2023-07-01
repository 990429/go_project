package main

import (
	"fmt"
	"time"
)

var (
	numChan  chan int
	resChan  chan [2]int
	exitChan chan bool
)

func store(numChan chan int) {
	for i := 1; i <= 800000; i++ {
		numChan <- i
	}
	close(numChan)
}
func sub_thread(numChan chan int) {
	for {
		if n, ok := <-numChan; ok {
			resChan <- [2]int{n, n * (n + 1) / 2}
		} else {
			fmt.Println("一个线程结束")
			exitChan <- true
			break
		}
	}
}

func main() {
	numChan = make(chan int, 1000)
	resChan = make(chan [2]int, 1000)
	exitChan = make(chan bool, 100)
	go store(numChan)
	time.Sleep(time.Millisecond)
	for i := 1; i <= 100; i++ {
		go sub_thread(numChan)
	}
	go func() {
		for i := 1; i <= 100; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	for {
		if n, ok := <-resChan; ok {
			fmt.Printf("res[%d]=%d\n", n[0], n[1])
		} else {
			fmt.Println("全部结束！")
			//close(resChan)
			break
		}
	}
}
