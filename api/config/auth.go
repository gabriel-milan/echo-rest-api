package config

import "os"

var DefaultJWTSecret string = os.Getenv("JWT_SECRET")

func GetDefaultJWTSecret() string {
	if DefaultJWTSecret == "" {
		return "secret"
	}
	return DefaultJWTSecret
}
