package main

import (
	"fmt"
	"mypkg"

	"github.com/headfirstgo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	playList(gadget.TapePlayer{}, mixtape)
	playList(gadget.TapeRecorder{}, mixtape)

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
	value := mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())

	s := Switch("off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}
