package linkedlist_test

import (
	"log"
	"reflect"
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

func TestAddToEnd(t *testing.T) {
	testList.AddToEnd(3)
	arr := testList.AsArray()
	if arr[len(arr)-1] != 3 {
		t.Errorf("got %q, want %q", arr[len(arr)-1], 3)
	}
}

func TestAddAfter(t *testing.T) {
	testList.AddAfter(1, 11)
	arr := testList.AsArray()
	if arr[1] != 11 {
		t.Errorf("got %q, want %q", arr[1], 11)
	}
}

func TestSort(t *testing.T) {
	testList.Sort()
	arr := testList.AsArray()
	if !isSorted(arr) {
		t.Errorf("arr is not sorted")
	}
}
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			log.Printf("%d < %d (%d < %d)", arr[i], arr[i-1], i, i-1)
			return false
		}
	}
	return true
}

func TestDeleteNode(t *testing.T) {
	testList.DeleteNode(1)
	arr := testList.AsArray()
	if arr[0] == 1 {
		t.Errorf("got %q, dont want that", arr[0])
	}
}

func TestAddLinkedLists(t *testing.T) {
	ll1, err := linkedlist.New(1, 2)
	if err != nil {
		t.Errorf("got an error: %s", err.Error())
	}
	ll2, err := linkedlist.New(3, 4)
	if err != nil {
		t.Errorf("got an error: %s", err.Error())
	}
	ll3 := linkedlist.AddLinkedLists(ll1, ll2)
	arr := ll3.AsArray()
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4}) {
		t.Errorf("got %q, want %q", arr, []int{1, 2, 3, 4})
	}
}
