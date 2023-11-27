package tests

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	_cmd := exec.Command("sh", "-c", "pwd")
	_cmd.Stdout = os.Stdout
	_cmd.Stderr = os.Stderr
	if err := _cmd.Run(); err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}
