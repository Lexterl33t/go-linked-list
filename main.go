package linked

import (
	"errors"
	"fmt"
)

//////////////////////
// 	  INIT TYPE		 //
//////////////////////

type Node struct {
	Value any
	Index int
	Next  *Node
	Prev  *Node
}

type List struct {
	Head *Node
}

//////////////////////
// 	  INIT List		//
//////////////////////

func NewList(defaultValue any) *List {
	var node *Node = &Node{
		Value: defaultValue,
		Index: 0,
		Next:  nil,
		Prev:  nil,
	}

	var list *List = &List{
		Head: node,
	}

	return list
}

//////////////////////
// 	  INSERT 		//
//////////////////////

func (list *List) Append(value any) *Node {
	var n_node *Node = &Node{
		Value: value,
		Index: list.Head.Index,
		Next:  list.Head,
		Prev:  nil,
	}

	list.Head.Prev = n_node
	list.Head = n_node
	list.Head.Index++
	return n_node
}

func (list *List) AppendAt(n, value any) (*Node, error) {
	var chunk_head *Node = list.Head
	for {
		if chunk_head == nil {
			return nil, errors.New("Index not found")
		}
		// not working...
	}
}

func (head *List) Show() {
	var chunk_head *Node = head.Head
	for {
		if chunk_head == nil {
			break
		}

		fmt.Println(chunk_head.Value, chunk_head.Index)

		chunk_head = chunk_head.Next
	}
}

func (head *List) Get(index int) (*Node, error) {
	var chunk_head *Node = head.Head
	for {
		if chunk_head == nil {
			return nil, errors.New("Index not found")
		}

		if index == -1 {
			if chunk_head.Next == nil {
				return chunk_head, nil
			}
		} else if index < 0 && index != -1 {
			index *= -1
		}

		if chunk_head.Index == index {
			return chunk_head, nil
		}

		chunk_head = chunk_head.Next
	}
}

func (head *List) Pop() {

	var chunk_head *Node = head.Head
	for {
		if chunk_head == nil {
			return
		}

		if chunk_head.Prev != nil {
			chunk_head.Prev.Next.Index--
		} else {
			chunk_head.Index--
		}

		if chunk_head.Next.Next == nil {
			chunk_head.Prev.Next = chunk_head.Next
			chunk_head.Next = nil
			return
		}

		chunk_head = chunk_head.Next
	}
}

func (head *List) Delete(n int) error {
	var chunk_head *Node = head.Head

	for {

		if chunk_head == nil {
			return errors.New("index not found")
		}

		_, err := head.Get(n)
		if err != nil {
			return errors.New("index not found")
		}

		if chunk_head.Prev != nil {
			chunk_head.Prev.Next.Index--
		} else {
			chunk_head.Index--
		}

		if chunk_head.Index == n {
			chunk_head.Prev.Next = chunk_head.Next
			chunk_head.Next = nil
			return nil
		}

		chunk_head = chunk_head.Next
	}
}
