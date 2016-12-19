// object_test.go
package main

import "testing"

func TestCollisionable(t *testing.T) {
	patterns := []struct {
		A      Collisionable
		B      Collisionable
		Result bool
	}{
		{NewCollisionableObject(0, 0, 10, 10, "#"), NewCollisionableObject(30, 30, 30, 30, "#"), false},
		{NewCollisionableObject(0, 0, 10, 10, "#"), NewCollisionableObject(5, 5, 10, 10, "#"), true},
		{NewCollisionableObject(0, 0, 10, 10, "#"), NewCollisionableObject(5, 5, 1, 1, "#"), true},
		{NewCollisionableObject(0, 0, 10, 10, "#"), NewCollisionableObject(5, 5, 0, 0, "#"), true},
		{NewCollisionableObject(1, 1, 1, 1, "#"), NewCollisionableObject(1, 1, 1, 1, "#"), true},
		{NewCollisionableObject(0, 0, 3, 3, "#"), NewCollisionableObject(4, 4, 2, 2, "#"), false},
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
func TestMovable(t *testing.T) {
	o := NewMovableObject(0, 0, 1, 1, "*", 1, 1)
	o.Next()
	o.Next()
	o.Next()

	if o.Point().X != 3 && o.Point().Y != 3 {
		t.Errorf("3,3!=%d,%d", o.Point().X, o.Point().Y)
	}
	o.Turn(HORIZONAL)
	o.Next()

	if o.Point().X != 4 && o.Point().Y != 2 {
		t.Errorf("4,2!=%d,%d", o.Point().X, o.Point().Y)
	}
	o.Move(-4, -2)
	if o.Point().X != 0 && o.Point().Y != 0 {
		t.Errorf("0,0!=%d,%d", o.Point().X, o.Point().Y)
	}
}
