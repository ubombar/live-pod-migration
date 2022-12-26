package structures

import (
	"errors"
	"sync"
)

type Store interface {
	Add(name string, obj interface{})
	Fetch(name string) (interface{}, error)
	Name() string
	Delete(name string) error
}

type store struct {
	mutex sync.Mutex
	store map[string]interface{}
	name  string
}

func NewStore(name string) *store {
	m := &store{
		store: make(map[string]interface{}),
		name:  name,
		mutex: sync.Mutex{},
	}

	return m
}

// Adds the given object to the memory.
func (m *store) Add(name string, obj interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.store[name] = obj
}

// Fetches the object, returns an error if not found.
func (m *store) Fetch(name string) (interface{}, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if obj, ok := m.store[name]; ok {
		return obj, nil
	}

	return nil, errors.New("cannot find interface with given name")
}

// Get the name of the store
func (m *store) Name() string {
	return m.name
}

func (m *store) Delete(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.store[name]

	if !ok {
		return errors.New("item is not in the store")
	}

	delete(m.store, name)

	return nil
}
