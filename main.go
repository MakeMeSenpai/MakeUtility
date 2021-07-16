package main

import (
	"fmt"
	"time"
)

func main() {
	for (true)  {
		time.Sleep(1 * time.Minute)
		fmt.Print("hello world")
	}
}