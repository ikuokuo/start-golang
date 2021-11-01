package main

import (
	"fmt"
	"net"
)

func main() {
	ifs, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}

	for _, ifi := range ifs {
		addr := ifi.HardwareAddr
		if addr != nil && (ifi.Flags&net.FlagUp > 0) {
			fmt.Println("------")
			fmt.Printf(" Name: %s\n", ifi.Name)
			fmt.Printf("  MAC: %v\n", ifi.HardwareAddr)
			fmt.Printf("Flags: %v\n", ifi.Flags)
		}
	}
}
