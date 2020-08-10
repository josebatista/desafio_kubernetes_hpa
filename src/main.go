package main

import "math"

func squareRoot(value float64) float64 {
	return math.Sqrt(value)
}

func main() {

	var x = 0.0001

	for i := 0; i < 1000000; i++ {
		x += squareRoot(x)
	}

	println("Code.education Rocks!")

}
