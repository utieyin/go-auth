package utils

import (
	"fmt"
	"log"
	"net/http"

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

type facebookSession struct {
	userID    string
	userToken string
	appToken  string
}

// HashPassword creates a hash of a given password
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	hashStr := string(hash)

	return hashStr, err
}

func (fbcred *facebookSession) TokenValidation(s, fburl string, body []byte, r *http.Request) *http.Response {
	url := fmt.Sprintf("%sinput_token=%s&access_token=%s", fburl, fbcred.appToken, fbcred.userToken)
	req, _ := http.NewRequest("GET", url, nil)
	hresp, err := Client.Do(req)
	if err != nil {
		log.Fatal("Request not successful")
	}
	return hresp
}
