package mat

import (
	"reflect"
	"testing"
)

var (
	d2   = Dim3{2, 2, 2}
	d8   = Dim3{8, 8, 8}
	d4x4 = Dim3{4, 4, 4}
)

func TestArr3_i(t *testing.T) {
	type fields struct {
		Dim3 Dim3
		dat  []int
	}
	type args struct {
		v Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// basic 2x2x2 checks
		{"2x2x2: change in x", fields{Dim3: d2}, args{Vec3{1, 0, 0}}, 1},
		{"2x2x2: change in y", fields{Dim3: d2}, args{Vec3{0, 1, 0}}, 4},
		{"2x2x2: change in z", fields{Dim3: d2}, args{Vec3{0, 0, 1}}, 2},

		// multiple dimensions
		{"2x2x2: x and y", fields{Dim3: d2}, args{Vec3{1, 1, 0}}, 5},
		{"2x2x2: y and z", fields{Dim3: d2}, args{Vec3{0, 1, 1}}, 6},
		{"2x2x2: x, y, and z", fields{Dim3: d2}, args{Vec3{1, 1, 1}}, 7},

		// 8x8x8
		{"8x8x8: change in x", fields{Dim3: d8}, args{Vec3{1, 0, 0}}, 1},
		{"8x8x8: change in y", fields{Dim3: d8}, args{Vec3{0, 1, 0}}, 64},
		{"8x8x8: change in z", fields{Dim3: d8}, args{Vec3{0, 0, 1}}, 8},
		{"8x8x8: x and y", fields{Dim3: d8}, args{Vec3{4, 3, 0}}, 64*3 + 4},
		{"8x8x8: y and z", fields{Dim3: d8}, args{Vec3{0, 5, 7}}, 5*64 + 7*8},
		{"8x8x8: x, y, and z", fields{Dim3: d8}, args{Vec3{4, 2, 5}}, 2*64 + 5*8 + 4},
		{"8x8x8: last index", fields{Dim3: d8}, args{Vec3{7, 7, 7}}, 7*64 + 7*8 + 7},

		// out of bounds tests (how do without checking bounds on get?)
		// this should return the value regardless because there's no bounds checks (and I'm not sure if I want them)
		{"oob: x is beyond width", fields{Dim3: d4x4}, args{Vec3{5, 0, 0}}, 5},
		{"oob: y is beyond height", fields{Dim3: d4x4}, args{Vec3{0, 5, 0}}, 80},
		{"oob: z is beyond depth", fields{Dim3: d4x4}, args{Vec3{0, 0, 5}}, 20},

		{"oob: x is negative", fields{Dim3: d4x4}, args{Vec3{-5, 0, 0}}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Arr3{
				Dim3: tt.fields.Dim3,
				Dat:  tt.fields.dat,
			}
			if got := a.i(tt.args.v); got != tt.want {
				t.Errorf("Arr3.i() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArr3_v(t *testing.T) {
	type fields struct {
		Dim3 Dim3
		Dat  []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec3
	}{
		{"2x2x2: zero", fields{Dim3: d2}, args{0}, Vec3{0, 0, 0}},
		{"2x2x2: x=1", fields{Dim3: d2}, args{1}, Vec3{1, 0, 0}},
		{"2x2x2: z=1", fields{Dim3: d2}, args{2}, Vec3{0, 0, 1}},
		{"2x2x2: y=1", fields{Dim3: d2}, args{4}, Vec3{0, 1, 0}},

		{"8x8x8: x=1", fields{Dim3: d8}, args{1}, Vec3{1, 0, 0}},
		{"8x8x8: z=1", fields{Dim3: d8}, args{8}, Vec3{0, 0, 1}},
		{"8x8x8: y=1", fields{Dim3: d8}, args{64}, Vec3{0, 1, 0}},

		{"8x8x8: {2, 5, 3}", fields{Dim3: d8}, args{5*64 + 3*8 + 2}, Vec3{2, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Arr3{
				Dim3: tt.fields.Dim3,
				Dat:  tt.fields.Dat,
			}
			if got := a.v(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arr3.v() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArr3_Get(t *testing.T) {
	type fields struct {
		s, i, v int
	}
	type args struct {
		v Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"2x2x2: get x", fields{2, 1, 1}, args{Vec3{1, 0, 0}}, 1},
		{"8x8x8: get x", fields{8, 1, 1}, args{Vec3{1, 0, 0}}, 1},
		{"2x2x2: get y", fields{2, 4, 1}, args{Vec3{0, 1, 0}}, 1},
		{"8x8x8: get y", fields{8, 8 * 8 * 3, 1}, args{Vec3{0, 3, 0}}, 1},
		{"2x2x2: get z", fields{2, 2, 1}, args{Vec3{0, 0, 1}}, 1},
		{"8x8x8: get z", fields{8, 8 * 3, 1}, args{Vec3{0, 0, 3}}, 1},

		{"2x2x2: get x, y, and z", fields{2, 7, 1}, args{Vec3{1, 1, 1}}, 1},
		{"8x8x8: get x, y, and z", fields{8, 8*8*3 + 8*5 + 6, 1}, args{Vec3{6, 3, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArr3(Dim3{tt.fields.s, tt.fields.s, tt.fields.s})
			a.Dat[tt.fields.i] = tt.fields.v
			if got := a.Get(tt.args.v); got != tt.want {
				t.Errorf("Arr3.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArr3_Set(t *testing.T) {
	type fields struct {
		s, i int
	}
	type args struct {
		v Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"2x2x2: set x", fields{2, 1}, args{Vec3{1, 0, 0}}, 1},
		{"8x8x8: set x", fields{8, 1}, args{Vec3{1, 0, 0}}, 1},
		{"2x2x2: set y", fields{2, 4}, args{Vec3{0, 1, 0}}, 1},
		{"8x8x8: set y", fields{8, 8 * 8 * 3}, args{Vec3{0, 3, 0}}, 1},
		{"2x2x2: set z", fields{2, 2}, args{Vec3{0, 0, 1}}, 1},
		{"8x8x8: set z", fields{8, 8 * 3}, args{Vec3{0, 0, 3}}, 1},

		{"2x2x2: set x, y, and z", fields{2, 7}, args{Vec3{1, 1, 1}}, 1},
		{"8x8x8: set x, y, and z", fields{8, 8*8*3 + 8*5 + 6}, args{Vec3{6, 3, 5}}, 1},

		{"oob: x is negative", fields{8, 8*8*3 + 8*5 + 6}, args{Vec3{6, 3, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArr3(Dim3{tt.fields.s, tt.fields.s, tt.fields.s})
			a.Set(tt.args.v, tt.want)
			if a.Dat[tt.fields.i] != tt.want {
				t.Errorf("Arr3.Get() = %v, want %v", a.Dat[tt.fields.i], tt.want)
			}
		})
	}
}

func TestDim3_Oob(t *testing.T) {
	type fields struct {
		w int
		h int
		d int
	}
	type args struct {
		v Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Oob false (v is 0)", fields{2, 3, 5}, args{Vec3{0, 0, 0}}, false},
		{"Oob false (v is max)", fields{2, 3, 5}, args{Vec3{1, 2, 4}}, false},
		{"Oob true (x oob)", fields{5, 3, 1}, args{Vec3{5, 2, 0}}, true},
		{"Oob true (y oob)", fields{5, 3, 1}, args{Vec3{1, 4, 0}}, true},
		{"Oob true (y neg)", fields{5, 3, 1}, args{Vec3{0, -5, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dim3{
				W: tt.fields.w,
				H: tt.fields.h,
				D: tt.fields.d,
			}
			if got := d.Oob(tt.args.v); got != tt.want {
				t.Errorf("Dim3.Oob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Add(t *testing.T) {
	type fields struct {
		v Vec3
	}
	type args struct {
		w Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec3
	}{
		{"basic addition", fields{Vec3{1, 1, 1}}, args{Vec3{2, 2, 2}}, Vec3{3, 3, 3}},
		{"also basic addition", fields{Vec3{32, 64, 16}}, args{Vec3{32, 64, 16}}, Vec3{64, 128, 32}},
		{"negative numbers", fields{Vec3{-32, -64, -16}}, args{Vec3{32, 64, 16}}, Vec3{0, 0, 0}},
		{"negative numbers", fields{Vec3{-32, -64, -16}}, args{Vec3{-32, -64, -16}}, Vec3{-64, -128, -32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.v.Add(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
