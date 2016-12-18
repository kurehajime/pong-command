// object.go
package main

//Alignment ------------------------------------------------
type Alignment int

const (
	//VERTICAL ,
	VERTICAL Alignment = iota
	//HORIZONAL ,
	HORIZONAL
)

//Objective interface------------------------------------------------
type Objective interface {
	Point() Point
	Size() Size
	Str() string
}

//Collisionable interface
type Collisionable interface {
	Objective
	Collision(Collisionable) bool
}

//Movable interface
type Movable interface {
	Objective
	Move(int, int)
	Next()
	Prev()
	Turn()
}

//CollisionableMovable interface
type CollisionableMovable interface {
	Objective
	Collision(Collisionable) bool
	Move(int, int)
	Next()
	Prev()
	Turn()
}

//Point :Object point------------------------------------------------
type Point struct {
	X int
	Y int
}

//Size : Object size------------------------------------------------
type Size struct {
	Width  int
	Height int
}

//Object : Object------------------------------------------------
type Object struct {
	point Point
	size  Size
	str   string
}

//NewObject : make new Object
func NewObject(x, y, w, h int, s string) Object {
	return Object{
		point: Point{X: x, Y: y},
		size:  Size{Width: w, Height: h},
		str:   s,
	}
}

//Point : get Point
func (o Object) Point() Point {
	return o.point
}

//Size : get Size
func (o Object) Size() Size {
	return o.size
}

//Str get Str
func (o Object) Str() string {
	return o.str
}

//MovableObject is movable------------------------------------------------
type MovableObject struct {
	Object
	Arrow Point
}

//NewMovableObject : NewMovableObject
func NewMovableObject(x, y, w, h int, s string, ax, ay int) MovableObject {
	return MovableObject{
		Object: Object{
			point: Point{X: x, Y: y},
			size:  Size{Width: w, Height: h},
			str:   s,
		},
		Arrow: Point{X: ax, Y: ay},
	}
}

//Move get Moved
func (o *MovableObject) Move(addX, addY int) {
	o.point.X += addX
	o.point.Y += addY
}

//Next get Moved
func (o *MovableObject) Next() {
	o.point.X += o.Arrow.X
	o.point.Y += o.Arrow.Y

}

//Prev get Move Cancel
func (o *MovableObject) Prev() {
	o.point.X -= o.Arrow.X
	o.point.Y -= o.Arrow.Y
}

//Turn Turn
func (o *MovableObject) Turn(a Alignment) {
	if a == VERTICAL {
		o.Arrow.X = -1 * o.Arrow.X
	} else {
		o.Arrow.Y = -1 * o.Arrow.Y
	}
}

//CollisionableObject is Collisionable------------------------------------------------
type CollisionableObject struct {
	Object
}

//NewCollisionableObject : get new NewCollisionableObject
func NewCollisionableObject(x, y, w, h int, s string) CollisionableObject {
	return CollisionableObject{
		Object{point: Point{X: x, Y: y},
			size: Size{Width: w, Height: h},
			str:  s,
		},
	}
}

//Collision : check Collision
func (o CollisionableObject) Collision(o2 Collisionable) bool {
	points1 := []Point{
		Point{X: o.Point().X, Y: o.Point().Y},
		Point{X: o.Point().X + o.Size().Width - 1, Y: o.Point().Y + o.Size().Height - 1},
		Point{X: o.Point().X + o.Size().Width - 1, Y: o.Point().Y},
		Point{X: o.Point().X, Y: o.Point().Y + o.Size().Height - 1},
	}
	points2 := []Point{
		Point{X: o2.Point().X, Y: o2.Point().Y},
		Point{X: o2.Point().X + o2.Size().Width - 1, Y: o2.Point().Y + o2.Size().Height - 1},
		Point{X: o2.Point().X + o2.Size().Width - 1, Y: o2.Point().Y},
		Point{X: o2.Point().X, Y: o2.Point().Y + o2.Size().Height - 1},
	}
	for i := range points2 {
		if o.Point().X <= points2[i].X && o.Point().X+o.Size().Width-1 >= points2[i].X &&
			o.Point().Y <= points2[i].Y && o.Point().Y+o.Size().Height-1 >= points2[i].Y {
			return true
		}
	}
	for i := range points1 {
		if o2.Point().X <= points1[i].X && o2.Point().X+o2.Size().Width-1 >= points1[i].X &&
			o2.Point().Y <= points1[i].Y && o2.Point().Y+o2.Size().Height-1 >= points1[i].Y {
			return true
		}
	}
	return false
}

//CollisionableMovableObject : Collisionable+MovableObject
type CollisionableMovableObject struct {
	MovableObject
}

//NewCollisionableMovableObject : NewCollisionableMovableObject
func NewCollisionableMovableObject(x, y, w, h int, s string, ax, ay int) CollisionableMovableObject {
	return CollisionableMovableObject{
		MovableObject: MovableObject{
			Object: Object{
				point: Point{X: x, Y: y},
				size:  Size{Width: w, Height: h},
				str:   s,
			},
			Arrow: Point{X: ax, Y: ay},
		},
	}
}

//Collision : check Collision
func (o CollisionableMovableObject) Collision(o2 Collisionable) bool {
	return NewCollisionableObject(o.Point().X, o.Point().Y, o.Size().Width, o.Size().Height, o.Str()).Collision(o2)
}
