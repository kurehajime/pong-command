// pong_test.go
package main

import (
	"testing"
)

func TestCollision(t *testing.T) {
	patterns := []struct {
		A      object
		B      object
		Result bool
	}{
		{NewObject(0, 0, 10, 10, "#"), NewObject(30, 30, 30, 30, "#"), false},
		{NewObject(0, 0, 10, 10, "#"), NewObject(5, 5, 10, 10, "#"), true},
		{NewObject(0, 0, 10, 10, "#"), NewObject(5, 5, 1, 1, "#"), true},
		{NewObject(0, 0, 10, 10, "#"), NewObject(5, 5, 0, 0, "#"), true},
		{NewObject(1, 1, 1, 1, "#"), NewObject(1, 1, 1, 1, "#"), true},
		{NewObject(0, 0, 3, 3, "#"), NewObject(4, 4, 2, 2, "#"), false},
	}

	for i := range patterns {
		if patterns[i].A.Collision(patterns[i].B) != patterns[i].Result {
			t.Errorf("Falt Pattern %d", i)
			continue
		}
		if patterns[i].B.Collision(patterns[i].A) != patterns[i].Result {
			t.Errorf("Falt Pattern %d", i)
			continue
		}
	}
}
