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

// experiment
func (r *RobustIterator) insertBefore(key int, value int) {
	r.Current++

	r.Data.insert(key, value)
}

func (r *RobustIterator) insertAfter(key int, value int) {
	r.Data.insert(key, value)
}

func (r *RobustIterator) removeBefore(value int) bool {
	r.Current--

	return r.Data.remove(value)
}

func (r *RobustIterator) removeAfter(value int) bool {
	return r.Data.remove(value)
}
// exp end

func (r *RobustIterator) Stringify() string {
	s := ""

	for r.isDone() != true {
		s +=  strconv.FormatInt(int64(r.next()), 10)
	}

	return s
}
