package main

import (
	"testing"
)

func TestAggregateInsert(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2

	l.add(firstElem)
	l.add(lastElem)

	a := Aggregate{
		version: 0,
		List: l,
	}

	a.insert(1, 3)

	assertIntEqual(t, a.version, 1)
	assertIntEqual(t, a.history[0].action, 0)
	assertIntEqual(t, a.history[0].index, 1)

	a.insert(3, 5)

	assertIntEqual(t, a.version, 2)
	assertIntEqual(t, a.history[0].action, 0)
	assertIntEqual(t, a.history[0].index, 1)
	assertIntEqual(t, a.history[1].action, 0)
	assertIntEqual(t, a.history[1].index, 2)
}

func TestAggregateRemove(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2

	l.add(firstElem)
	l.add(lastElem)

	a := Aggregate{
		version: 0,
		List: l,
	}

	a.remove(1)

	assertIntEqual(t, a.version, 1)
	assertIntEqual(t, a.history[0].action, 1)
	assertIntEqual(t, a.history[0].index, 0)

	a.remove(2)

	assertIntEqual(t, a.version, 2)
	assertIntEqual(t, a.history[0].action, 1)
	assertIntEqual(t, a.history[0].index, 0)
	assertIntEqual(t, a.history[1].action, 1)
	assertIntEqual(t, a.history[1].index, 0)
}

func TestAggregateGetHistory(t *testing.T) {
	l := &List{
		rootNode: nil,
		endNode: nil,
	}

	var firstElem = 1
	var lastElem = 2

	l.add(firstElem)
	l.add(lastElem)

	a := Aggregate{
		version: 0,
		List: l,
	}

	a.insert(1, 3)
	a.insert(3, 5)

	h := a.getHistory(0)

	assertIntEqual(t, h[0].action, 0)
	assertIntEqual(t, h[0].index, 1)
	assertIntEqual(t, h[1].action, 0)
	assertIntEqual(t, h[1].index, 2)
}