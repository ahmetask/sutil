package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetask/sutil"
)

type C struct {
	C1 string
	C2 int64
}
type B struct {
	C *C
}
type A struct {
	B B
}

func main() {
	v := &A{B: B{C: &C{
		C1: "test",
		C2: 1,
	}}}

	res, err := sutil.New(v).WithValue("main").WithPath("B.C.C1", false).Set()

	j, _ := json.Marshal(v)
	fmt.Printf("success:%v error:%v\nresult:%v", res, err, string(j))

	res, err = sutil.New(v).WithValue(B{}).WithPath("B", false).Set()

	j, _ = json.Marshal(v)
	fmt.Printf("success:%v error:%v\nresult:%v", res, err, string(j))
}
