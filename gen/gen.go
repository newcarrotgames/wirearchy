package gen

import (
	"math"
	"math/rand"
	"time"

	"github.com/newcarrotgames/wirearchy/dis"
	"github.com/newcarrotgames/wirearchy/mat"
	"github.com/newcarrotgames/wirearchy/plan"
)

var (
	s1  = rand.NewSource(time.Now().UnixNano())
	RND = rand.New(s1)
)

// eventually, use adjusted weights rather than purely random values
func RndRelDim(rs *rand.Rand) mat.RelDim {
	b := rs.Float64()/2 + 0.25
	w := b + rs.Float64()/4
	h := b + rs.Float64()/4
	d := b + rs.Float64()/4
	return mat.RelDim{W: w, H: h, D: d}
}

func RndRelPos(rs *rand.Rand) mat.RelVec {
	x := 1 - rs.Float64()*2
	y := 1 - rs.Float64()*2
	z := 1 - rs.Float64()*2
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

func RndNode(rs *rand.Rand, e Evolution) *plan.Node {
	n := &plan.Node{
		RelDim:   RndRelDim(rs),
		RelVec:   RndRelPos(rs),
		Material: rs.Intn(dis.NUMBER_OF_MATERIALS) + 2,
	}

	if rs.Intn(10) > 7 {
		ns := e.N() + 1
		nodes := make([]*plan.Node, ns)
		for i := 0; i < ns; i++ {
			node := RndNode(rs, e)
			nodes[i] = node
		}
		n.Nodes = nodes
	}

	return n
}

func RndBasePlan(rs *rand.Rand, e Evolution) *plan.Plan {
	p := plan.Blank()
	p.Name = Name()
	n := e.N() + 1
	nodes := make([]*plan.Node, n)
	for i := 0; i < n; i++ {
		node := RndNode(rs, e)
		nodes[i] = node
	}
	p.Nodes = nodes
	return p
}

type Evolution struct {
	RateOfChange float64
	Complexity   float64
}

func (e Evolution) N() int {
	b := int(math.Floor(4.0 * e.Complexity)) // how to determine this constant?
	return RND.Intn(b)
}

func (e Evolution) Dilute() Evolution {
	return Evolution{
		RateOfChange: e.RateOfChange * e.RateOfChange, // RateOfChange < 1
		Complexity:   e.Complexity / 2,
	}
}

func RndEvolution() Evolution {
	return Evolution{0.05, 1}
}
