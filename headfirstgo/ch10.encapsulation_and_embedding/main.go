package main

import (
	"calendar"
	"fmt"
	"geo"
	"log"
)

func main() {
	date := calendar.Date{}
	if err := date.SetYear(2019); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(5); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(27); err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	event := calendar.Event{}
	if err := event.SetTitle("Mom's birthday"); err != nil {
		log.Fatal(err)
	}
	if err := event.SetYear(2019); err != nil {
		log.Fatal(err)
	}
	if err := event.SetMonth(5); err != nil {
		log.Fatal(err)
	}
	if err := event.SetDay(27); err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Date.Day())

	value := calendar.MyType{}
	value.ExportedMethod()

	coordinates := geo.Coordinates{}
	if err := coordinates.SetLatitude(37.42); err != nil {
		log.Fatal(err)
	}
	if err := coordinates.SetLongitude(-122.08); err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

	location := geo.Landmark{}
	if err := location.SetName("The Googleplex"); err != nil {
		log.Fatal(err)
	}
	if err := location.SetLatitude(37.42); err != nil {
		log.Fatal(err)
	}
	if err := location.SetLongitude(-122.08); err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
