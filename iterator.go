package main

type Iterator struct {
	Current int
	CurrentNode *Node
	Data Aggregate
	version int
}

func (i *Iterator) next() int {
	i.Current++;

	aggregateV := i.Data.Version()

	if (aggregateV != i.version) {
		i.version = aggregateV
		return i.Data.get(i.Current - 1)
	} else {
		i.CurrentNode = i.CurrentNode.Next
		i.version = aggregateV
		return i.CurrentNode.Data;
	}
}

func (i *Iterator) isDone() bool {
	return i.Current != i.Data.count()
}