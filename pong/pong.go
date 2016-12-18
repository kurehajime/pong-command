// pong.go
package main

type Point struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Object interface {
	Point() Point
	Size() Size
	Collision(Object) bool
	Next(int, int) Object
	Str() string
}

type object struct {
	point Point
	size  Size
	str   string
}

func NewObject(x, y, w, h int, s string) object {
	return object{
		point: Point{X: x, Y: y},
		size:  Size{Width: w, Height: h},
		str:   s,
	}
}

func (o object) Point() Point {
	return o.point
}

func (o object) Size() Size {
	return o.size
}

func (o1 object) Collision(o2 Object) bool {
	points1 := []Point{
		Point{X: o1.Point().X, Y: o1.Point().Y},
		Point{X: o1.Point().X + o1.Size().Width - 1, Y: o1.Point().Y + o1.Size().Height - 1},
		Point{X: o1.Point().X + o1.Size().Width - 1, Y: o1.Point().Y},
		Point{X: o1.Point().X, Y: o1.Point().Y + o1.Size().Height - 1},
	}
	points2 := []Point{
		Point{X: o2.Point().X, Y: o2.Point().Y},
		Point{X: o2.Point().X + o2.Size().Width - 1, Y: o2.Point().Y + o2.Size().Height - 1},
		Point{X: o2.Point().X + o2.Size().Width - 1, Y: o2.Point().Y},
		Point{X: o2.Point().X, Y: o2.Point().Y + o2.Size().Height - 1},
	}
	for i := range points2 {
		if o1.Point().X <= points2[i].X && o1.Point().X+o1.Size().Width-1 >= points2[i].X &&
			o1.Point().Y <= points2[i].Y && o1.Point().Y+o1.Size().Height-1 >= points2[i].Y {
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

func (o object) Next(addX, addY int) Object {
	o.point.X += addX
	o.point.Y += addY
	return o
}

func (o object) Str() string {
	return o.str
}
