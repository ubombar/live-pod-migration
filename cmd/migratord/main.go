package main

import (
	"flag"

	"github.com/ubombar/live-pod-migration/pkg/migratord/v1alpha1"
)

func main() {
	var address string
	var port int

	flag.StringVar(&address, "address", "localohst", "Specifies the node address which the migratord is running.")
	flag.IntVar(&port, "port", 4545, "Specifies the port which the migratord is listening.")
	flag.Parse()

	mig, _ := v1alpha1.NewMigratord("localhost", 4545)

	mig.Run()
}
