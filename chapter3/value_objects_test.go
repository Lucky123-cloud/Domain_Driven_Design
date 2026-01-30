package chapter3

import (
	"testing"
)

func TestValue_Object(t *testing.T) {
	a := NewPoint(1, 1)
	b := NewPoint(1, 1)

	if a != b {
		t.Fatal("a and b are not equal")
	}
}
