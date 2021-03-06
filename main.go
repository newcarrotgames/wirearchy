package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/newcarrotgames/wirearchy/dis"
	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/gen"
	"github.com/newcarrotgames/wirearchy/mat"
)

func handler(w http.ResponseWriter, r *http.Request) {
	bb, _ := json.Marshal(getNewPoi())
	fmt.Fprintf(w, string(bb))
}

func rndMaterial() int {
	return gen.RND.Intn(11) + 2
}

var singleNodePlan = []byte(`{
	"BaseName": "",
	"Created": "2020-11-14T10:46:18.8482241-05:00",
	"Nodes": [
		{
			"W": 0.4,
			"H": 0.4,
			"D": 0.4,
			"X": 0,
			"Y": 0,
			"Z": 0,
			"Material": 2,
			"Nodes": [
				{
					"W": 0.8,
					"H": 0.8,
					"D": 0.8,
					"X": 1.2,
					"Y": -1.5,
					"Z": 0.9,
					"Material": 3
				},
				{
					"W": 0.9,
					"H": 0.9,
					"D": 0.9,
					"X": -1.2,
					"Y": 1.5,
					"Z": 0.9,
					"Material": 4
				},
				{
					"W": 0.7,
					"H": 0.7,
					"D": 0.7,
					"X": 1.2,
					"Y": -1.5,
					"Z": -0.9,
					"Material": 6
				}
			]
		}
	]
}`)

func getNewPoi() form.Form {
	var leader form.Form
	best := 0.0
	d := dis.CostDiscriminator{}
	size := mat.Dim3{W: 13, H: 13, D: 13}
	for i := 0; i < 1000; i++ {
		a := mat.NewArr3(mat.SqDim3(30))
		p := gen.RndBasePlan(gen.RND, gen.RndEvolution())
		f := p.Build(a, size)
		score := d.Discriminate(&f)
		if score > best {
			leader = f
			best = score
		}
	}
	t := form.Terrain(mat.SqDim3(32))
	t.Inset(leader.Arr3, mat.Vec3{})
	return t
}

func main() {
	// static web stuff
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// api endpoints
	http.HandleFunc("/api", handler)

	// run server
	log.Println("Listening on http://localhost:3000/...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
