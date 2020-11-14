package mat

// Relative vector expressed as 0.1
type RelVec struct {
	X float32
	Y float32
	Z float32
}

func (r RelVec) Pos(d Dim3) Vec3 {
	return Vec3{
		X: int(r.X * float32(d.W)),
		Y: int(r.Y * float32(d.H)),
		Z: int(r.Z * float32(d.D)),
	}
}

type RelDim struct {
	W float32
	H float32
	D float32
}

func (r RelDim) Size(d Dim3) Dim3 {
	return Dim3{
		W: int(r.W * float32(d.W)),
		H: int(r.H * float32(d.H)),
		D: int(r.D * float32(d.D)),
	}
}
