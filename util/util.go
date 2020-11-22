package util

import "fmt"

func Dbg(v interface{}) {
	fmt.Printf("%+v\n", v)
}
