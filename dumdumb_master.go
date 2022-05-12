package main

import (
	"fmt"
	"net"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	sysInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, j := range sysInterfaces {
		k := &j
		address, err := k.Addrs()
		if err != nil {
			fmt.Printf(err.Error())
		}
		for _, add := range address {
			fmt.Println(k.Index, " ", k.Name, " ", add.String(), " network ", add.Network())
		}
	}
	IPs, err := net.LookupIP("localhost")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, ip := range IPs {
		fmt.Println(ip)
	}

}
