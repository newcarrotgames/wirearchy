package mat

import (
	"reflect"
	"testing"
)

func TestRelVec_Pos(t *testing.T) {
	tests := []struct {
		name string
		v    RelVec
		d    Dim3
		want Vec3
	}{
		{"exact center", RelVec{}, Dim3{5, 5, 5}, Vec3{2, 2, 2}},
		{"exact center", RelVec{}, Dim3{9, 9, 9}, Vec3{4, 4, 4}},
		{"positive half", RelVec{0.5, 0.5, 0.5}, Dim3{10, 10, 10}, Vec3{7, 7, 7}},
		{"negative half", RelVec{-0.5, -0.5, -0.5}, Dim3{10, 10, 10}, Vec3{2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RelVec{
				X: tt.v.X,
				Y: tt.v.Y,
				Z: tt.v.Z,
			}
			if got := r.Pos(tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelVec.Pos() = %v, want %v", got, tt.want)
			}
		})
	}
}
