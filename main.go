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
	case cliArgs.Login != nil:
		cmd.UserSignIn()
	case cliArgs.Upload != nil:
		cmd.UploadHandler()
	default:
		log.Fatal("Unrecognized command.")
	}

}
