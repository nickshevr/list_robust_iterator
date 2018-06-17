package main

import "strconv"

type RobustIterator struct {
	Current int
	CurrentIterator *ListIterator
	Data *Aggregate
	version int
}

func (r *RobustIterator) next() int {
	r.syncData()

	r.Current++;

	return r.CurrentIterator.next();
}

func (r *RobustIterator) isDone() bool {
	return r.CurrentIterator.isDone()
}

func (r *RobustIterator) syncData() {
	aggregateV := r.Data.Version()

	if (aggregateV != r.version) {
		itertaror := &ListIterator{
			ListData: r.Data.GetData(),
			Current: r.Data.List.getNode(r.Current - 1),
		}

		r.CurrentIterator = itertaror;
	}

	r.version = aggregateV
}

func (r *RobustIterator) Stringify() string {
	s := ""

	for r.isDone() != true {
		s +=  strconv.FormatInt(int64(r.next()), 10)
	}

	return s
}
