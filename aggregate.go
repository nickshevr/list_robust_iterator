package main

type Aggregate struct {
	List    *List
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
	return a.List;
}

func (a *Aggregate) get(index int) int {
	return a.List.get(index);
}

func (a *Aggregate) count() int {
	return a.List.count();
}

func (a *Aggregate) add(value int) {
	a.IncVersion()

	a.List.add(value);
}

func (a *Aggregate) insert(key int, value int) {
	a.IncVersion()

	a.List.insert(key, value);
}

func (a *Aggregate) remove(value int) bool {
	a.IncVersion()

	return a.List.remove(value);
}
