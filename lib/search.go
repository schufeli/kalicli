package lib

import (
	"fmt"
	"log"
	"os/exec"
)

func Search(packagename string) {
	out, err := exec.Command("apt", "search", packagename).Output()

	if err != nil {
		RemoveRepository()
		log.Fatalf("%s[x]%s Failed to search: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Println(string(out))
}
