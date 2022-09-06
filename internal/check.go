package internal

import (
	"errors"
	"os"
	"os/exec"
)

func CheckKeyExists() bool {
	out, _ := exec.Command("gpg", "--list-keys", "|", "grep", "Kali Linux Repository").Output()
	return len(out) != 0
}

func CheckFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
