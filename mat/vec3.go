package mat

import "github.com/newcarrotgames/wirearchy/util"

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vec3) IsNeg() bool {
	return v.X < 0 || v.Y < 0 || v.Z < 0
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{util.Mulifi(v.X, s), util.Mulifi(v.Y, s), util.Mulifi(v.Z, s)}
}
