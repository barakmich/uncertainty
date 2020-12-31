package main

import (
	"fmt"

	"github.com/barakmich/uncertainty"
)

func main() {
	x := uncertainty.NewGaussian(30.0, 1.0)
	y := uncertainty.NewGaussian(5.0, 1.0)
	z := uncertainty.Div(x, y)
	m := uncertainty.Materialize(z, 100000)
	for _, v := range m.Samples {
		fmt.Println(v)
	}
}
