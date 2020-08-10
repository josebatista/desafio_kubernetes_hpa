package main

import "testing"

func TestSquareRoot(t *testing.T) {

	x := squareRoot(4)

	if 2 != x {
		t.Errorf("Esperava o valor %d, mas recebeu o valor %f", 2, x)
	}

}
