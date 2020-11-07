package plan

import (
	"time"

	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
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

// Returns Arr3 of "final" form according to Plan's relative structure definition
func (p *Plan) Build(size mat.Dim3) *mat.Arr3 {
	a := mat.NewArr3(size)
	if p.Nodes != nil || len(p.Nodes) > 0 {
		for _, n := range p.Nodes {
			recurse(&a, n)
		}
	}
	return &a
}

func recurse(a *mat.Arr3, node *Node) {
	if node.Nodes != nil {
		for _, n := range node.Nodes {
			recurse(a, n)
		}
	}
	form.Room(node.Size(a.Dim3), material())
}

func material() int {
	return 1
}

func (p *Plan) Age() time.Duration {
	return time.Now().Sub(p.Created)
}
