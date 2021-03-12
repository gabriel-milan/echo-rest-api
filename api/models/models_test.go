package models

import (
	"encoding/json"
	"testing"
)

func TestUser(t *testing.T) {
	user := User{Name: "Tester", Email: "test@test.com"}
	res, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	if string(res) != "{\"id\":0,\"name\":\"Tester\",\"email\":\"test@test.com\"}" {
		t.Fatalf("res was %s, should be ", res)
	}
}
