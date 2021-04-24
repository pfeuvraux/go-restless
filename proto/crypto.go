package proto

import (
	"github.com/mazen160/go-random"
	"golang.org/x/crypto/argon2"
)

func derive_key(passphrase []byte) []byte {

	salt, _ := random.Bytes(8)
	time := 3
	mem := 32 * 1024
	threads := 4
	outLen := 192

	return argon2.IDKey(passphrase, salt, uint32(time), uint32(mem), uint8(threads), uint32(outLen))
}
