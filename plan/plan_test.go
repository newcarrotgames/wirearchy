package plan

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/newcarrotgames/wirearchy/mat"
)

var singleNodePlan = `{
	"Name": "singleNode",
	"BaseName": "",
	"Created": "2020-11-14T10:46:18.8482241-05:00",
	"Nodes": [
		{
			"W": 0.3,
			"H": 0.5,
			"D": 0.3,
			"X": 0,
			"Y": 0,
			"Z": 0,
			"Root": null,
			"Nodes": null
		}
	]
}`

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

func TestPlan_Build(t *testing.T) {
	type args struct {
		size mat.Dim3
	}
	tests := []struct {
		name string
		data []byte
		args args
		want *mat.Arr3
	}{
		{"build singleNodePlan", []byte(singleNodePlan), args{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := Decode(tt.data)
			got := p.Build(tt.args.size)
			fmt.Println(got)
		})
	}
}

func TestPlan_Encode(t *testing.T) {
	type fields struct {
		Name     string
		BaseName string
		Created  time.Time
		Nodes    []*Node
	}
	var nodes []*Node

	n := Node{
		RelDim: mat.RelDim{0.3, 0.5, 0.3},
		RelVec: mat.RelVec{0, 0, 0},
	}

	nodes = append(nodes, &n)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"first structure", fields{"singleNode", "", time.Now(), nodes}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plan{
				Name:     tt.fields.Name,
				BaseName: tt.fields.BaseName,
				Created:  tt.fields.Created,
				Nodes:    tt.fields.Nodes,
			}
			got, err := Encode(p)
			fmt.Println(string(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("Plan.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("Plan.Encode() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Plan.Encode() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestPlan_Decode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"decode singleNodePlan", args{[]byte(singleNodePlan)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Plan.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, _ := Encode(p)
			if !reflect.DeepEqual([]byte(singleNodePlan), got) {
				t.Errorf("Plan.Decode() result different")
			}
			fmt.Println(string(got))
		})
	}
}
