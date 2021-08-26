package linkedlist

import "fmt"

type Node struct{
	value	interface{}
	next	*Node
}

type LinkedList struct {
	head	*Node
	length	int
}

func (l *LinkedList) PrintList() {
	if l.length == 0 {
		fmt.Println("No Data in the List")
		return
	}
	currentHead := l.head
	for i:=0; i<l.length;i++ {
		fmt.Printf("%d: %v\n", i+1, currentHead.value)
		currentHead = currentHead.next
	}
}

func (l *LinkedList) Add(value interface{}) {
	newNode := Node{value: value}
	if l.head == nil {
		l.head = &newNode
		l.length++
		return
	}
	currentNode := l.head
	//for i:=0; i<l.length; i++ {
	for currentNode != nil {
		if currentNode.next == nil {
			currentNode.next = &newNode
			l.length++
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := Node{value: value}
	if l.length == 0 {
		l.head = &newNode
		l.length++
		return
	}
	newNode.next = l.head
	l.head = &newNode
	l.length++
	//l.head.next = currentHead.next
}

func (l *LinkedList) Search(val interface{}) int{
	currentNode := l.head
	for i := 0; i < l.length; i++ {
		if currentNode.value == val {
			return i+1
		}
		currentNode = currentNode.next
	}
	return -1
}

func (l *LinkedList) SearchAt(pos int) interface{}{
	currentNode := l.head

	for i:=1; i <= l.length; i++ {
		if i == pos {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return -1
}

func (l *LinkedList) DeleteAt(pos int) {
	currentNode := l.head
	if pos == 1 {
		l.head = currentNode.next
		l.length--
		return
	} else if pos > l.length {
		fmt.Println("Position Not Exists")
		return
	}
	for i := 1; i<= l.length; i++ {
		if i == pos-1 {
			currentNode.next = currentNode.next.next
			l.length--
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList) InsertAt(pos int,value interface{}){
	currentNode := l.head

	if pos == 1 {
		l.head.value = value
		return
	} else if pos > l.length || pos <= 0 {
		fmt.Println("Position Does Not Exists")
		return
	}

	for i := 1; i <= l.length; i++ {
		if i == pos {
			currentNode.value = value
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList) InsertBefore(pos int, value interface{}) {
	currentNode := l.head
	newNode := Node{
		value: value,
	}

	if pos < 1 {
		fmt.Println("Index cannot be less than 1")
		return
	} else if pos > l.length+1 {
		fmt.Println("Index Does not exist")
		return
	}

	if pos == 1 {
		newNode.next = l.head
		l.head = &newNode
		l.length++
		return
	}

	for i := 1; i <= l.length; i++ {
		if i == pos-1 {
			newNode.next = currentNode.next
			currentNode.next = &newNode
			l.length++
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList) InsertAfter(pos int, value interface{}) {
	currentNode := l.head
	newNode := Node{value: value}
	if pos < 0  {
		fmt.Println("Index cannot be less than 0")
		return
	} else if pos > l.length {
		fmt.Printf("Index Cannot be greater than %v\n", l.length)
		return
	}

	for i := 1; i <= l.length; i++ {
		if i == pos && pos < l.length {
			newNode.next = currentNode.next.next
			currentNode.next = &newNode
			return
		}
		if i == pos && pos == l.length {
			currentNode.next = &newNode
			l.length++
			return
		}
		currentNode = currentNode.next
	}
}

func SinglyListTestingCases() {
	list := LinkedList{}

	list.Add(23)
	list.Add(43)
	//
	list.PrintList()
	fmt.Println("..............Prepend........")

	list.Prepend("First Item")
	list.Prepend("A new Added Item")

	list.Add("Good to go")

	list.PrintList()

	fmt.Println("..............Search Value..........")

	//found := list.Search(23)
	found := list.Search(43)
	//found := list.Search("Good to go")
	if found >= 0 {
		fmt.Printf("The value found at position %d\n", found)
	}else {
		fmt.Printf("Not Found\n")
	}

	fmt.Printf("The Value at Position 3 and 5 is: %v, %v\n", list.SearchAt(3), list.SearchAt(5))

	fmt.Printf("%v\n",list.SearchAt(6))

	fmt.Println("...........Delete At Position.........")

	list.DeleteAt(5)
	list.PrintList()

	fmt.Println("...........Insert At Position.........")

	list.InsertAt(1, "First of its Kind")
	list.InsertAt(3, 27)
	list.PrintList()

	fmt.Println("...........Insert At Position.........")

	list.InsertAt(1, "First of its Kind")
	list.InsertAt(3, 27)
	list.PrintList()

	fmt.Println("...........Insert Before Position.........")

	list.InsertBefore(1, "First of its Kind")
	//list.InsertAt(3, 27)
	list.PrintList()

	fmt.Println("...........Insert After Position.........")

	list.InsertAfter(1, "Second of its Kind")
	list.InsertAfter(5, 42)
	list.PrintList()
}