package main

import (
	"testing"
)

func TestListAdd(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2

	l.add(firstElem)
	l.add(lastElem)

	if l.rootNode.Data != firstElem {
		t.Errorf("First elem isnt right: %d, want: %d.", l.rootNode.Data, firstElem)
	}

	if l.endNode.Data != lastElem {
		t.Errorf("Last elem isnt right: %d, want: %d.", l.endNode.Data, lastElem)
	}
}

func TestListFind(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2

	l.add(firstElem)
	l.add(lastElem)


	var findedElem = l.find(lastElem)

	if l.endNode != findedElem {
		t.Errorf("Finding elem isnt right: %d, want: %d.", l.rootNode, findedElem)
	}
}

func TestListInsert(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2
	var insertElem = 3

	l.add(firstElem)
	l.add(lastElem)

	insertNode := l.insert(firstElem, insertElem)

	if insertNode.Data != insertElem {
		t.Errorf("Insert wrong elem: %d, want: %d.", insertNode.Data, insertElem)
	}

	nextElem := l.find(lastElem)

	if insertNode.Next !=  nextElem {
		t.Errorf("Wrong chain")
	}
}

func TestListGetZero(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	firstElem := 1


	l.add(firstElem)

	node := l.get(0)

	if node != firstElem {
		t.Errorf("Getting wrong elem: %d, want: %d.", node, firstElem)
	}
}

func TestListGetFirst(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	firstElem := 1
	secondElem := 2

	l.add(firstElem)
	l.add(secondElem)

	node := l.get(1)

	if node != secondElem {
		t.Errorf("Getting wrong elem: %d, want: %d.", node, firstElem)
	}
}

func TestListCount(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	firstElem := 1
	secondElem := 2

	l.add(firstElem)
	l.add(secondElem)

	length := l.count()

	if length != 2 {
		t.Errorf("Getting wrong elem: %d, want: %d.", length, 2)
	}
}

func TestListRemove(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	firstElem := 1
	secondElem := 2
	thirdElem := 3

	l.add(firstElem)
	l.add(secondElem)
	l.add(thirdElem)

	l.remove(secondElem)

	length := l.count()

	if l.rootNode.Next != l.endNode {
		t.Errorf("Deletion was wrong: %d, want: %d.", l.rootNode.Next.Data, l.endNode.Data)
	}

	if length != 2 {
		t.Errorf("Deletion was wrong length: %d, want: %d.", length, 1)
	}
}