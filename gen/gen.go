package gen

import (
	"math"
	"math/rand"
	"time"

	"github.com/newcarrotgames/wirearchy/mat"
)

var (
	s1  = rand.NewSource(time.Now().UnixNano())
	RND = rand.New(s1)
)

// eventually, use adjusted weights rather than purely random values
func RndRelDim(rs *rand.Rand) mat.RelDim {
	w := rs.Float32()
	h := rs.Float32()
	d := rs.Float32()
	return mat.RelDim{W: w, H: h, D: d}
}

func RndRelPos(rs *rand.Rand) mat.RelVec {
	x := 1 - rs.Float32()*2
	y := 1 - rs.Float32()*2
	z := 1 - rs.Float32()*2
	return mat.RelVec{X: x, Y: y, Z: z}
}

// v should be a positive value between 0 and 1
func RndN(r int, v float64) int {
	n := float64(RND.Intn(r))
	o := float64(r) / v
	e := n * n / o
	e = math.Exp(float64(-1.0 * e))
	return int(float64(r) * e)
}
