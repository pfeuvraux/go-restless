package proto

/*
#include "crypto.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

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

func NewFileUploader(file []byte) *File {
	return &File{
		chunks: SplitFile(file),
	}
}

func (f *File) Upload() {

	for i := 0; i < len(f.chunks); i++ {
		// b64chk := base64.StdEncoding.EncodeToString(f.chunks[i])
		var cEncryptedChunk *C.char = C.encrypt_chunk(C.CString(string(f.chunks[i])))
		encryptedChunk := C.GoString(cEncryptedChunk)
		C.free(unsafe.Pointer(cEncryptedChunk))

		fmt.Printf("%s", encryptedChunk)
	}
}
