package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "crypto/sha256" // https://github.com/Kong/go-srp/issues/1

	"github.com/kong/go-srp"
	"github.com/mazen160/go-random"
	"github.com/pfeuvraux/go-restless/internal/args"
	"github.com/pfeuvraux/go-restless/proto"
)

type RegisterUserAttributes struct {
	Username     string `json:"username"`
	Srp_verifier string `json:"srp_verifier"`
	Srp_salt     string `json:"srp_salt"`
	KekSalt      string `json:"master_salt"`
	CekKey       string `json:"content_encryption_key"`
}

func NewUserAttributes(username string) *RegisterUserAttributes {
	return &RegisterUserAttributes{
		Username: username,
	}
}

func (r *RegisterUserAttributes) SetAttributesFromBytes(s []byte, vkey []uint8) {
	r.Srp_salt = base64.StdEncoding.EncodeToString(s)
	r.Srp_verifier = base64.RawStdEncoding.EncodeToString(vkey)
}

func (r *RegisterUserAttributes) GenerateKeys(password string) {
	r.KekSalt, r.CekKey = proto.GenUserKeys(password)
}

func computeVerifier(username string, password string) ([]uint8, []byte) {

	salt, err := random.Bytes(4)
	if err != nil {
		log.Fatal("Error while generating salt...")
	}

	srp_params := srp.GetParams(2048)
	verifier := srp.ComputeVerifier(srp_params, salt, []byte(username), []byte(password))

	return verifier, salt
}

func MakeHttpRequest(user *RegisterUserAttributes, host string, port string) (string, int) {

	jsonPayload, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error while marshaling json...")
	}

	bufferedPayload := bytes.NewBuffer(jsonPayload)
	url := "http://" + host + ":" + port + "/auth/register"
	contentType := "application/json"

	resp, err := http.Post(url, contentType, bufferedPayload)
	if err != nil {
		log.Fatal("Something wrong happened when making POST request.")
	}
	defer resp.Body.Close()

	body_b, _ := ioutil.ReadAll(resp.Body)
	stringifiedBody := string(body_b)
	return stringifiedBody, resp.StatusCode
}

func RegisterUser(args *args.Args) {

	vkey, salt := computeVerifier(args.Register.Username, args.Register.Password)
	user := NewUserAttributes(args.Register.Username)
	user.SetAttributesFromBytes(salt, vkey)
	user.GenerateKeys(args.Register.Password)

	resp, statusCode := MakeHttpRequest(user, args.Register.Host, args.Register.Port)

	switch statusCode {
	case 201:
		println("User successfully registered.")
	case 409:
		log.Fatal("User already exists.")
	default:
		println(resp)
		log.Fatalf("Unhandlded status code %v.", statusCode)
	}
}
