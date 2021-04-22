package proto

/*
#include "crypto.h"
*/
import "C"
import "fmt"

var MAX_CHUNK_SIZE int = 1000

type File struct {
	chunks [][]byte
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

func NewFileHandler(file []byte) *File {

	return &File{
		chunks: SplitFile(file),
	}
}

func (f *File) Encrypt() {

	for i := 0; i < len(f.chunks); i++ {
		// b64chk := base64.StdEncoding.EncodeToString(f.chunks[i])
		cChunk := C.CString(string(f.chunks[i]))
		encrypted_chunk := C.encrypt_chunk(cChunk)
		fmt.Println(C.GoString(encrypted_chunk))
	}

}
