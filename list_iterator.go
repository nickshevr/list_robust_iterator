package  main

import "strconv"

type ListIterator struct {
	Current *Node
	ListData *List
}

func (i *ListIterator) next() int {
	if (i.Current == nil) {
		i.Current = i.ListData.rootNode

		return i.Current.Data
	}

	i.Current = i.Current.Next
	data := i.Current.Data

	return data
}

func (i *ListIterator) isDone() bool {
	return i.Current != nil && i.Current.Next == nil || i.ListData.rootNode == nil
}

func (i *ListIterator) Stringify() string {
	s := ""

	for i.isDone() != true {
		s +=  strconv.FormatInt(int64(i.next()), 10)
	}

	return s
}
