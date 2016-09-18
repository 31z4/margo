package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	items map[string]interface{}
	sync.RWMutex
}

func New() *Storage {
	return &Storage{items: make(map[string]interface{})}
}

func (storage *Storage) Set(key string, i interface{}) error {
	storage.Lock()
	defer storage.Unlock()

	switch value := i.(type) {
	case string, []interface{}, map[string]interface{}:
		storage.items[key] = value
		return nil
	}
	return errors.New("Unsupported type")
}

func (storage *Storage) Get(key string) (interface{}, error) {
	storage.RLock()
	defer storage.RUnlock()

	if value, ok := storage.items[key]; ok {
		return value, nil
	}
	return nil, errors.New("Not found")
}

func (storage *Storage) Remove(key string) {
	storage.Lock()
	defer storage.Unlock()

	delete(storage.items, key)
}

func (storage *Storage) Keys() []string {
	storage.RLock()
	defer storage.RUnlock()

	keys := make([]string, 0, len(storage.items))
	for k := range storage.items {
		keys = append(keys, k)
	}
	return keys
}

func (storage *Storage) Update(key string, i interface{}) error {
	storage.Lock()
	defer storage.Unlock()

	currentValue, ok := storage.items[key]
	if !ok {
		return errors.New("Not found")
	}

	switch newValue := i.(type) {
	case string:
		if _, ok := currentValue.(string); ok {
			storage.items[key] = newValue
			return nil
		}
	case []interface{}:
		if l, ok := currentValue.([]interface{}); ok {
			storage.items[key] = append(l, newValue...)
			return nil
		}
	case map[string]interface{}:
		if m, ok := currentValue.(map[string]interface{}); ok {
			for k, v := range newValue {
				m[k] = v
			}
			return nil
		}
	default:
		return errors.New("Unsupported type")
	}
	return errors.New("Incompatible types")
}
