package main

import (
	"fmt"
	"go-playground/examples/gadget"
)

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// TryOut type assertion
//如果不使用第二个返回值ok来接收，会在运行时出现异常
func TryOut(player gadget.Player) {
	player.Play("test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("player is not a TapeRecorder")
	}
}

func main() {
	mixtape := []string{
		"Jessie's Girl",
		"Whip It",
		"9 to 5",
	}
	var player gadget.Player
	player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	s := gadget.Switch("off")
	var t gadget.Toggleable
	t = &s
	t.Toggle()
	s1 := gadget.Switch1("on")
	t = s1
	t.Toggle()

	TryOut(gadget.TapePlayer{})
	TryOut(player)
}
