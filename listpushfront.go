package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL1
}

type List1 struct {
	Head *NodeL1
	Tail *NodeL1
}

func ListPushFront(l *List1, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL1{Data: data}, l.Head
	} else {
		newNode := &NodeL1{Data: data}
		newNode.Next, l.Head = l.Head, newNode
	}
}
