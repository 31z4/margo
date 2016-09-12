package storage

type Storage map[string]interface{}

func New() Storage {
	return make(map[string]interface{})
}

func (storage *Storage) Set(key string, i interface{}) bool {
	switch value := i.(type) {
	case string, []string, map[string]string:
		(*storage)[key] = value
		return true
	}
	return false
}

func (storage *Storage) Get(key string) (interface{}, bool) {
	value, ok := (*storage)[key]
	return value, ok
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

func (storage *Storage) Update(key string, i interface{}) bool {
	currentValue, ok := (*storage)[key]
	if !ok {
		return false
	}

	switch newValue := i.(type) {
	case string:
		if _, ok := currentValue.(string); ok {
			(*storage)[key] = newValue
			return true
		}
	case []string:
		if l, ok := currentValue.([]string); ok {
			(*storage)[key] = append(l, newValue...)
			return true
		}
	case map[string]string:
		if m, ok := currentValue.(map[string]string); ok {
			for k, v := range newValue {
				m[k] = v
			}
			return true
		}
	}
	return false
}
