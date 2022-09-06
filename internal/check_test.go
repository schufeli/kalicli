package internal_test

import (
	"testing"

	"github.com/schufeli/kalicli/internal"
)

func TestFileExists_FileExists(t *testing.T) {
	if !internal.CheckFileExists("/etc/apt/sources.list") {
		t.Fail()
	}
}

func TestFileExists_FileNotExists(t *testing.T) {
	if internal.CheckFileExists("/ets/apt/sources1234.list") {
		t.Fail()
	}
}
