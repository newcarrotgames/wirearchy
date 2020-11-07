package plan

import (
	"regexp"
	"testing"
)

func Test_name(t *testing.T) {
	var re = regexp.MustCompile(`(?m)[A-Z][a-z]+[A-Z][a-z]+[A-Z][a-z]+`)

	tests := []struct {
		name string
	}{
		{"three words"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Name()
			if len(re.FindAllString(got, -1)) == 0 {
				t.Errorf("didn't get three capitalized words back, got: %s", got)
			}
		})
	}
}
