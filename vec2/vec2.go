package vec2

type Vec2 struct {
	X, Y int
}

func New(x int, y int) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) InRange(x1, x2, y1, y2 int) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2
}

func (va Vec2) Add(vb Vec2) Vec2 {
	return Vec2{va.X + vb.X, va.Y + vb.Y}
}

func (v Vec2) Mul(s int) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}
