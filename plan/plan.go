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

func New(nodes []*Node) *Plan {
	return &Plan{
		Name:    Name(),
		Created: time.Now(),
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
	if p.Name == "" {
		p.Name = Name()
	}
	return p, err
}

// Returns Arr3 of "final" form according to Plan's relative structure definition
func (p *Plan) Build(size mat.Dim3) form.Form {
	result := form.Form{
		PlanName: &p.Name,
		Arr3:     mat.NewArr3(size),
		Vec3:     mat.Vec3{},
	}
	for _, n := range p.Nodes {
		n.Build(&result, &result)
	}
	return result
}

func (p *Plan) Age() time.Duration {
	return time.Now().Sub(p.Created)
}
