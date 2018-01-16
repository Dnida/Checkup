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

	checkserv := [3]string{
		"www.rooksandkings.com",
		"www.google.com",
		"www.facebook.com",
	}

	for {
		clearscreen()
		//run check on ip/domains in the array
		checkup(checkserv[0])
		checkup(checkserv[1])
		checkup(checkserv[2])

		//Poll every 90 seconds.
		time.Sleep(90 * time.Second)
	}

}
