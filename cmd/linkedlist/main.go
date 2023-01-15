package main

import (
	"github.com/batt0s/go-linkedlist/pkg/linkedlist"
)

func main() {
	ll := linkedlist.LinkedList{}
	ll.AddToHead(3)
	ll.AddToEnd(5)
	ll.AddAfter(3, 33)
	ll.AddToEnd(1)

	ll.IterateList()

	ll.Sort()

	ll.IterateList()

}
