package linkedlist

import "fmt"

type Nodes struct {
	next 	*Nodes
	data	interface{}
}

type LinkedLists struct {
	head 	*Nodes
	length	int
}

func (l *LinkedLists) Prepend(value interface{}) {
	newNode := Nodes{data: value}
	
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	
}

func (l *LinkedLists) PrintLists() {
	if l.head == nil {
		return
	} 
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedLists) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
	
}