package structures

import "errors"

type Store interface {
	Add(name string, obj interface{})
	Fetch(name string) (interface{}, error)
	Name() string
	Delete(name string) error
}

type store struct {
	store map[string]interface{}
	name  string
}

func NewStore(name string) *store {
	m := &store{
		store: make(map[string]interface{}),
		name:  name,
	}

	return m
}

// Adds the given object to the memory.
func (m *store) Add(name string, obj interface{}) {
	m.store[name] = obj
}

// Fetches the object, returns an error if not found.
func (m *store) Fetch(name string) (interface{}, error) {
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
	_, ok := m.store[name]

	if !ok {
		return errors.New("item is not in the store")
	}

	delete(m.store, name)

	return nil
}
