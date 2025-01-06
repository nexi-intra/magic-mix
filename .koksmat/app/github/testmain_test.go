package github

import (
	"os"
	"testing"

	"github.com/magicbutton/magic-mix/utils"
)

func TestMain(m *testing.M) {
	utils.Setup("../.env")

	code := m.Run()

	os.Exit(code)
}
