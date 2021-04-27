package apiclient

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/HimbeerserverDE/srp"
)

type ApiClient struct {
	Url       string // host + port
	AuthToken string
}

func NewApiClient(url string, u string, p string) *ApiClient {

	ApiObj := &ApiClient{
		Url: url,
	}

	A, a, err := srp.InitiateHandshake()
	if err != nil {
		panic(err)
	}
	A_b64 := base64.StdEncoding.EncodeToString(A)

	B_b64, s_b64 := SrpAToServer(url, u, A_b64)
	B, _ := base64.StdEncoding.DecodeString(B_b64)
	s, _ := base64.StdEncoding.DecodeString(s_b64)

	K, err := srp.CompleteHandshake(A, a, []byte(u), []byte(p), s, B)
	if err != nil {
		panic(err)
	}
	clientProof := srp.Hash(K)
	clientProof_b64 := base64.StdEncoding.EncodeToString(clientProof)
	fmt.Println(clientProof_b64)

	M2_b64 := M1ToServer(url, u, clientProof_b64, A_b64)
	M2, _ := base64.StdEncoding.DecodeString(M2_b64)

	if !bytes.Equal(M2, K) {
		panic("Doesn't server and client prooves don't match.")
	}

	return ApiObj
}
