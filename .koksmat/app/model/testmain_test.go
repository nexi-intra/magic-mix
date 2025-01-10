package model

import (
	"os"
	"testing"

	"github.com/magicbutton/magic-mix/config"
)

func TestMain(m *testing.M) {
	config.Setup("../.env")
	code := m.Run()

	os.Exit(code)
}
