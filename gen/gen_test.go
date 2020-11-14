package gen

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/newcarrotgames/wirearchy/mat"
)

func TestRndRelDim(t *testing.T) {
	d16 := mat.Dim3{16, 16, 16}
	type args struct {
		rs *rand.Rand
	}
	tests := []struct {
		name string
		args args
		want mat.RelDim
	}{
		{"first", args{RND}, mat.RelDim{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RndRelDim(tt.args.rs)
			fmt.Printf("got: %+v\n", got)
			d := got.Size(d16)
			fmt.Printf("d: %+v\n", d)
			want := int(float32(d16.D) * got.D)
			if d.D != want {
				fmt.Printf("incorrect depth: %+v - %+v\n", d, got)
			}
		})
	}
}

func TestRndRelPos(t *testing.T) {
	type args struct {
		rs *rand.Rand
	}
	tests := []struct {
		name string
		args args
	}{
		{"first", args{RND}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				got := RndRelPos(tt.args.rs)
				if got.X > 1 || got.X < -1 || got.Y > 1 || got.Y < -1 || got.Z > 1 || got.Z < -1 {
					t.Errorf("out of bounds: %+v", got)
				}
				fmt.Printf("%+v\n", got)
			}
		})
	}
}

func TestRndN(t *testing.T) {
	type args struct {
		r int
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test", args{16, 0.25}, 8.0},
		{"test", args{16, 0.50}, 8.0},
		{"test", args{16, 0.1}, 11.0},
	}
	for _, tt := range tests {
		total := 0
		rounds := 10000
		for i := 0; i < rounds; i++ {
			t.Run(tt.name, func(t *testing.T) {
				got := RndN(tt.args.r, tt.args.v)
				total += got
			})
		}
		avg := float64(total) / float64(rounds)
		if avg > tt.want {
			t.Errorf("gaussian average %f above %f for requested varibility of %f\n", avg, tt.want, tt.args.v)
		}
	}
}
