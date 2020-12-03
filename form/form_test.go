package form

import (
	"testing"

	"github.com/newcarrotgames/wirearchy/mat"
	"github.com/newcarrotgames/wirearchy/util"
)

func TestTerrain(t *testing.T) {
	tests := []struct {
		name string
		s    mat.Dim3
	}{
		{"test 1", mat.Dim3{16, 16, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Terrain(tt.s)
			util.Dbg(got)
		})
	}
}
