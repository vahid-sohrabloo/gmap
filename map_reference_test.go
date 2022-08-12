package gmap_test

import "sync"

// This file contains reference map implementations for unit-tests.

// mapInterface is the interface Map implements.
type mapInterface interface {
	Load(string) (string, bool)
	Store(key, value string)
	LoadOrStore(key, value string) (actual string, loaded bool)
	LoadAndDelete(key string) (value string, loaded bool)
	Delete(string)
	Range(func(key, value string) (shouldContinue bool))
}

// SyncMap is an implementation of mapInterface using a sync.map.
type SyncMap struct {
	m sync.Map
}

func (m *SyncMap) Load(key string) (string, bool) {
	value, ok := m.m.Load(key)
	return value.(string), ok
}

func (m *SyncMap) Store(key, value string) {
	m.m.Store(key, value)
}

func (m *SyncMap) LoadOrStore(key, value string) (string, bool) {
	actual, loaded := m.m.LoadOrStore(key, value)
	return actual.(string), loaded
}

func (m *SyncMap) LoadAndDelete(key any) (string, bool) {
	value, loaded := m.m.LoadAndDelete(key)
	return value.(string), loaded
}

func (m *SyncMap) Delete(key string) {
	m.m.Delete(key)
}

func (m *SyncMap) Range(f func(key, value string) bool) {
	m.m.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(string))
	})
}
