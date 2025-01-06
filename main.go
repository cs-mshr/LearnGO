package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	cir := Circle{Radius: 7}

	fmt.Println("Area of Rectangle:", calculateArea(rect))
	fmt.Println("Area of Circle:", calculateArea(cir))

	mysterBox := interface{}(10)
	describeValue(mysterBox)

	retrievedInt, ok := mysterBox.(int)
	if ok {
		fmt.Println("Retrieved int:", retrievedInt)
	} else {
		fmt.Println("Failed to retrieve int")
	}

}

func describeValue(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
	switch v := i.(type) {
	case int:
		fmt.Println("This is an int:", v)
	case string:
		fmt.Println("This is a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
