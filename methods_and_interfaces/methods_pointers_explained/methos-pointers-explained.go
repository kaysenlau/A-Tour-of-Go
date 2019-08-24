package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Go interprets the statement v.Scale(5) as (&v).Scale(5) if the method's receiver is pointer, but won't transfer Scale(v,10)
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
