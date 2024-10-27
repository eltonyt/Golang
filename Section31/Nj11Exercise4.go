package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("match error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	val := -10.23
	root, err := sqrt(val)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Square root of %v is %v. \n", val, root)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := sqrtError{
			lat:  "50.2289 N",
			long: "90.4656 W",
			err:  errors.New("You cannot have square root on a negative number."),
		}
		return f, e
	}
	return math.Sqrt(f), nil
}
