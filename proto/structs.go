package proto

type EncryptedData struct {
	iv   []byte // aka nonce
	ad   []byte // authentication data
	data []byte // encrypted data
}
