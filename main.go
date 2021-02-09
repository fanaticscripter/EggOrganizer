package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/fanaticscripter/EggOrganizer/cmd"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	cmd.Execute()
}
