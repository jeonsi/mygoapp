package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func main() {
	AcceptAnything(1)
	fmt.Println(Gallons(12.09248342))
}
