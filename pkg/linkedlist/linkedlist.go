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

func (ll *LinkedList) String() string {
	return fmt.Sprintf("LinkedList{%v}", ll.AsArray())
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{prop: %d, next: %v}", n.property, n.nextNode)
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

// Length/Number of nodes
func (ll *LinkedList) Length() int {
	length := 0
	for n := ll.headNode; n != nil; n = n.nextNode {
		length++
	}
	return length
}

// Print list
func (ll *LinkedList) IterateList() {
	for n := ll.headNode; n != nil; n = n.nextNode {
		fmt.Printf("%+v ", n.property)
	}
}

// Return properties as array
func (ll *LinkedList) AsArray() []int {
	var arr []int
	for n := ll.headNode; n != nil; n = n.nextNode {
		arr = append(arr, n.property)
	}
	return arr
}

// Add two linkedlists
func AddLinkedLists(ll1 *LinkedList, ll2 *LinkedList) *LinkedList {
	var ll *LinkedList
	n, ok := ll1.LastNode()
	if !ok {
		return nil
	}
	n.nextNode = ll2.headNode
	ll.headNode = ll1.headNode
	return ll
}

// Sort the list
func (ll *LinkedList) Sort() {
	if ll.headNode == nil {
		return
	}
	if ll.headNode.nextNode == nil {
		return
	}
	n := ll.Length()
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		h := ll.headNode
		swapped := false
		for j := 0; j < n-i-1; j++ {
			n1 := h
			n2 := n1.nextNode
			fmt.Printf("n1: %s, n1.next: %s\n", n1.String(), n1.nextNode.String())
			if n1.property > n2.property {
				swap(n1, n2)
				swapped = true
			}
			h = h.nextNode
		}
		if !swapped {
			break
		}
	}
}

func swap(node1 *Node, node2 *Node) {
	temp := node2.nextNode
	node2.nextNode = node1
	node1.nextNode = temp
}
