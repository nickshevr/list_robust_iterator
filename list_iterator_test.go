package main

import (
	"testing"
)

func TestListIterate(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	firstElem := 1
	lastElem := 2
	thirdElem := 3

	l.add(firstElem)
	l.add(lastElem)
	l.add(thirdElem)

	iterator := ListIterator{
		ListData: l,
	}

	res := iterator.Stringify()

	if res != "123" {
		t.Errorf("Iteration isnt right: %s, want: %s.", res, "123")
	}
}

func TestNullListIterate(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	iterator := ListIterator{
		ListData: l,
	}

	if iterator.isDone() != true {
		t.Errorf("Error with null list")
	}
}