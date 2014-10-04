package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X,Y float64
}

//求开方
func (v *Vertex)Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v:=&Vertex{3,4}
	fmt.Println(v.Abs())
}


