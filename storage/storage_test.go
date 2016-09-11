package storage

import (
	"testing"
	"reflect"
)

func TestStorage(t *testing.T) {
	storage := New()

	t.Run("Empty", func(t *testing.T) {
		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}

		storage.Remove("nonexistent")
	})

	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		validValues := map[string]interface{}{
			"": "empty",
			"string1": "test",
			"string2": "",
			"list1": []string{"test1", "test2"},
			"list2": []string{},
			"dict1": map[string]string{"key1": "test1", "key2": "test2"},
			"dict2": map[string]string{},
		}

		for k, v := range validValues {
			if !storage.Set(k, v) {
				t.Errorf("Failed to set (%#v, %#v)", k, v)
			}
		}

		for _, k := range storage.Keys() {
			expected, ok := validValues[k]

			if !ok {
				t.Errorf("%#v is not expected", k)
			}

			if actual, _ := storage.Get(k); !reflect.DeepEqual(actual, expected) {
				t.Errorf("Expected %#v, got %#v", expected, actual)
			}

			storage.Remove(k)
			if _, ok := storage.Get(k); ok {
				t.Errorf("Not expected to get %#v", k)
			}
		}

		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		t.Parallel()

		invalidValues := map[string]interface{}{
			"test1": 1,
			"test2": []int{1, 2},
			"test3": []interface{}{"1", 2},
			"test4": map[string]int{"key": 1},
			"test5": map[int]string{1: "test"},
			"test6": map[string]interface{}{"test": []string{"test"}},
		}

		for k, v := range invalidValues {
			if storage.Set(k, v) {
				t.Errorf("Not expected to set (%#v, %#v)", k, v)
			}

			if _, ok := storage.Get(k); ok {
				t.Errorf("Not expected to get (%#v, %#v)", k, v)
			}
		}

		for _, k := range storage.Keys() {
			if _, ok := invalidValues[k]; ok {
				t.Errorf("%#v is not expected", k)
			}
		}
	})
}
