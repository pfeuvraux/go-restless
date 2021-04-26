package proto

import (
	"fmt"

	"github.com/mazen160/go-random"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

func Argon2idSaltLen() int {
	return 8
}

func DeriveKey(params ...[]byte) ([]byte, []byte) {

	passphrase := params[0]
	var salt []byte

	if len(params) > 1 {
		salt = params[1]
	} else {
		salt, _ = random.Bytes(Argon2idSaltLen())
	}

	time := 3
	mem := 32 * 1024
	threads := 4
	outLen := 32 // 256-bit key

	return argon2.IDKey(passphrase, salt, uint32(time), uint32(mem), uint8(threads), uint32(outLen)),
		salt
}

func Encrypt(data []byte, key []byte) *EncryptedData {
	cipher, err := chacha20poly1305.NewX(key)
	if err != nil {
		fmt.Printf("Error while init chacha")
		panic(err)
	}

	var nonce = make([]byte, cipher.NonceSize(), cipher.NonceSize()+len(data)+cipher.Overhead())
	ad, _ := random.Bytes(32)
	nonce, _ = random.Bytes(cipher.NonceSize())

	encryptedData := cipher.Seal(nonce, nonce, data, ad)
	return &EncryptedData{
		iv:   encryptedData[:cipher.NonceSize()],
		ad:   ad,
		data: encryptedData[cipher.NonceSize():],
	}
}
