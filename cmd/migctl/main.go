package main

import "github.com/ubombar/live-pod-migration/pkg/migctl/cmd"

func main() {

	// var addrClient, addrServer string
	// var portClient, portServer int

	// // Default values are for debugging
	// flag.StringVar(&addrClient, "addrc", "localhost", "Client address")
	// flag.StringVar(&addrServer, "addrs", "localhost", "Server address")
	// flag.IntVar(&portClient, "portc", 4545, "Client address")
	// flag.IntVar(&portServer, "ports", 4546, "Server port")

	cmd.Execute()
}
