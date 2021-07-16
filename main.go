package main

import (
	"fmt"
	"log"
	"time"
	"os/exec"
)

func Alert() {
	out, err := exec.Command("/bin/sh", "./notify.sh").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Look away!", out)
}

func main() {
	for (true)  {
		time.Sleep(20 * time.Minute)
		Alert()
		time.Sleep(20 * time.Second)
	}
}
