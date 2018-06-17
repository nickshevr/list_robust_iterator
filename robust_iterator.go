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
	r.syncData()

	return r.CurrentIterator.isDone()
}

func (r *RobustIterator) apllyLog(log Log) {
	//delete
	if (log.action == 1) {
		if (r.Current < log.index || r.Current - 1 == log.index)  {
			r.Current--
		}
	}
	//insert
	if (log.action == 0) {
		if (r.Current > log.index)  {
			r.Current++
		}
	}
}

func (r *RobustIterator) syncIndex() {
	h := r.Data.getHistory(r.version)

	for i := range h {
		r.apllyLog(h[i])
	}
}

func (r *RobustIterator) syncData() {
	aggregateV := r.Data.Version()

	if (aggregateV != r.version) {
		r.syncIndex()

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
