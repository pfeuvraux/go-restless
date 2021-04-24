package proto

import "fmt"

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
	toto := SplitFile(file)
	return &File{
		Chunks: toto,
	}
}

func (f *File) Upload() {
	//key := derive_key([]byte(password))
	//fmt.Println(key)

	for i := 0; i < len(f.Chunks); i++ {
		fmt.Println("tes")
	}
}
