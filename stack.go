package main

import (
	"fmt"
)

type (
	SymbolNode struct {
		Symbol rune
		Right  *SymbolNode
		Left   *SymbolNode
	}

	SymbolStack struct {
		List       *SymbolNode
		Head       *SymbolNode
		noElements uint
		State      string
	}
)

func NewSymboStack() *SymbolStack {
	return &SymbolStack{}
}

func (s *SymbolStack) Add(symbol rune) {
	if s.List == nil {
		s.List = &SymbolNode{
			Symbol: symbol,
		}
		s.Head = s.List
		s.noElements++
		return
	}

	// nos colocamos al final de la lista
	lastNode := s.List
	for lastNode.Right != nil {
		lastNode = lastNode.Right
	}
	lastNode.Right = &SymbolNode{
		Symbol: symbol,
	}
	lastNode.Right.Left = lastNode
	s.noElements++
}

func (s *SymbolStack) Last() *SymbolNode {
	nav := s.List
	for nav.Right != nil {
		nav = nav.Right
	}
	return nav
}

func (s *SymbolStack) Remove(node *SymbolNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		// se desplaza uno a la izuierda
		node.Left.Right = node.Right
	}

	if node.Right != nil {
		node.Right.Left = node.Left
	}
	node.Left = nil
	node.Right = nil
	node = nil
	s.noElements--
}

func (node *SymbolNode) JoinWith(toJoin *SymbolStack) {
	originalRight := node.Right
	node.Right = toJoin.Head
	toJoin.Head.Left = node
	tail := toJoin.Last()
	tail.Right = originalRight
}

func (s *SymbolStack) insertString(str string) {
	if s.Head == nil {
		return
	}
	// el nodo actual se desecha
	leftNode := s.Head.Left
	s.Remove(s.Head)
	toInsertList := NewSymboStack()
	toInsertList.AddFromString(str)
	// si tenemos nodo izquierda
	if leftNode != nil {
		leftNode.JoinWith(toInsertList)
	} else {
		s.Head = toInsertList.Head
	}
}

func (s *SymbolStack) AddFromString(str string) {
	for _, char := range str {
		s.Add(char)
	}
}

func (s *SymbolStack) Debug() {
	nav := s.List
	for nav != nil {
		fmt.Printf("%c -> %p\n", nav.Symbol, nav.Left)
		nav = nav.Right
	}
}

func (s *SymbolStack) String() string {
	str := ""
	nav := s.List
	for nav != nil {
		str += string(nav.Symbol)
		nav = nav.Right
	}
	return str
}

func (s *SymbolStack) MoveRigth() {
	s.Head = s.Head.Right
}

func (s *SymbolNode) String() string {
	return string(s.Symbol)
}
