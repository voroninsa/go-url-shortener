package common

import "testing"

func TestEncoder(t *testing.T) {
	expected := "e1a69a7f"
	actual := EncodeString("yandex.ru")

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}
