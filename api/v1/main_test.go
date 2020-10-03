package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"github.com/utieyin/go-auth/api/v1/utils"
)

type MockClient struct{}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// type FBDebugTokenResp struct {
// 	BaseResp
// 	Data struct {
// 		Error struct {
// 			Code    int    `json:"code"`
// 			Message string `json:"message"`
// 		} `json:"error"`
// 		AppID               string `json:"app_id"`
// 		Type                string `json:"type"`
// 		Application         string `json:"application"`
// 		DataAccessExpiresAt int64  `json:"data_access_expires_at"`
// 		ExpiresAt           int64  `json:"expires_at"`
// 		IsValid             bool   `json:"is_valid"`
// 		IssuedAt            int64  `json:"issued_at"`
// 		Metadata            struct {
// 			SSO string `json:"sso"`
// 		} `json:"metadata"`
// 		Scopes []string `json:"scopes"`
// 		UserId string   `json:"user_id"`
// 	} `json:"data"`
// }

type FBDebugTokenResp struct {
	IsValid bool
	Name    string
}

var utils.Client = &MockClient{}
func TestUser(t *testing.T) {
	fbresp, _ := json.Marshal(FBDebugTokenResp{IsValid: false})
	r := ioutil.NopCloser(bytes.NewBuffer([]byte(fbresp)))
	req, _ := http.NewRequest("GET", "/user", nil)
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	res, err := utils.Client.Do(req)
	defer res.Body.Close()
	if err != nil {
		t.Errorf("Error occurred")
	}

	expected := http.StatusOK
	actual := res.StatusCode
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]interface{}
	json.Unmarshal(bodyBytes, &m)
	err = VerifyToken(m)
	if err != nil {
		t.Errorf("Token not valid")
	}
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
func VerifyToken(s map[string]interface{}) error {
	boolStr := s["IsValid"]
	fmt.Println(boolStr)
	if boolStr != true {
		fmt.Println("Something wrong")
		err := errors.New("Not valid")
		return err
	}
	return nil
}
