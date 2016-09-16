package storage

import (
	"testing"
	"reflect"
)

func TestStorage(t *testing.T) {
	storage := New()

	validValues := map[string]struct{
		set interface{}
		validUpdate interface{}
		updated interface{}
		invalidUpdate interface{}
	}{
		"": {
			"empty",
			"",
			"",
			[]string{},
		},
		"string1": {
			"test1",
			"test2",
			"test2",
			[]interface{}{"test"},
		},
		"string2": {
			"",
			"test",
			"test",
			map[string]interface{}{},
		},
		"list1": {
			[]interface{}{"test1", 2},
			[]interface{}{},
			[]interface{}{"test1", 2},
			"test",
		},
		"list2": {
			[]interface{}{nil},
			[]interface{}{"test1", "test2"},
			[]interface{}{nil, "test1", "test2"},
			map[string]interface{}{},
		},
		"dict1": {
			map[string]interface{}{"key1": 1, "key2": []interface{}{}},
			map[string]interface{}{"key1": "test1", "key3": "test2"},
			map[string]interface{}{"key1": "test1", "key2": []interface{}{}, "key3": "test2"},
			"",
		},
		"dict2": {
			map[string]interface{}{"key1": map[string]interface{}{}},
			map[string]interface{}{},
			map[string]interface{}{"key1": map[string]interface{}{}},
			[]interface{}{"test"},
		},
	}

	invalidValues := map[string]interface{}{
		"test1": 1,
		"test2": []int{1, 2},
		"test3": map[string]int{"key": 1},
		"test4": map[int]string{1: "test"},
	}

	t.Run("Empty", func(t *testing.T) {
		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}

		storage.Remove("nonexistent")
	})

	t.Run("Valid", func(t *testing.T) {
		for k, v := range validValues {
			if err := storage.Set(k, v.set); err != nil {
				t.Errorf("Failed to set (%#v, %#v): %v", k, v.set, err)
			}
		}

		for _, k := range storage.Keys() {
			expected, ok := validValues[k]

			if !ok {
				t.Errorf("%#v is not expected", k)
			}

			if actual, _ := storage.Get(k); !reflect.DeepEqual(actual, expected.set) {
				t.Errorf("Expected %#v, got %#v", expected.set, actual)
			}

			storage.Remove(k)

			if _, err := storage.Get(k); err == nil {
				t.Errorf("Not expected to get %#v", k)
			}

			if storage.Update(k, "test") == nil {
				t.Errorf("Not expected to update a removed key %#v", k)
			}
		}

		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		for k, v := range invalidValues {
			if storage.Set(k, v) == nil {
				t.Errorf("Not expected to set (%#v, %#v)", k, v)
			}

			if _, err := storage.Get(k); err == nil {
				t.Errorf("Not expected to get (%#v, %#v)", k, v)
			}
		}

		for _, k := range storage.Keys() {
			if _, ok := invalidValues[k]; ok {
				t.Errorf("%#v is not expected", k)
			}
		}
	})

	t.Run("Update", func(t *testing.T) {
		if storage.Update("nonexistent", "test") == nil {
			t.Error("Not expected to update a nonexistent key")
		}

		for k, vv := range validValues {
			if err := storage.Set(k, vv.set); err != nil {
				t.Errorf("Failed to set (%#v, %#v): %v", k, vv.set, err)
			}

			for _, iv := range invalidValues {
				if storage.Update(k, iv) == nil {
					t.Errorf("Not expected to update (%#v, %#v)", k, iv)
				}

				if v, _ := storage.Get(k); v == iv {
					t.Errorf("Expected %#v, got %#v", vv, iv)
				}
			}

			if err := storage.Update(k, vv.validUpdate); err != nil {
				t.Errorf("Failed to update (%#v, %#v): %v", k, vv.validUpdate, err)
			}

			if actual, _ := storage.Get(k); !reflect.DeepEqual(actual, vv.updated) {
				t.Errorf("Expected %#v, got %#v", vv.updated, actual)
			}

			if storage.Update(k, vv.invalidUpdate) == nil {
				t.Errorf("Not expected to update (%#v, %#v)", k, vv.invalidUpdate)
			}

			if v, _ := storage.Get(k); v == vv.invalidUpdate {
				t.Errorf("Expected %#v, got %#v", vv.updated, vv.invalidUpdate)
			}
		}
	})
}
