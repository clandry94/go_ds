package queue

import (
	"net"
)

type Conn struct {
	Connection net.Conn
	Id         int
}

type Node struct {
	Next  *Node
	Value interface{}
}

type Queue struct {
	Front *Node
	Back  *Node
	Size  int
}

func New() Queue {
	return Queue{Size: 0}
}

func (q *Queue) Push(item Node) {
	if q.Size == 0 {
		q.Front = &item
		q.Back = &item
	} else {
		q.Back.Next = &item
		q.Back = &item
	}
	q.Size += 1
}

func (q *Queue) Pop() {
	q.Front = q.Front.Next
	q.Size -= 1
}
