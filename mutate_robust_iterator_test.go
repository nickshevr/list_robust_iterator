package main

import (
	"testing"
)

func TestRemoveFromRobust(t *testing.T) {
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

	iterator := &ListIterator{
		ListData: l,
	}

	aggregate := &Aggregate{
		version: 0,
		List:    l,
	}

	robust := RobustIterator{
		Current: 0,
		CurrentIterator: iterator,
		Data: aggregate,
		version: aggregate.version,
	}

	first := robust.next()

	assertIntEqual(t, first, firstElem)

	flag := robust.removeBefore(firstElem)

	assertBoolEqual(t, flag, true)

	second := robust.next()

	assertIntEqual(t, second, secondElem)

	assertIntEqual(t, robust.next(), thirdElem)
}

func TestRemoveTailFromRobust(t *testing.T) {
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

	iterator := &ListIterator{
		ListData: l,
	}

	aggregate := &Aggregate{
		version: 0,
		List:    l,
	}

	robust := RobustIterator{
		Current: 0,
		CurrentIterator: iterator,
		Data: aggregate,
		version: aggregate.version,
	}

	first := robust.next()

	assertIntEqual(t, first, firstElem)

	second := robust.next()

	assertIntEqual(t, second, secondElem)

	flag := robust.removeAfter(thirdElem)

	assertBoolEqual(t, flag, true)

	assertBoolEqual(t, robust.isDone(), true)
}

func TestAddBeforeFromRobust(t *testing.T) {
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

	iterator := &ListIterator{
		ListData: l,
	}

	aggregate := &Aggregate{
		version: 0,
		List:    l,
	}

	robust := RobustIterator{
		Current: 0,
		CurrentIterator: iterator,
		Data: aggregate,
		version: aggregate.version,
	}

	first := robust.next()

	assertIntEqual(t, first, firstElem)

	second := robust.next()

	assertIntEqual(t, second, secondElem)

	robust.insertBefore(firstElem, 6)
	robust.insertBefore(firstElem, 7)

	assertIntEqual(t, robust.next(), thirdElem)
}