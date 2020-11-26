package mat

import (
	"fmt"
	"testing"
)

var (
	d15 = Dim3{15, 15, 15}
	v0  = Vec3{}
	v1  = Vec3{1, 1, 1}
)

func aToS(a Arr3) string {
	return fmt.Sprintf("%v", a)
}

func TestArr3_Inset(t *testing.T) {
	tests := []struct {
		name string
		a    Arr3
		b    Arr3
		p    Vec3
		e    bool
	}{
		{"d2 true", NewArr3(d2), NewArr3(d2), v0, true},
		{"d2 false", NewArr3(d2), NewArr3(d2), v1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Set(v1, 1)
			tt.a.Inset(tt.b, tt.p)
			fmt.Println("a: " + aToS(tt.a))
			fmt.Println("b: " + aToS(tt.b))
			if (aToS(tt.a) == aToS(tt.b)) != tt.e {
				t.Errorf("test failed")
			}
		})
	}
}
