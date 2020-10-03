package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockClient struct{}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func TestUser(t *testing.T) {
	client := &MockClient{}
	json := `{"is_valid":"true","name":"uti"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	req, _ := http.NewRequest("GET", "/user", nil)
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error occurred")
	}
	fmt.Println(res.Body)

}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
