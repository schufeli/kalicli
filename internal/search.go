package internal

import (
	"fmt"
	"log"
	"os/exec"
)

func Search(packageName string) {
	out, err := exec.Command("apt", "search", packageName).Output()

	if err != nil {
		log.Fatal()
	}

	fmt.Println(string(out))
}
