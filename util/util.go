package util

import (
	"encoding/json"
	"fmt"
	"math"
	"runtime/debug"
	"strings"
)

func Dbg(vs ...interface{}) {
	lines := strings.Split(string(debug.Stack()), "\n")
	fmt.Println("---------------------- STACK -------------------------")
	fmt.Println(strings.TrimSpace(lines[6]))
	fmt.Println("------------------------------------------------------")
	for _, v := range vs {
		switch v.(type) {
		case string:
			fmt.Println(v)
		default:
			fmt.Println(js(v))
		}
	}
	fmt.Println("------------------------------------------------------")
}

func Mulifi(i int, f float64) int {
	return int(math.Floor(float64(i) * f))
}

func js(v interface{}) string {
	bb, _ := json.Marshal(v)
	return string(bb)
}
