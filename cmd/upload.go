package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pfeuvraux/go-restless/internal/args"
	"github.com/pfeuvraux/go-restless/internal/config"
	"github.com/pfeuvraux/go-restless/proto"
)

func UploadHandler(u *args.Args) {

	config := config.Parse(u.ConfigPath)

	if _, err := os.Stat(u.Upload.Src); err != nil {
		log.Fatalln("File doesn't exist.")
	}

	fd, err := os.Open(u.Upload.Src)
	if err != nil {
		log.Fatalln("Couldn't open source file.")
	}
	defer fd.Close()

	fileBuffer, _ := ioutil.ReadAll(fd)
	uploader := proto.NewFileUploader(fileBuffer)

	url := "http://" + config.Api.Host + ":" + config.Api.Port
	uploader.Upload(config.Api.Username, config.Api.Password, url)
}
