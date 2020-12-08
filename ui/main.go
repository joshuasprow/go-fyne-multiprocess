package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	log.SetPrefix("client: ")

	a := app.New()
	w := a.NewWindow("Hello")

	w.SetOnClosed(func() {})

	hello := widget.NewLabel("Hello Fyne!")

	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			reply, err := sayHello()
			if err != nil {
				log.Fatalf("sayHello: %v", err)
			}
			hello.SetText(reply)
		}),
	))

	w.ShowAndRun()
}
