package dis

import (
	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
)

type Discriminator interface {
	Discriminate(f *form.Form) float64
}

type SimpleDiscriminator struct {
}

// simple rating: total blocks / total size
func (s SimpleDiscriminator) Discriminate(f *form.Form) float64 {
	score := 0
	f.Each(func(p mat.Vec3, val int) {
		if val > 0 {
			score++
		}
	})
	return float64(score) / float64(f.Size())
}
