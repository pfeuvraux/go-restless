package proto

import (
	"encoding/base64"
)

func GenUserKeys(password string) (string, string) {
	kek, saltKek := DeriveKey([]byte(password))
	cek, saltCek := DeriveKey(kek)

	cek = append(cek, saltCek...)

	encedCek := Encrypt(cek, kek)

	finalCek := append(encedCek.data, encedCek.iv...)
	finalCek = append(finalCek, encedCek.ad...)
	finalCek = append(finalCek, saltCek...)

	finalCekB64 := base64.StdEncoding.EncodeToString(finalCek)
	kekSaltB64 := base64.StdEncoding.EncodeToString(saltKek)
	return kekSaltB64, finalCekB64
}
