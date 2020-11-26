package mat

// Dim3
type Dim3 struct {
	W int
	H int
	D int
}

// Returns d's volume
func (d Dim3) Size() int {
	return d.W * d.H * d.D
}

// Checks if Oob
func (d Dim3) Oob(v Vec3) bool {
	return v.X >= d.W || v.Y >= d.H || v.Z >= d.D || v.IsNeg()
}

// Offset return's offset vector for use with inset
func (d Dim3) Offset() Vec3 {
	return Vec3{d.W / 2, d.H / 2, d.D / 2}
}

func SqDim3(s int) Dim3 {
	return Dim3{s, s, s}
}
