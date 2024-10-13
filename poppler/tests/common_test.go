package tests

import (
	"os"
	"testing"
)

var (
	cwd string
	err error
)

func setup() {
	os.Chdir("../testpdfs/")
	cwd, err = os.Getwd()
	if err != nil {
		panic("Setup Failed!")
	}
}

func TestMain(m *testing.M) {
	setup()

	m.Run()
}
