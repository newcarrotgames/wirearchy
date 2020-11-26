package mat

import (
	"log"

	"github.com/newcarrotgames/wirearchy/util"
)

// Arr3 is the main type for storing arrays of voxels
type Arr3 struct {
	Dim3
	Dat []int
}

// Returns 2d array index given a 3d position within arr3
func (a Arr3) i(v Vec3) int {
	return v.Y*a.W*a.D + v.Z*a.W + v.X
}

// Reverse of Arr3.i()
func (a Arr3) v(i int) Vec3 {
	l := a.W * a.D
	y := i / l
	r := i % l
	z := r / a.W
	x := r % a.W
	return Vec3{x, y, z}
}

// Returns value within Arr3 at v
func (a Arr3) Get(v Vec3) int {
	return a.Dat[a.i(v)]
}

// Sets voxel at v to val
func (a *Arr3) Set(v Vec3, val int) {
	a.Dat[a.i(v)] = val
}

// Iterator func type for each method
type Arr3Itr func(Vec3, int)

// Helper method to iterate arr3
func (a Arr3) Each(itr Arr3Itr) {
	for i, v := range a.Dat {
		itr(a.v(i), v)
	}
}

// Iterator func type for find method
type Arr3ItrBool func(Vec3, int) bool

// Helper method to iterate arr3
func (a Arr3) Find(itr Arr3ItrBool) {
	for i, v := range a.Dat {
		if !itr(a.v(i), v) {
			break
		}
	}
}

// Returns a blank arr3 of size d
func NewArr3(d Dim3) Arr3 {
	return Arr3{d, make([]int, d.Size())}
}

// Changes behavior of the blend method
type BlendMode uint8

// Blending mode contants
const (
	OVERWRITE BlendMode = 1 // overwrite the current value
	IGNORE    BlendMode = 2 // only write when current value is 0
	AVERAGE   BlendMode = 3 // writes the average of old and new values
)

// Writes b into a at pos p (p is the position of b's origin)
func (a *Arr3) Inset(b Arr3, p Vec3) { // todo: blend mode
	// convert p to the "real" position
	pos := p.Add(a.Offset().Sub(b.Offset()))
	pos.Y = 0
	util.Dbg("real pos: ", pos)
	a.blend(b, pos, OVERWRITE)
}

func (a *Arr3) blend(b Arr3, p Vec3, mode BlendMode) {
	b.Each(func(v Vec3, bVal int) {
		if bVal == 0 {
			return
		}
		pos := v.Add(p)
		if !a.Oob(pos) {
			if mode == OVERWRITE {
				a.Set(pos, bVal)
			} else {
				aVal := a.Get(pos)
				if mode == IGNORE && aVal == 0 {
					a.Set(pos, bVal)
				} else if mode == AVERAGE {
					a.Set(pos, (aVal+bVal)/2)
				} else {
					log.Printf("unknown blend mode: %d", mode)
				}
			}
		}
	})
}
