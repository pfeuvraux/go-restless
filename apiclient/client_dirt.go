package apiclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SrpSaltFromServer(url string, username string) string {

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
	return res.Salt
}

func SrpAToServer(url string, username string, A string) (string, string) {

	type SrpParamsRq struct {
		A string `json:"srpA"`
	}

	type SrpAToServerRq struct {
		Username  string      `json:"username"`
		SrpParams SrpParamsRq `json:"srp_params"`
	}

	srpParams := SrpParamsRq{
		A: A,
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
		s string
	}

	res := &SrpBFromServer{}
	err = json.Unmarshal(resBlob, &res)
	if err != nil {
		panic(err)
	}

	return res.B, res.s

}

func M1ToServer(url string, username string, M1 string, srpA string) string {

	type SrpParamsRq struct {
		M1   string `json:"M1"`
		SrpA string `json:"srpA"`
	}

	type SrpAToServerRq struct {
		Username  string      `json:"username"`
		SrpParams SrpParamsRq `json:"srp_params"`
	}

	srpParams := SrpParamsRq{
		M1:   M1,
		SrpA: srpA,
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

	return res.M2
}
