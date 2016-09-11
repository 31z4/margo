package storage

type Storage map[string]interface{}

func New() Storage {
	return make(map[string]interface{})
}

func (storage *Storage) Set(key string, i interface{}) bool {
	switch v := i.(type) {
	case string, []string, map[string]string:
		(*storage)[key] = v
		return true
	default:
		return false
	}
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
