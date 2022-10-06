package internal

import (
	"fmt"
	"log"
	"os/exec"
)

func Search(packageName string) {
	out, err := exec.Command("apt", "search", packageName).Output()

	if err != nil {
		RemoveRepository()
		log.Fatalf("%s[x]%s Failed to search: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Println(string(out))
}
