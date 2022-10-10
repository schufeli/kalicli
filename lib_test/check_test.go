package lib_test

import (
	"testing"

	"github.com/schufeli/kalicli/lib"
)

func TestFileExists_FileExists(t *testing.T) {
	if !lib.CheckFileExists("/etc/apt/sources.list") {
		t.Fail()
	}
}

func TestFileExists_FileNotExists(t *testing.T) {
	if lib.CheckFileExists("/ets/apt/sources1234.list") {
		t.Fail()
	}
}
