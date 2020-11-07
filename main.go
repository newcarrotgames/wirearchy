package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/newcarrotgames/wirearchy/form"
	"github.com/newcarrotgames/wirearchy/gen"
	"github.com/newcarrotgames/wirearchy/mat"
)

func handler(w http.ResponseWriter, r *http.Request) {
	bb, _ := json.Marshal(getNewPoi())
	fmt.Fprintf(w, string(bb))
}

/*
'tiles/blocks/dirt.png',
'tiles/blocks/stonebrick.png',
'tiles/blocks/stone_slab_top.png',
'tiles/blocks/cobblestone.png',
'tiles/blocks/brick.png',
'tiles/blocks/nether_brick.png',
'tiles/blocks/planks_acacia.png',
'tiles/blocks/planks_big_oak.png',
'tiles/blocks/planks_jungle.png',
'tiles/blocks/planks_oak.png',
'tiles/blocks/planks_spruce.png',
'tiles/blocks/iron_block.png',
*/

func rndMaterial() int {
	return gen.RND.Intn(11) + 2
}

func getNewPoi() form.Room {
	// create poi and add random room
	poi := form.NewBase(mat.Dim3{W: 32, H: 32, D: 32}, 1)
	n := gen.RND.Intn(100) + 10
	var rooms []form.Room
	for i := 0; i < n; i++ {
		s := gen.RND.Intn(8) + 4
		pos := mat.Vec3{X: gen.RND.Intn(28) + 2, Y: s/2 + 1, Z: gen.RND.Intn(28) + 2}
		room := form.NewRoom(mat.Dim3{W: s, H: s, D: s}, rndMaterial())
		room.Vec3 = pos
		rooms = append(rooms, room)
	}

	// toss overlapping rooms for now
	for _, r := range rooms {
		empty := true
		r.Find(func(p mat.Vec3, val int) bool {
			q := r.Offset().Add(p.Add(r.Vec3))
			if !poi.Oob(q) {
				empty = poi.Get(q) == 0
			} else {
				empty = false
			}
			return empty
		})
		if empty {
			poi.Inset(r.Arr3, r.Vec3)
		}
	}
	return poi
}

func main() {
	// static web stuff
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// api endpoints
	http.HandleFunc("/api", handler)

	// run server
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
