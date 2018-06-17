package main

import (
	"testing"
)

func TestRobustBaseIterate(t *testing.T) {
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

	res := robust.Stringify()

	if res != "123" {
		t.Errorf("Iteration isnt right: %s, want: %s.", res, "123")
	}
}

func TestRobustIterateAddition(t *testing.T) {
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

	insertElem := 6

	aggregate.insert(secondElem, insertElem);

	third := robust.next()

	assertIntEqual(t, third, insertElem)

	assertIntEqual(t, robust.next(), thirdElem)
}

func TestRobustIterateNextDeletion(t *testing.T) {
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

	flag := aggregate.remove(2);

	assertBoolEqual(t, flag, true)

	second := robust.next()

	assertIntEqual(t, second, thirdElem)
}

func TestRobustIterateAddRemove(t *testing.T) {
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

	aggregate.add(6);
	aggregate.add(7);

	third := robust.next()

	assertIntEqual(t, third, thirdElem)
	assertIntEqual(t, robust.next(), 6)
	assertIntEqual(t, robust.next(), 7)
}

func TestRobustIterateStory(t *testing.T) {
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

	aggregate.add(6);
	aggregate.add(7);

	third := robust.next()

	assertIntEqual(t, third, thirdElem)
	flag := aggregate.remove(6);

	assertBoolEqual(t, flag, true)

	assertIntEqual(t, robust.next(), 7)
}
