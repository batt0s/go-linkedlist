package main

import (
	"fmt"
	"log"

	"github.com/batt0s/go-linkedlist/pkg/linkedlist"
)

func main() {
	ll, err := linkedlist.New(3, 33, 5, 1)
	if err != nil {
		log.Fatalf("Error while creating new linkedlist.\n%v", err)
	}

	ll.IterateList()
	fmt.Println()

	ll.Sort()

	ll.IterateList()
	fmt.Println()

	ok := ll.DeleteNode(33)
	if !ok {
		log.Fatalln("Error while deleting node.")
	}

	ll.IterateList()

}
