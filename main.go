package main

import (
	"log"

	argparser "github.com/alexflint/go-arg"
	"github.com/pfeuvraux/go-restless/cmd"
	"github.com/pfeuvraux/go-restless/internal/args"
)

func main() {
	cliArgs := args.NewArgs()
	argparser.MustParse(cliArgs)

	switch {
	case cliArgs.Register != nil:
		cmd.RegisterUser(cliArgs)
	case cliArgs.Upload != nil:
		cmd.UploadHandler(cliArgs)
	default:
		log.Fatal("Unrecognized command.")
	}

}
