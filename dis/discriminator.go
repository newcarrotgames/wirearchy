package dis

import (
	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
)

const NUMBER_OF_MATERIALS = 11

type Discriminator interface {
	Discriminate(f *form.Form) float64
}

type SizeDiscriminator struct {
	// config?
}

// simple rating: total blocks / total size
func (s SizeDiscriminator) Discriminate(f *form.Form) float64 {
	score := 0
	f.Each(func(p mat.Vec3, val int) {
		if val > 0 {
			score++
		}
	})
	return float64(score) / float64(f.Size())
}

type CostDiscriminator struct {
	// config?
}

// simple rating: total blocks / total size
func (c CostDiscriminator) Discriminate(f *form.Form) float64 {
	score := 0
	f.Each(func(p mat.Vec3, val int) {
		score += val
	})
	return float64(score) / (float64(f.Size() * NUMBER_OF_MATERIALS))
}
