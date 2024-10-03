package luas

import "fmt"

type Shape interface {
	Area()
}

type Circle struct {
	JariJari float64
}

type Rectangle struct {
	Panjang float32
	Lebar   float32
}

func (c Circle) Area() {
	pi := 3.14
	fmt.Println(pi * c.JariJari * c.JariJari)
}
func (r Rectangle) Area() {
	fmt.Println(r.Panjang * r.Lebar)
}

func ViewArea(s Shape) {
	s.Area()
}
