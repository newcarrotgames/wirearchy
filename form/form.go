package form

import (
	"github.com/newcarrotgames/wirearchy/mat"
)

type Type string

type Form struct {
	mat.Arr3
	mat.Vec3
}

func Room(s mat.Dim3, material int) Form {
	r := Form{}
	r.Arr3.Dim3 = s
	r.Arr3.Dat = make([]int, s.Size())
	r.Each(func(p mat.Vec3, val int) {
		// set if block is part of an edge
		if p.X == 0 || p.Y == 0 || p.Z == 0 || p.X == r.W-1 || p.Y == r.H-1 || p.Z == r.D-1 {
			r.Set(p, material)
		}
	})
	return r
}

func Base(s mat.Dim3, material int) Form {
	r := Form{}
	r.Arr3.Dim3 = s
	r.Arr3.Dat = make([]int, s.Size())
	// todo: optimize
	r.Each(func(p mat.Vec3, val int) {
		if p.Y == 0 {
			r.Set(p, material)
		}
	})
	return r
}
