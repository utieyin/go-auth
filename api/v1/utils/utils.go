package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	//FacebookTokenURL is the base url for fb debug
	FacebookTokenURL = "https://graph.facebook.com/debug_token?"
)

var (
	//Client is an interface of type http client
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

type Data struct {
	AppID string `json:"app_id"`
	Valid bool   `json:"is_valid"`
}
type FBDebugTokenResp struct {
	Data Data `json:"data"`
}

type FacebookSession struct {
	UserAccessToken string
	FbAppID         string
	FbAppToken      string
}

var FbSess = &FacebookSession{FbAppID: os.Getenv("FbAppID"), FbAppToken: os.Getenv("SECRET_FACEBOOK_ACCESS_TOKEN")}

// HashPassword creates a hash of a given password
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	hashStr := string(hash)

	return hashStr, err
}

func (fbcred *FacebookSession) TokenValidation(fburl string) (bool, error) {
	url := fmt.Sprintf("%sinput_token=%s&access_token=%s", fburl, fbcred.UserAccessToken, fbcred.FbAppToken)
	req, _ := http.NewRequest("GET", url, nil)
	fmt.Println(url)
	req.Header.Set("Content-Type", "application/json")
	hresp, err := Client.Do(req)

	if err != nil {
		log.Fatal("Request not successful")
	}
	if hresp.StatusCode != http.StatusOK {
		log.Fatalf("Could not validate facebook token, response code: %d", hresp.StatusCode)
	}
	defer hresp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(hresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]interface{}
	json.Unmarshal(bodyBytes, &m)
	data := m["data"].(map[string]interface{})
	if data["is_valid"] != true {
		err := errors.New("Not valid token")
		return false, err
	}
	return true, err
}

func DebugToken(req *http.Request) (bool, error) {
	keys, ok := req.URL.Query()["tokens"]

	if ok {
		log.Println("Url Param ", keys)

	}

	FbSess.UserAccessToken = keys[0]

	resp, err := FbSess.TokenValidation(FacebookTokenURL)
	if err != nil {
		log.Fatal("Some error occurred")
	}
	return resp, nil
}
