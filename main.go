package main

import (
	"fmt"
	"time"
	"github.com/gen2brain/beeep"
)

func main() {
	for (true)  {
		time.Sleep(20 * time.Second) // time.Minute)
		title := "20-20-20"
		message := "Take A break! Look at something 20ft away for at least 20 seconds"
		icon := "./eyes.png"
		err := beeep.Alert(title, message, icon)
		if err != nil {
			panic(err)
		}
		fmt.Print("Debug Message at: ", time.Now())
		time.Sleep(20 * time.Second)
	}
}