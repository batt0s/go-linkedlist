package linkedlist

import (
	"fmt"
	"sync"
)

type LinkedList struct {
	headNode *Node
	mu       sync.Mutex
}

type Node struct {
	property int
	nextNode *Node
}

func (ll *LinkedList) NodeWithVal(prop int) (*Node, bool) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == prop {
			return node, true
		}
	}
	return nil, false
}

func (ll *LinkedList) LastNode() (*Node, bool) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node, true
		}
	}
	return nil, false
}

func (ll *LinkedList) AddToHead(prop int) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	var node = &Node{
		property: prop,
		nextNode: nil,
	}
	if ll.headNode != nil {
		node.nextNode = ll.headNode
	}
	ll.headNode = node
}

func (ll *LinkedList) AddToEnd(prop int) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	node := &Node{
		property: prop,
		nextNode: nil,
	}
	lastNode, _ := ll.LastNode()
	lastNode.nextNode = node
}

func (ll *LinkedList) AddAfter(addAfter, prop int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	node := &Node{
		property: prop,
		nextNode: nil,
	}
	var nodeAddAfter *Node
	for nodeAddAfter = ll.headNode; nodeAddAfter != nil; nodeAddAfter = nodeAddAfter.nextNode {
		if nodeAddAfter.property == addAfter {
			node.nextNode = nodeAddAfter.nextNode
			nodeAddAfter.nextNode = node
			return true
		}
	}
	return false
}

func (ll *LinkedList) DeleteNode(prop int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == prop {
			break
		}
	}
	if node == ll.headNode {
		if node.nextNode != nil {
			ll.headNode = node.nextNode
			return true
		} else {
			return false
		}
	}
	var prevNode *Node = ll.headNode
	for prevNode != nil {
		if prevNode.nextNode == node {
			prevNode.nextNode = node.nextNode
			return true
		}
	}
	return false
}

func (ll *LinkedList) IterateList() {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	for n := ll.headNode; n != nil; n = n.nextNode {
		fmt.Printf("%+v ", n.property)
	}
}
