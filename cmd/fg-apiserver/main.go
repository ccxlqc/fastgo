package main

import (
	"os"

	"github.com/onexstarck/fastgo/cmd/fg-apiserver/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewFastGOCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
