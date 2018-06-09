// Package memstore is an in-memory mock data source.
//
// Use memstore until a working persistence layer is in place.
package memstore

// Player describes a single player entity.
// TODO: Mind file declarations.
type Player struct {
	Name string
	Wins int
}

// Player describes a data memstore.
type MemStore struct {
	store map[string]int
}

// New initializes a MemStore.
func New() *MemStore {
	return &MemStore{map[string]int{}}
}

// RecordWin increment a player's score.
func (i *MemStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore retrieves a player's score.
func (i *MemStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// GetLeague ...
func (i *MemStore) GetLeague() []Player {
	return nil
}
