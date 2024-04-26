package magicapp

import (
	"os"
	"testing"

	"github.com/magicbutton/magic-mix/utils"
)

func TestMain(m *testing.M) {
	cwd, _ := os.Getwd()

	utils.MakeEnvFile(cwd)
	utils.Setup("./.env-test")
	code := m.Run()

	os.Exit(code)
}
