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
	return &SymbolStack{
		State: "q0",
	}
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

	if node == s.List { // Si es el primer nodo
		s.List = node.Right
		if s.List != nil {
			s.List.Left = nil
		}
	} else if node.Right != nil { // Nodo intermedio
		node.Right.Left = node.Left
	}
	if node.Left != nil {
		node.Left.Right = node.Right
	}

	if node == s.Head { // Si es el nodo actual
		s.Head = node.Right
	}

	node.Left = nil
	node.Right = nil
	s.noElements--
}

func (node *SymbolNode) JoinWith(toJoin *SymbolStack) {
	originalRight := node.Right
	node.Right = toJoin.Head
	toJoin.Head.Left = node
	tail := toJoin.Last()
	tail.Right = originalRight
	if originalRight != nil {
		originalRight.Left = tail
	}
}

func (s *SymbolStack) insertString(str string) {
	if s.Head == nil {
		return
	}
	leftNode := s.Head.Left
	rightNode := s.Head.Right

	// Remueve el nodo actual
	s.Remove(s.Head)

	// Inserta la nueva sublista
	toInsertList := NewSymboStack()
	toInsertList.AddFromString(str)

	// Conecta la sublista con los nodos izquierdo y derecho
	if leftNode != nil {
		leftNode.Right = toInsertList.List
		toInsertList.List.Left = leftNode
	} else {
		s.List = toInsertList.List
	}

	tail := toInsertList.Last()
	if rightNode != nil {
		tail.Right = rightNode
		rightNode.Left = tail
	}
	s.Head = toInsertList.Head
}

func (s *SymbolStack) AddFromString(str string) {
	for _, char := range str {
		s.Add(char)
	}
}

func (s *SymbolStack) MoveRigth() {
	if s.Head != nil && s.Head.Right != nil {
		s.Head = s.Head.Right
	}
}

func (s *SymbolStack) MoveLeft() {
	if s.Head != nil && s.Head.Left != nil {
		s.Head = s.Head.Left
	}
}
func (s *SymbolStack) Debug() {
	nav := s.List
	for nav != nil {
		fmt.Printf("%c -> Node: %p (L: %p R: %p)\n", nav.Symbol, nav, nav.Left, nav.Right)
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

func (s *SymbolNode) String() string {
	return string(s.Symbol)
}
