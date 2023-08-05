package main

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList(head *Node) *LinkedList {
	ll := &LinkedList{
		Head:   head,
		Tail:   head,
		Length: 1,
	}
	return ll
}

func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length += 1
}

func (l *LinkedList) Prepend(value interface{}) *LinkedList {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Length += 1

	return l
}

func (l *LinkedList) Insert(index int, value interface{}) *LinkedList {
	if index >= l.Length {
		return l.Prepend(value)
	}

	newNode := &Node{
		Value: value,
		Next:  nil,
	}
	leader := l.traverseToIndex(index - 1)
	holdingPointer := leader.Next
	leader.Next = newNode
	newNode.Next = holdingPointer
	l.Length++

	return nil
}

func (l *LinkedList) traverseToIndex(index int) *Node {
	counter := 0
	currentNode := l.Head
	for counter < index {
		currentNode = currentNode.Next
		counter++
	}

	return currentNode
}
