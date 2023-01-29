package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/ubombar/live-pod-migration/pkg/migratord/daemon"
)

func main() {
	var address string
	var port int

	// TODO: Add a way to add docker client information
	flag.StringVar(&address, "address", "localhost", "Specifies the node address which the migratord is running.")
	flag.IntVar(&port, "port", 9213, "Specifies the port which the migratord is listening.")
	flag.Parse()

	logrus.Infof("Starting migratord on address %s and port %d\n", address, port)

	config := daemon.DaemonConfig{
		SelfAddress: address,
		SelfPort:    port,
		QueueSize:   64,
	}

	migratorDaemon := daemon.NewDaemon(&config)

	stopCh := make(chan interface{})

	err := migratorDaemon.Start()
	defer migratorDaemon.Stop()

	if err != nil {
		logrus.Errorln(err)
	}

	<-stopCh
}
