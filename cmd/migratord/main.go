package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/ubombar/live-pod-migration/pkg/migratord/v1alpha1"
)

func main() {
	var address string
	var port int

	// TODO: Add a way to add docker client information
	flag.StringVar(&address, "address", "localhost", "Specifies the node address which the migratord is running.")
	flag.IntVar(&port, "port", 4545, "Specifies the port which the migratord is listening.")
	flag.Parse()

	logrus.Printf("Starting migratord on address %s and port %d\n", address, port)

	mig, _ := v1alpha1.NewMigratord(address, port)

	mig.Run()
}
