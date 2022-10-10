package lib

import (
	"fmt"
	"log"
	"os/exec"
)

func Uninstall(packages []string) {
	cmd := exec.Command("apt", "remove")

	cmd.Args = append(cmd.Args, packages...)
	cmd.Args = append(cmd.Args, "-y")

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		RemoveRepository()
		log.Fatalf("%s[x]%s Failed to uninstall tools: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Removed successfully\n", ColorGreen, ColorNone)
}
