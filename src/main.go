package main

import (
	"fmt"
	"math"
	"net/http"
)

func squareRoot(value float64) float64 {
	return math.Sqrt(value)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var x = 0.0001

	for i := 0; i < 1000000; i++ {
		x += squareRoot(x)
	}

	fmt.Fprintf(w, "Code.education Rocks!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
