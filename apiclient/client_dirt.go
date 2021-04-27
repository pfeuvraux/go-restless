package apiclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

func SrpAToServer(url string, username string, A string) string {

	type SrpAToServerRq struct {
		Username string `json:"username"`
		A        string `json:"srp_params"`
	}

	jsonPayload, err := json.Marshal(&SrpAToServerRq{
		Username: username,
		A:        A,
	})
	if err != nil {
		panic(err)
	}

	bufPayload := bytes.NewBuffer(jsonPayload)
	contentType := "application/json"
	endpoint := url + "/auth/login/challenge"

	resp, err := http.Post(endpoint, contentType, bufPayload)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resBlob, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("Status code wasn't 200.\n%s", string(resBlob))
	}

	type SrpBFromServer struct {
		B string
	}

	res := &SrpBFromServer{}
	err = json.Unmarshal(resBlob, &res)
	if err != nil {
		panic(err)
	}

	return res.B

}

func M1ToServer(url string, username string, M1 string, srpA string, secret string) []byte {

	type SrpParamsRq struct {
		M1     string `json:"M1"`
		SrpA   string `json:"srpA"`
		Secret string `json:"secret"`
	}

	type SrpAToServerRq struct {
		Username  string      `json:"username"`
		SrpParams SrpParamsRq `json:"srp_params"`
	}

	srpParams := SrpParamsRq{
		M1:     M1,
		SrpA:   srpA,
		Secret: secret,
	}

	jsonPayload, err := json.Marshal(&SrpAToServerRq{
		Username:  username,
		SrpParams: srpParams,
	})
	if err != nil {
		panic(err)
	}

	bufPayload := bytes.NewBuffer(jsonPayload)
	contentType := "application/json"
	endpoint := url + "/auth/login/verify"

	resp, err := http.Post(endpoint, contentType, bufPayload)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resBlob, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("Status code wasn't 200.\n%s", string(resBlob))
	}

	type SrpM2FromServer struct {
		M2 string
	}

	res := &SrpM2FromServer{}
	err = json.Unmarshal(resBlob, &res)
	if err != nil {
		panic(err)
	}

	return []byte(res.M2)
}
