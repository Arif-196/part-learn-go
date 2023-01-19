package dsa

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  int64
}

type BinaryTree struct {
	Root *BinaryNode
}

func (t *BinaryTree) Insert(Data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Data: Data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(Data)
	}
	return t
}

func (n *BinaryNode) Insert(Data int64) {
	if n == nil {
		return
	} else if Data <= n.Data {
		if n.Left == nil {
			n.Left = &BinaryNode{Data: Data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(Data)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Data: Data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(Data)
		}
	}
}

func Print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	Print(w, node.Left, ns+2, 'L')
	Print(w, node.Right, ns+2, 'R')

}
