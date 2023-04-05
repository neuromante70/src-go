package main

import (
	// "fmt"
	"gadget"
)

type MusicPlayer interface {
	Play(string)
	Stop()
}

func TryOut(player MusicPlayer) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
