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

type FBDebugTokenResp struct {
	AppID       string         `json:"app_id"`
	Application string         `json:"application"`
	ExpiresAt   int            `json:"expires_at"`
	Valid       bool           `json:"is_valid"`
	IssuedAt    int            `json:"issued_at"`
	Scopes      []string       `json:"scopes"`
	UserID      string         `json:"user_id"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	Error       *facebookError `json:"error"`
}

type FacebookSession struct {
	userID    string
	UserToken string
	AppToken  string
}

// HashPassword creates a hash of a given password
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	hashStr := string(hash)

	return hashStr, err
}

func (fbcred *FacebookSession) TokenValidation(fburl string) *http.Response {
	url := fmt.Sprintf("%sinput_token=%s&access_token=%s", fburl, fbcred.AppToken, fbcred.UserToken)
	req, _ := http.NewRequest("GET", url, nil)
	hresp, err := Client.Do(req)
	if err != nil {
		log.Fatal("Request not successful")
	}
	return hresp
}
