package linkedlist_test

import (
	"testing"

	"github.com/batt0s/go-linkedlist/pkg/linkedlist"
)

var testList *linkedlist.LinkedList = new(linkedlist.LinkedList)

func TestAddToHead(t *testing.T) {
	testList.AddToHead(1)
	arr := testList.AsArray()
	if arr[0] != 1 {
		t.Errorf("got %q, want %q", arr[0], 1)
	}
}
