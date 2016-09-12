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
			[]string{"test"},
		},
		"string2": {
			"",
			"test",
			"test",
			map[string]string{},
		},
		"list1": {
			[]string{"test1", "test2"},
			[]string{},
			[]string{"test1", "test2"},
			"test",
		},
		"list2": {
			[]string{"test1"},
			[]string{"test2", "test3"},
			[]string{"test1", "test2", "test3"},
			map[string]string{},
		},
		"dict1": {
			map[string]string{"key1": "test1", "key2": "test2"},
			map[string]string{"key1": "test3", "key3": "test4"},
			map[string]string{"key1": "test3", "key2": "test2", "key3": "test4"},
			"",
		},
		"dict2": {
			map[string]string{"key1": "test1"},
			map[string]string{},
			map[string]string{"key1": "test1"},
			[]string{"test"},
		},
	}

	invalidValues := map[string]interface{}{
		"test1": 1,
		"test2": []int{1, 2},
		"test3": []interface{}{"1", 2},
		"test4": map[string]int{"key": 1},
		"test5": map[int]string{1: "test"},
		"test6": map[string]interface{}{"test": []string{"test"}},
	}

	t.Run("Empty", func(t *testing.T) {
		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}

		storage.Remove("nonexistent")
	})

	t.Run("Valid", func(t *testing.T) {
		for k, v := range validValues {
			if !storage.Set(k, v.set) {
				t.Errorf("Failed to set (%#v, %#v)", k, v.set)
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

			if _, ok := storage.Get(k); ok {
				t.Errorf("Not expected to get %#v", k)
			}

			if storage.Update(k, "test") {
				t.Errorf("Not expected to update a removed key %#v", k)
			}
		}

		if l := len(storage.Keys()); l != 0 {
			t.Errorf("Expected len is 0, got %d", l)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
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

	t.Run("Update", func(t *testing.T) {
		if storage.Update("nonexistent", "test") {
			t.Error("Not expected to update a nonexistent key")
		}

		for k, vv := range validValues {
			if !storage.Set(k, vv.set) {
				t.Errorf("Failed to set (%#v, %#v)", k, vv.set)
			}

			for _, iv := range invalidValues {
				if storage.Update(k, iv) {
					t.Errorf("Not expected to update (%#v, %#v)", k, iv)
				}

				if v, _ := storage.Get(k); v == iv {
					t.Errorf("Expected %#v, got %#v", vv, iv)
				}
			}

			if !storage.Update(k, vv.validUpdate) {
				t.Errorf("Failed to update (%#v, %#v)", k, vv.validUpdate)
			}

			if actual, _ := storage.Get(k); !reflect.DeepEqual(actual, vv.updated) {
				t.Errorf("Expected %#v, got %#v", vv.updated, actual)
			}

			if storage.Update(k, vv.invalidUpdate) {
				t.Errorf("Not expected to update (%#v, %#v)", k, vv.invalidUpdate)
			}

			if v, _ := storage.Get(k); v == vv.invalidUpdate {
				t.Errorf("Expected %#v, got %#v", vv.updated, vv.invalidUpdate)
			}
		}
	})
}
