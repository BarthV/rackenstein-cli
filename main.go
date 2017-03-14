package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/BarthV/rackenstein-cli/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Execute failed")
	}
}
