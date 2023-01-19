package dsa

import "fmt"

type Node struct {
	Prev *Node
	Next *Node
	Key  interface{}
}

type List struct {
	Head *Node
	Tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		Next: L.Head,
		Key:  key,
	}
	if L.Head != nil {
		L.Head.Prev = list
	}
	L.Head = list

	l := L.Head
	for l.Next != nil {
		l = l.Next
	}
	L.Tail = l
}

func (l *List) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v ->", list.Key)
		list = list.Next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.Key)
		list = list.Next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.Key)
		list = list.Prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.Head
	var prev *Node
	l.Tail = l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
	Display(l.Head)
}
