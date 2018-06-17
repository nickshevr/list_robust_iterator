package main

type Aggregate struct {
	List    *List
	version int
	history []Log
}

func (a *Aggregate) Version() int {
	return a.version;
}

func (a *Aggregate) IncVersion() int {
	a.version++;

	return a.Version();
}

func (a *Aggregate) getHistory(version int) []Log {
	return a.history[version:a.version]
}

// actions: insert - 0, remove - 1
func (a *Aggregate) pushHistory(index int, action int) {
	a.history = append(a.history, Log{
		action,
		index,
	})
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
	// добавление в хвост не влияет на робастность
	a.List.add(value);
}

func (a *Aggregate) insert(key int, value int) {
	a.IncVersion()

	node := a.List.insert(key, value);
	index := a.List.findNode(node)

	a.pushHistory(index, 0)
}

func (a *Aggregate) remove(value int) bool {
	a.IncVersion()

	index, err := a.List.rm(value);
	if err != nil {
		return false
	}

	a.pushHistory(index, 1)

	return true
}
