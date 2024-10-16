package main

import "fmt"

type r struct {
	first string
}

func vS(r r, v string) r {
	r.first = v
	return r
}

func pS(r *r, v string) {
	r.first = v
}

func main() {
	r1 := r{
		first: "first name",
	}
	fmt.Println(r1)
	r1 = vS(r1, "second name")
	fmt.Println(r1)

	pS(&r1, "third name")
	fmt.Println(r1)
}
