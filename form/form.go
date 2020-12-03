package form

import (
	"math"
	"math/rand"

	"github.com/newcarrotgames/wirearchy/mat"
	"github.com/newcarrotgames/wirearchy/tex"
	opensimplex "github.com/ojrac/opensimplex-go"
)

var noise = opensimplex.New(rand.Int63())

type Type string

type Form struct {
	PlanName *string
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

func Terrain(s mat.Dim3) Form {
	w, h, r := s.W, s.H, Empty(s)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)
			nh := noise.Eval2(xFloat, yFloat)
			fh := int(math.Floor(nh*(float64(h)/2) + float64(h)/2))
			r.Set(mat.Vec3{x, fh / 4, y}, int(tex.Grass))
		}
	}
	return r
}

func Empty(s mat.Dim3) Form {
	r := Form{}
	r.Arr3.Dim3 = s
	r.Arr3.Dat = make([]int, s.Size())
	return r
}
