package main

import (
	"os"

	"github.com/apex/log"
	handler "github.com/apex/log/handlers/cli"

	"github.com/wesleimp/gitignore/cmd/cli"
)

var (
	version = "v0.1.0"
)

func main() {
	log.SetHandler(handler.Default)

	err := cli.Execute(version, os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
