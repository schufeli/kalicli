package internal

import (
	"fmt"
	"log"
	"os/exec"
)

func Install(packages []string) {
	cmd := exec.Command("apt", "install")

	cmd.Args = append(cmd.Args, packages...)
	cmd.Args = append(cmd.Args, "-y")

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		RemoveRepository()
		log.Fatalf("%s[x]%s Failed to install tools: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Installation successful\n", ColorGreen, ColorNone)
}
