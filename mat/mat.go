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

// Returns center's offset vector
func (d Dim3) Offset() Vec3 {
	return Vec3{d.W / 2, d.H / 2, d.D / 2}
}
