package config

import (
	"log"
	"os"
	"strconv"
)

var DefaultRequestsPerSecond string = os.Getenv("MAX_REQUESTS_PER_SECOND")

func GetDefaultRequestsPerSecond() float64 {
	if DefaultRequestsPerSecond == "" {
		return 100
	}
	val, err := strconv.ParseFloat(DefaultRequestsPerSecond, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
