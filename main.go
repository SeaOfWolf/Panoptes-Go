package main

import (
	"fmt"

	"github.com/SeaOfWolf/Panoptes-Go/port"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Panoptes Go")

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	open := port.ScanPort("tcp","localhost", 443)
	if open.State == "Open" {
		fmt.Printf("Port 443 State: %s\n", green(open.State))
	} else {
		fmt.Printf("Port 443 State: %s\n", red(open.State))
	}

	results := port.InitialScan("localhost")
	fmt.Println("Scan Results:")
	for _, result := range results {
		if result.State == "Open" {
            fmt.Printf("Port %s: %s\n", result.Port, green(result.State))
        } else {
            fmt.Printf("Port %s: %s\n", result.Port, red(result.State))
        }
	}
}