package storage

import "sync"

type NotFoundError struct {
	detail string
}

func (e *NotFoundError) Error() string {
	return e.detail
}

type TypeError struct {
	detail string
}

func (e *TypeError) Error() string {
	return e.detail
}

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
	return &TypeError{"unsupported type"}
}

func (storage *Storage) Get(key string) (interface{}, error) {
	storage.RLock()
	defer storage.RUnlock()

	if value, ok := storage.items[key]; ok {
		return value, nil
	}
	return nil, &NotFoundError{"key not found"}
}

func (storage *Storage) Remove(key string) error {
	storage.Lock()
	defer storage.Unlock()

	if _, ok := storage.items[key]; ok {
		delete(storage.items, key)
		return nil
	}
	return &NotFoundError{"key not found"}
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
		return &NotFoundError{"key not found"}
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
		return &TypeError{"unsupported type"}
	}
	return &TypeError{"incompatible types"}
}
