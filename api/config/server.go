package config

import "os"

var DefaultApiUrl string = os.Getenv("API_URL")

func GetDefaultApiUrl() string {
	if DefaultApiUrl == "" {
		return ":1323"
	}
	return DefaultApiUrl
}
