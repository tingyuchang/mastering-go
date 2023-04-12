package main

import (
	"testing"
	"testing/quick"
)

var N = 1000000

func TestWithItSelf(t *testing.T) {
	condition := func(a, b Point2D) bool {
		return Add(a, b) == Add(b, a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})

	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
