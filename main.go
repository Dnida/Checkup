package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/sparrc/go-ping"
)

func checkup(server string) {
	pinger, err := ping.NewPinger(server)
	if err != nil {
		panic(err)
	}

	pinger.Count = 1
	pinger.Timeout = 1000 * time.Millisecond

	pinger.Run()

	if pinger.PacketsRecv == 0 {
		fmt.Println(server, "is down or has packet loss.")
	}
}

func clearscreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func main() {
	for {
		clearscreen()
		//run check on ip address using goroutines to call function at once.
		go checkup("13.59.137.165")
		go checkup("13.59.137.165")
		go checkup("13.59.137.165")
		go checkup("13.59.137.165")

		//Poll every 2 minutes.
		time.Sleep(10000 * time.Millisecond)
	}

}
