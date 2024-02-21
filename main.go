package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateClock(clock *widget.Label) {
	formattedTime := time.Now().Format("The time is: 03:04:05")
	clock.SetText(formattedTime)
}

func main() {

	// A simple app starts by creating an app instance with app.New()
	application := app.New()

	// then opening a window with app.NewWindow()
	window := application.NewWindow("Go Desktop App")

	window.SetTitle("Go Desktop App")

	clock := widget.NewLabel("Hello world")
	window.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateClock(clock)
		}
	}()
	// The app UI is then shown by calling ShowAndRun() on the window
	window.ShowAndRun()
}
