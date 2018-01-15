package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/sparrc/go-ping"
)

func checkUp(server string) {
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

func main() {
	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		//run check on throwaway ec2 instance
		checkUp("13.59.137.165")
		//Poll every 12 seconds.
		time.Sleep(12000 * time.Millisecond)
	}

}
