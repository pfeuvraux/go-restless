package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pfeuvraux/go-restless/internal/args"
	"github.com/pfeuvraux/go-restless/proto"
)

func UploadHandler(host string, port string, file *args.FileUpload) {
	if _, err := os.Stat(file.Src); err != nil {
		log.Fatalln("File doesn't exist.")
	}

	fd, err := os.Open(file.Src)
	if err != nil {
		log.Fatalln("Couldn't open source file.")
	}
	defer fd.Close()

	fileBuffer, _ := ioutil.ReadAll(fd)
	ciphertext := proto.NewFileHandler(fileBuffer)
	ciphertext.Encrypt()

}
