package dsa

import "fmt"

type Nodes struct {
	Value int
}

type Stack struct {
	nodes []*Nodes
	count int
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
    nodes []*Nodes
    size  int
    head  int
    tail  int
    count int
}

func (n *Nodes) String() string {
	return fmt.Sprint(n.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Nodes) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Nodes {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
func NewQueue(size int) *Queue {
	return &Queue{
		nodes []*Nodes
		size int
		head int
		tail int
		count int
	}
}

