package main

type Aggregate struct {
	Data *List
	version int
}

func (a *Aggregate) Version() int {
	return a.version;
}

func (a *Aggregate) IncVersion() int {
	a.version++;

	return a.Version();
}

func (a *Aggregate) GetData() *List {
	return a.Data;
}

func (a *Aggregate) get(index int) int {
	return a.Data.get(index);
}

func (a *Aggregate) count() int {
	return a.Data.count();
}

func (a *Aggregate) add(value int) {
	a.IncVersion()

	a.Data.add(value);
}

func (a *Aggregate) remove(value int) bool {
	a.IncVersion()

	return a.Data.remove(value);
}
