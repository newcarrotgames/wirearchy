package plan

import (
	"encoding/json"
	"time"

	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
	"github.com/newcarrotgames/wirearchy/util"
)

type Node struct {
	mat.RelDim
	mat.RelVec
	Root  *Node
	Nodes []*Node
}

func (n *Node) Add(a *Node) {
	n.Nodes = append(n.Nodes, a)
	a.Root = n
}

func (n *Node) Form(area mat.Dim3, material int) form.Form {
	return form.Room(n.Size(area), material)
}

type Plan struct {
	Name     string
	BaseName string // name of plan this one was derived from
	Created  time.Time
	Nodes    []*Node
}

func New(name string, nodes []*Node) *Plan {
	return &Plan{
		Name:    name,
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
	return p, err
}

// Returns Arr3 of "final" form according to Plan's relative structure definition
func (p *Plan) Build(size mat.Dim3) form.Form {
	a := mat.NewArr3(size)
	if p.Nodes != nil || len(p.Nodes) > 0 {
		for _, n := range p.Nodes {
			recurse(&a, n)
		}
	}
	return form.Form{
		PlanName: &p.Name,
		Arr3:     a,
		Vec3:     mat.Vec3{}}
}

func recurse(a *mat.Arr3, node *Node) {
	size := node.Size(a.Dim3)
	pos := node.RelVec.Pos(size)
	util.Dbg("------------------------------------------------------")
	util.Dbg(pos)
	util.Dbg("------------------------------------------------------")
	f := form.Room(size, material())
	a.Inset(f.Arr3, pos)
	if node.Nodes != nil {
		for _, n := range node.Nodes {
			recurse(a, n)
		}
	}
}

func material() int {
	return 2
}

func (p *Plan) Age() time.Duration {
	return time.Now().Sub(p.Created)
}
