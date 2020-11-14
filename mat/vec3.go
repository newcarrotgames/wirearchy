package mat

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vec3) IsNeg() bool {
	return v.X < 0 || v.Y < 0 || v.Z < 0
}
