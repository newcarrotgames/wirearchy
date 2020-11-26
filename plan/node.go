package plan

import (
	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/mat"
)

type Node struct {
	mat.RelDim
	mat.RelVec
	Root     *Node
	Material int
	Nodes    []*Node
}

func (n *Node) Add(a *Node) {
	n.Nodes = append(n.Nodes, a)
	a.Root = n
}

func (n *Node) Build(root *form.Form, parent *form.Form) {
	// convert relative dimensions and position to real values
	size := n.Size(parent.Dim3)
	pos := n.Pos(parent.Dim3).Add(parent.Vec3)
	// build the form (only rooms atm) from size and pos
	f := form.Room(size, n.Material)
	f.Vec3 = pos
	for _, n := range n.Nodes {
		n.Build(root, &f)
	}
	// place node's form in root form
	root.Inset(f.Arr3, pos)
}
