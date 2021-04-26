package proto

import (
	"encoding/base64"
	"fmt"
)

var MAX_CHUNK_SIZE int = 1000

type File struct {
	Chunks [][]byte
}

func SplitFile(file []byte) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(file)/MAX_CHUNK_SIZE)

	for len(file) >= MAX_CHUNK_SIZE {
		chunk, file = file[:MAX_CHUNK_SIZE], file[MAX_CHUNK_SIZE:]
		chunks = append(chunks, chunk)
	}
	if len(file) > 0 {
		chunks = append(chunks, file)
	}
	return chunks
}

func NewFileUploader(file []byte) *File {
	return &File{
		Chunks: SplitFile(file),
	}
}

func (f *File) Upload(username string, passwd string, url string) {
	fileEncryptionKey, _ := DeriveKey([]byte(passwd))

	for i := 0; i < len(f.Chunks); i++ {
		chunkEncryptionKey, _ := DeriveKey(fileEncryptionKey)

		encedChunk := Encrypt(f.Chunks[i], chunkEncryptionKey)
		fmt.Println(base64.StdEncoding.EncodeToString(encedChunk.data))
	}
}
