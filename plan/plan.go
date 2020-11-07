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
	BaseName string
	Created  time.Time
	RootNode *Node
}

func New() *Plan {
	return &Plan{
		Name:    Name(),
		Created: time.Now(),
	}
}

func NewWithRoot(root *Node) Plan {
	return Plan{
		Name:     Name(),
		Created:  time.Now(),
		RootNode: root,
	}
}

// Returns Arr3 of "final" form according to structure defined by nodes.
func (p *Plan) Build(size mat.Dim3) *mat.Arr3 {
	return parseNodes(p.RootNode, size)
}

func parseNodes(node *Node, size mat.Dim3) *mat.Arr3 {
	a := mat.NewArr3(size)
	recurse(&a, node)
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
