package main

//design
type LRUCache struct {
	capacity   int
	store      map[int][2]*Listnode //pos1:pre list,pos2:cur list
	Queue_head *Listnode
	Queue_tail *Listnode
	Queue_len  int
}

///使用链表来实现key的删除添加，使查找复杂度为O（1）
type Listnode struct {
	Key   int
	Value int
	Next  *Listnode
}

func Constructor(capacity int) LRUCache {
	var head *Listnode
	head = new(Listnode)
	LRU := LRUCache{capacity: capacity, Queue_head: head, Queue_len: 0, Queue_tail: head}
	LRU.store = make(map[int][2]*Listnode)

	return LRU
}

func (this *LRUCache) Get(key int) int {
	if List, ok := this.store[key]; ok {
		if List[1] == this.Queue_tail {
			return List[1].Value
		}
		List[0].Next = List[1].Next //把当前key从队列中拿出来
		this.store[key] = [2]*Listnode{this.Queue_tail, List[1]}
		//放入队尾，更新队尾
		this.Queue_tail.Next = List[1]
		this.Queue_tail = List[1]
		List[1].Next = nil
		return List[1].Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.store[key]; ok {
		this.Get(key)
		this.store[key][1].Value = value
	} else {
		if this.Queue_len == this.capacity {

			temp := this.Queue_head.Next
			this.Queue_head.Next = this.Queue_head.Next.Next //error
			if this.Queue_head.Next != nil {
				this.store[this.Queue_head.Next.Key] = [2]*Listnode{this.Queue_head, this.Queue_head.Next}
			}
			delete(this.store, temp.Key)
			this.Queue_len--
			if this.capacity == 1 {
				this.Queue_tail = this.Queue_head
			}
			//将最近最少使用的删除
		}
		newNode := &Listnode{Key: key, Value: value, Next: nil}

		this.store[key] = [2]*Listnode{this.Queue_tail, newNode}
		this.Queue_tail.Next = newNode
		this.Queue_tail = this.Queue_tail.Next
		this.Queue_len++
	}
}
