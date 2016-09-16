package storage

import (
	"errors"
)

type Storage map[string]interface{}

func New() Storage {
	return make(map[string]interface{})
}

func (storage *Storage) Set(key string, i interface{}) error {
	switch value := i.(type) {
	case string, []interface{}, map[string]interface{}:
		(*storage)[key] = value
		return nil
	}
	return errors.New("Unsupported type")
}

func (storage *Storage) Get(key string) (interface{}, error) {
	if value, ok := (*storage)[key]; ok {
		return value, nil
	}
	return nil, errors.New("Not found")
}

func (storage *Storage) Remove(key string) {
	delete(*storage, key)
}

func (storage *Storage) Keys() []string {
	keys := make([]string, 0, len(*storage))
	for k := range *storage {
		keys = append(keys, k)
	}
	return keys
}

func (storage *Storage) Update(key string, i interface{}) error {
	currentValue, ok := (*storage)[key]
	if !ok {
		return errors.New("Not found")
	}

	switch newValue := i.(type) {
	case string:
		if _, ok := currentValue.(string); ok {
			(*storage)[key] = newValue
			return nil
		}
	case []interface{}:
		if l, ok := currentValue.([]interface{}); ok {
			(*storage)[key] = append(l, newValue...)
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
