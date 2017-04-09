package main

import (
	"fmt"
	"net"

	"pault.ag/go/maclookup"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, device := range interfaces {
		vendor, err := maclookup.Lookup(device.HardwareAddr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", vendor)
	}
}
