package apiclient

import (
	"encoding/base64"
	"fmt"

	"github.com/kong/go-srp"
)

type ApiClient struct {
	Url       string // host + port
	AuthToken string
}

func NewApiClient(url string, u string, p string) *ApiClient {

	ApiObj := &ApiClient{
		Url: url,
	}

	srpParams := srp.GetParams(4096)
	srpSecret := srp.GenKey()
	srpSecret_B64 := base64.StdEncoding.EncodeToString(srpSecret)
	fmt.Println(srpSecret_B64)

	srpSalt := SrpSaltFromServer(url, u)
	srpClient := srp.NewClient(srpParams, srpSalt, []byte(u), []byte(p), srpSecret)

	srpA := srpClient.ComputeA()
	srpA_B64 := base64.StdEncoding.EncodeToString(srpA)

	srpB_B64 := SrpAToServer(url, u, srpA_B64)
	srpB, _ := base64.StdEncoding.DecodeString(srpB_B64)

	srpClient.SetB(srpB)
	srpM1 := srpClient.ComputeM1()
	srpM1_B64 := base64.StdEncoding.EncodeToString(srpM1)

	M1ToServer(url, u, srpM1_B64, srpA_B64, srpSecret_B64)

	return ApiObj
}
