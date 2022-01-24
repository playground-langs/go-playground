package gadget

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
}

type TapeRecorder struct {
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("recorder playing ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("recorder stop")
}

func (t TapeRecorder) Record() {
	fmt.Println("recorder recording")
}

func (t TapePlayer) Play(song string) {
	fmt.Println("tape playing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("tape stop")
}
