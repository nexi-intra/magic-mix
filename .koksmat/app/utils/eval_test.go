package utils

import (
	"log"
	"testing"
)

func TestEval(t *testing.T) {
	s, _ := EvalTest("https://example.com/api/user/{{.UserID}}/details?age={{.Age}}&greeting={{.Greeting}}")
	log.Println(*s)
}
