package main

// import "the_go_workshop/04-Complex_Types/interfaces/hfgo/gadget"

import "gadget"

// "github.com/headfirstgo/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{
		Batteries: "",
	}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
