package mat

import "math"

// Relative vector expressed as 0.1
type RelVec struct {
	X float64
	Y float64
	Z float64
}

func (r RelVec) Pos(d Dim3) Vec3 {
	// find d's origin
	return Vec3{
		X: int(math.Floor(r.X * float64(d.W/2.0))),
		Y: int(math.Floor(r.Y * float64(d.H/2.0))),
		Z: int(math.Floor(r.Z * float64(d.D/2.0))),
	}
}

type RelDim struct {
	W float64
	H float64
	D float64
}

func (r RelDim) Size(d Dim3) Dim3 {
	return Dim3{
		W: int(r.W * float64(d.W)),
		H: int(r.H * float64(d.H)),
		D: int(r.D * float64(d.D)),
	}
}
