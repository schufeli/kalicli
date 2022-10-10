package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RemoveRepository() {
	err := os.Remove(RepositoryFileLocation)
	if err != nil {
		log.Fatalf("%s[x]%s Failed to remove repository file -> %s\n", ColorRed, ColorNone, RepositoryFileLocation)
	}

	fmt.Printf("%s[-]%s Removed repository file -> %s \n", ColorGreen, ColorNone, RepositoryFileLocation)
}

func RemoveRepositoryKey() {
	_, err := exec.Command("gpg", "--delete-keys", "44C6513A8E4FB3D30875F758ED444FF07D8D0BF6").Output()

	if err != nil {
		log.Fatalf("%s[x]%s Failed removing repository key: %s\n", ColorRed, ColorNone, err)
	}
}

func RemoveRepositoryKeyFile() {
	err := os.Remove(KeyFileLocation)
	if err != nil {
		log.Fatalf("%s[x]%s Failed to remove repository key file -> %s\n", ColorRed, ColorNone, KeyFileLocation)
	}

	fmt.Printf("%s[-]%s Removed repository key file -> %s \n", ColorGreen, ColorNone, KeyFileLocation)
}
