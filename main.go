package main

import (
	"sync"

	"github.com/batt0s/go-linkedlist/linkedlist"
)

func main() {

	ll := new(linkedlist.LinkedList)
	wg := new(sync.WaitGroup)

	wg.Add(5)

	go func() {
		ll.AddToHead(1)
		wg.Done()
	}()
	go func() {
		ll.AddToEnd(2)
		wg.Done()
	}()
	go func() {
		ll.AddToEnd(4)
		wg.Done()
	}()
	go func() {
		ll.AddToEnd(5)
		wg.Done()
	}()
	go func() {
		ll.AddAfter(2, 3)
		wg.Done()
	}()
	//	go func() {
	//		ll.DeleteNode(1)
	//		wg.Done()
	//	}()
	//	go func() {
	//		ll.DeleteNode(3)
	//		wg.Done()
	//	}()

	wg.Wait()

	ll.IterateList()

}
