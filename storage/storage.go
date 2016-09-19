package storage

import (
	"strconv"
	"sync"
	"time"
)

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

type item struct {
	value interface{}
	timer *time.Timer
}

type Storage struct {
	items map[string]*item
	sync.RWMutex
}

func New() *Storage {
	return &Storage{items: make(map[string]*item)}
}

func (storage *Storage) Set(key string, i interface{}, ttl time.Duration) error {
	storage.Lock()
	defer storage.Unlock()

	switch value := i.(type) {
	case string, []interface{}, map[string]interface{}:
		if item, ok := storage.items[key]; ok && item.timer != nil {
			item.timer.Stop()
		}

		var timer *time.Timer = nil
		if ttl > 0 {
			timer = time.AfterFunc(ttl, func() {
				storage.Lock()
				defer storage.Unlock()

				delete(storage.items, key)
			})
		}

		storage.items[key] = &item{value: value, timer: timer}
		return nil
	}
	return &TypeError{"unsupported type"}
}

func (storage *Storage) Get(key string) (interface{}, error) {
	storage.RLock()
	defer storage.RUnlock()

	if item, ok := storage.items[key]; ok {
		return item.value, nil
	}
	return nil, &NotFoundError{"key not found"}
}

func (storage *Storage) GetElement(key, element string) (interface{}, error) {
	storage.RLock()
	defer storage.RUnlock()

	item, ok := storage.items[key]
	if !ok {
		return nil, &NotFoundError{"key not found"}
	}

	switch v := item.value.(type) {
	case []interface{}:
		if i, err := strconv.Atoi(element); err == nil {
			if i >= 0 && i < len(v) {
				return v[i], nil
			}
		}
		return nil, &NotFoundError{"element not found"}
	case map[string]interface{}:
		if value, ok := v[element]; ok {
			return value, nil
		}
		return nil, &NotFoundError{"element not found"}
	}
	return nil, &TypeError{"string value do not have elements"}
}

func (storage *Storage) Remove(key string) error {
	storage.Lock()
	defer storage.Unlock()

	if item, ok := storage.items[key]; ok {
		if item.timer != nil {
			item.timer.Stop()
		}

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

	item, ok := storage.items[key]
	if !ok {
		return &NotFoundError{"key not found"}
	}

	switch newValue := i.(type) {
	case string:
		if _, ok := item.value.(string); ok {
			item.value = newValue
			return nil
		}
	case []interface{}:
		if l, ok := item.value.([]interface{}); ok {
			item.value = append(l, newValue...)
			return nil
		}
	case map[string]interface{}:
		if m, ok := item.value.(map[string]interface{}); ok {
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
