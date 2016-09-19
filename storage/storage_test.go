package storage

import (
	"reflect"
	"testing"
	"time"
)

var validValues = map[string]struct {
	set           interface{}
	validUpdate   interface{}
	updated       interface{}
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

var invalidValues = map[string]interface{}{
	"test1": 1,
	"test2": []int{1, 2},
	"test3": map[string]int{"key": 1},
	"test4": map[int]string{1: "test"},
}

func TestEmpty(t *testing.T) {
	storage := New()

	if l := len(storage.Keys()); l != 0 {
		t.Errorf("Expected len is 0, got %d", l)
	}

	if err := storage.Remove("nonexistent"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to remove nonexistent")
	}
}

func TestValid(t *testing.T) {
	storage := New()

	for k, v := range validValues {
		if err := storage.Set(k, v.set, 0); err != nil {
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

		if err := storage.Remove(k); err != nil {
			t.Errorf("Failed to remove %#v: %v", k, err)
		}

		if _, err := storage.Get(k); err != nil {
			if _, ok := err.(*NotFoundError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to get %#v", k)
		}

		if err := storage.Update(k, "test"); err != nil {
			if _, ok := err.(*NotFoundError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to update %#v", k)
		}

		if err := storage.Remove(k); err != nil {
			if _, ok := err.(*NotFoundError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to remove %#v", k)
		}
	}

	if l := len(storage.Keys()); l != 0 {
		t.Errorf("Expected len is 0, got %d", l)
	}
}

func TestInvalid(t *testing.T) {
	storage := New()

	for k, v := range invalidValues {
		if err := storage.Set(k, v, 0); err != nil {
			if _, ok := err.(*TypeError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to set (%#v, %#v)", k, v)
		}

		if _, err := storage.Get(k); err != nil {
			if _, ok := err.(*NotFoundError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to get (%#v, %#v)", k, v)
		}
	}

	for _, k := range storage.Keys() {
		if _, ok := invalidValues[k]; ok {
			t.Errorf("%#v is not expected", k)
		}
	}
}

func TestUpdate(t *testing.T) {
	storage := New()

	if err := storage.Update("nonexistent", "test"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to update a nonexistent key")
	}

	for k, vv := range validValues {
		if err := storage.Set(k, vv.set, 0); err != nil {
			t.Errorf("Failed to set (%#v, %#v): %v", k, vv.set, err)
		}

		for _, iv := range invalidValues {
			if err := storage.Update(k, iv); err != nil {
				if _, ok := err.(*TypeError); !ok {
					t.Errorf("Unexpected error: %#v", err)
				}
			} else {
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

		if err := storage.Update(k, vv.invalidUpdate); err != nil {
			if _, ok := err.(*TypeError); !ok {
				t.Errorf("Unexpected error: %#v", err)
			}
		} else {
			t.Errorf("Not expected to update (%#v, %#v)", k, vv.invalidUpdate)
		}

		if v, _ := storage.Get(k); v == vv.invalidUpdate {
			t.Errorf("Expected %#v, got %#v", vv.updated, vv.invalidUpdate)
		}
	}
}

func TestGetElement(t *testing.T) {
	storage := New()

	values := map[string]interface{}{
		"string": "test",
		"list":   []interface{}{"test1", "test2"},
		"dict":   map[string]interface{}{"key1": "test1", "key2": "test2"},
	}

	for k, v := range values {
		if err := storage.Set(k, v, 0); err != nil {
			t.Errorf("Failed to set (%#v, %#v): %v", k, v, err)
		}
	}

	if _, err := storage.GetElement("nonexistent", "nonexistent"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get a nonexistent key")
	}

	if _, err := storage.GetElement("string", "nonexistent"); err != nil {
		if _, ok := err.(*TypeError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get an element from a string")
	}

	if _, err := storage.GetElement("list", "nonexistent"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get a string element from a list")
	}

	if _, err := storage.GetElement("list", "-1"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get a negative element from a list")
	}

	if _, err := storage.GetElement("list", "2"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get a nonexistent element from a list")
	}

	if _, err := storage.GetElement("dict", "nonexistent"); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Error("Not expected to get a nonexistent element from a dict")
	}

	const (
		listKey      = "list"
		listElement  = "01"
		listExpected = "test2"
	)
	if actual, err := storage.GetElement(listKey, listElement); err == nil {
		if actual != listExpected {
			t.Errorf("Expected %#v, got %#v", listExpected, actual)
		}
	} else {
		t.Errorf("Failed to get (%#v, %#v): %v", listKey, listElement, err)
	}

	const (
		dictKey      = "dict"
		dictElement  = "key1"
		dictExpected = "test1"
	)
	if actual, err := storage.GetElement(dictKey, dictElement); err == nil {
		if actual != dictExpected {
			t.Errorf("Expected %#v, got %#v", dictExpected, actual)
		}
	} else {
		t.Errorf("Failed to get (%#v, %#v): %v", dictKey, dictElement, err)
	}
}

func TestTTL(t *testing.T) {
	storage := New()

	const (
		key      = "key"
		value    = "value"
		duration = time.Second / 2
	)

	// Update doesn't affect TTL
	if err := storage.Set(key, value, duration); err != nil {
		t.Errorf("Failed to set (%#v, %#v): %v", key, value, err)
	}

	if err := storage.Update(key, value); err != nil {
		t.Errorf("Failed to update (%#v, %#v): %v", key, value, err)
	}

	if actual, _ := storage.Get(key); actual != value {
		t.Errorf("Expected %#v, got %#v", value, actual)
	}

	time.Sleep(2 * duration)

	if _, err := storage.Get(key); err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Unexpected error: %#v", err)
		}
	} else {
		t.Errorf("Not expected to get %#v", key)
	}

	// Set overrides TTL
	if err := storage.Set(key, value, duration); err != nil {
		t.Errorf("Failed to set (%#v, %#v): %v", key, value, err)
	}

	if err := storage.Set(key, value, 0); err != nil {
		t.Errorf("Failed to set (%#v, %#v): %v", key, value, err)
	}

	if actual, _ := storage.Get(key); actual != value {
		t.Errorf("Expected %#v, got %#v", value, actual)
	}

	time.Sleep(2 * duration)

	if actual, _ := storage.Get(key); actual != value {
		t.Errorf("Expected %#v, got %#v", value, actual)
	}

	// Remove cancels the timer
	if err := storage.Set(key, value, duration); err != nil {
		t.Errorf("Failed to set (%#v, %#v): %v", key, value, err)
	}

	if err := storage.Remove(key); err != nil {
		t.Errorf("Failed to remove %#v: %v", key, err)
	}

	if err := storage.Set(key, value, 0); err != nil {
		t.Errorf("Failed to set (%#v, %#v): %v", key, value, err)
	}

	if actual, _ := storage.Get(key); actual != value {
		t.Errorf("Expected %#v, got %#v", value, actual)
	}

	time.Sleep(2 * duration)

	if actual, _ := storage.Get(key); actual != value {
		t.Errorf("Expected %#v, got %#v", value, actual)
	}
}

func TestNotFoundError(t *testing.T) {
	expected := "test"
	err := &NotFoundError{expected}

	if actual := err.Error(); actual != expected {
		t.Errorf("Expected %#v, got %#v", expected, actual)
	}
}

func TestTypeError(t *testing.T) {
	expected := "test"
	err := &TypeError{expected}

	if actual := err.Error(); actual != expected {
		t.Errorf("Expected %#v, got %#v", expected, actual)
	}
}
