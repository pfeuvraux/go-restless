package apiclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kong/go-srp"
)

type ApiClient struct {
	Url       string // host + port
	AuthToken string
}

func SrpSaltFromServer(url string, username string) []byte {

	// get SRP salt before anything else
	type SrpInitRq struct {
		Username string `json:"username"`
	}

	jsonPayload, err := json.Marshal(&SrpInitRq{
		Username: username,
	})
	if err != nil {
		panic(err)
	}

	bufPayload := bytes.NewBuffer(jsonPayload)
	contentType := "application/json"
	endpoint := url + "/auth/login/init"

	resp, err := http.Post(endpoint, contentType, bufPayload)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resBlob, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("Status code wasn't 200.\n%s", string(resBlob))
	}

	type SrpInitRes struct {
		Salt string
	}

	res := &SrpInitRes{}
	err = json.Unmarshal(resBlob, &res)
	if err != nil {
		panic(err)
	}
	return []byte(res.Salt)
}

func NewApiClient(url string, u string, p string) *ApiClient {

	ApiObj := &ApiClient{
		Url: url,
	}

	srpParams := srp.GetParams(2048)
	srpSecret := srp.GenKey()

	srpSalt := SrpSaltFromServer(url, u)
	srpClient := srp.NewClient(srpParams, srpSalt, []byte(u), []byte(p), srpSecret)
	srpA := srpClient.ComputeA()

	return ApiObj
}
