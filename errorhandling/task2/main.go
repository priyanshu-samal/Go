package main

import( 
	"errors"
	"fmt"
)
func find(m map[string]int, name string) (int, error) {
	if m == nil {
		return 0, errors.New("empty map")
	}
	age,ok:=m[name]
	if !ok{
		return 0, errors.New("Not found")
	}
	return age,nil
}

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 10
	m["c"] = 10

	age, error := find(m, "b")
	if error!=nil{
         fmt.Println("error:", error)
		return
	}
	fmt.Println("age: ",age)
}