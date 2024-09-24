package main

import "fmt"

func main() {
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine engine
		make   string
		model  string
		doors  int
		color  string
	}

	car1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Benz",
		model: "CLA",
		doors: 4,
		color: "Yellow",
	}

	car2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "Y",
		doors: 4,
		color: "White",
	}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.engine.electric)
	fmt.Println(car1.doors)

	fmt.Println(car2.engine.electric)
	fmt.Println(car1.color)
}
