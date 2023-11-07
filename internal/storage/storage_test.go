package storage

import "testing"

func TestNewStorage(t *testing.T) {
	expected := AppStorage
	actual := NewStorage("inmemory")

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}
