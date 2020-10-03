package config

import "os"

const (
	secretFacebookAccessToken = "SECRET_FACEBOOK_ACCESS_TOKEN"
)

var (
	facebookAccessToken = os.Getenv(secretFacebookAccessToken)
)

// GetFacebookAccessToken returns the GH access token
func GetFacebookAccessToken() string {
	return facebookAccessToken
}
