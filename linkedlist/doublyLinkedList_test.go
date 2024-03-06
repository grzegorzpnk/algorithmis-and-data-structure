package linkedlist

import (
	"testing"
)

func TestInsertFront(t *testing.T) {

	dll := DoubledLinkedList{}

	dll.InsertFront(1)
	if dll.Head.value != 1 {
		t.Errorf("Expected Head value to be 1, got %d", dll.Head.value)
	}
}

func TestInsertEnd(t *testing.T) {
	dll := DoubledLinkedList{}

	dll.InsertEnd(1)
	if dll.Tail.value != 1 {
		t.Errorf("Epected Tail value to be 1, got %d", dll.Tail.value)
	}

}
