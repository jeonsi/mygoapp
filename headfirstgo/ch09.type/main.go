package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	carFuel += Liters(40.0).ToGallons()
	busFuel += Gallons(30.0).ToLiters()
	fmt.Printf("Car fuel: %.1f gallons\nBus fuel: %.1f liters\n", carFuel, busFuel)

	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()

	number := Number(4)
	number.Double()
	number.Display()
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

type Number int

func (n *Number) Display() {
	fmt.Println(*n)
}

func (n *Number) Double() {
	*n *= 2
}
