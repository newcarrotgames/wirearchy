package plan

import (
	"encoding/json"
	"time"

	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
)

type Plan struct {
	Name     string
	BaseName string // name of plan this one was derived from
	Created  time.Time
	Nodes    []*Node
}

func Blank() *Plan {
	return &Plan{
		Created: time.Now(),
	}
}

func New(nodes []*Node) *Plan {
	return &Plan{
		Created: time.Now(),
		Nodes:   nodes,
	}
}

func Encode(p Plan) ([]byte, error) {
	return json.MarshalIndent(p, "", "\t")
}

func Decode(data []byte) (Plan, error) {
	var p Plan
	err := json.Unmarshal(data, &p)
	if err != nil {
		return p, err
	}
	return p, err
}

// Returns Arr3 of "final" form according to Plan's relative structure definition
func (p *Plan) Build(a mat.Arr3, size mat.Dim3) form.Form {
	root := form.Form{
		PlanName: &p.Name,
		Arr3:     a,
		Vec3:     mat.Vec3{},
	}
	result := form.Form{
		PlanName: &p.Name,
		Arr3:     mat.NewArr3(size),
		Vec3:     mat.Vec3{},
	}
	for _, n := range p.Nodes {
		n.Build(&root, &result)
	}
	return root
}

func (p *Plan) Age() time.Duration {
	return time.Now().Sub(p.Created)
}
