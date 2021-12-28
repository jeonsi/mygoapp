package main

import (
	"fmt"
	"greeting"
	"greeting/deutsch"
	"keyboard"
	"log"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()

	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
}
