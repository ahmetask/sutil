# Struct Util
*  Struct field getter/setter

# Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetask/sutil"
)

type C struct {
	C1 string `json:"c_1"`
	C2 int64 `json:"c_2"`
}
type B struct {
	C *C `json:"c"`
}
type A struct {
	B B `json:"b"`
}

func main() {
	v := &A{B: B{C: &C{
		C1: "test",
		C2: 1,
	}}}
	a := sutil.New(v).WithPath("B.C.C1", false).Get()

	fmt.Printf("get:%v\n", a.Value())


	res, err := sutil.New(v).WithValue("main").WithPath("b.c.c_1", true).Set()

	j, _ := json.Marshal(v)
	fmt.Printf("success:%v error:%v\nresult:%v\n", res, err, string(j))

	res, err = sutil.New(v).WithValue(B{}).WithPath("B", false).Set()

	j, _ = json.Marshal(v)
	fmt.Printf("success:%v error:%v\nresult:%v\n", res, err, string(j))
}
```
* Output
```text
get:test
success:true error:<nil>
result:{"b":{"c":{"c_1":"main","c_2":1}}}
success:true error:<nil>
result:{"b":{"c":null}}
````