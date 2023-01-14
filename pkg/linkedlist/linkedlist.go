package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	headNode *Node
}

type Node struct {
	property int
	nextNode *Node
}

// Get Node with Value
func (ll *LinkedList) NodeWithVal(prop int) (*Node, bool) {
	node := ll.headNode
	for node != nil {
		fmt.Println(node)
		if node.property == prop {
			return node, true
		}
		node = node.nextNode
	}
	return nil, false
}

// Get Last Node
func (ll *LinkedList) LastNode() (*Node, bool) {
	var node *Node
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node, true
		}
	}
	return nil, false
}

// Add to head
func (ll *LinkedList) AddToHead(prop int) {
	var node = &Node{
		property: prop,
		nextNode: nil,
	}
	if ll.headNode != nil {
		node.nextNode = ll.headNode
	}
	ll.headNode = node
}

// Add to end
func (ll *LinkedList) AddToEnd(prop int) {
	node := &Node{
		property: prop,
		nextNode: nil,
	}
	lastNode, _ := ll.LastNode()
	lastNode.nextNode = node
}

// Add after value
func (ll *LinkedList) AddAfter(addAfter, prop int) bool {
	node := &Node{
		property: prop,
		nextNode: nil,
	}
	nodeAddAfter, ok := ll.NodeWithVal(addAfter)
	if !ok {
		return false
	}
	node.nextNode = nodeAddAfter.nextNode
	nodeAddAfter.nextNode = node
	return true
}

// Delete Node
func (ll *LinkedList) DeleteNode(prop int) bool {
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
	prevNode := ll.headNode
	for prevNode != nil {
		if prevNode.nextNode == node {
			prevNode.nextNode = node.nextNode
			return true
		}
	}
	return false
}

func (ll *LinkedList) IterateList() {
	for n := ll.headNode; n != nil; n = n.nextNode {
		fmt.Printf("%+v ", n.property)
	}
}

// TODO Sort()
