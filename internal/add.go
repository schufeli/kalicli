package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func AddRepository() {
	err := os.WriteFile(RepositoryFileLocation, []byte(RepositoryFileContent), 0644)
	if err != nil {
		log.Fatalf("%s[x]%s Failed to write to file: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Created repository file -> %s \n", ColorGreen, ColorNone, RepositoryFileLocation)
}

func AddRepositoryKey() {
	_, err := exec.Command("gpg", "--keyserver", "keyserver.ubuntu.com", "--recv-key", "ED444FF07D8D0BF6").Output()

	if err != nil {
		log.Fatalf("%s[x]%s Failed downloading repository key: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Added repository key\n", ColorGreen, ColorNone)
}

func ExportRepositoryKey() {
	_, err := exec.Command("gpg", "-o", KeyFileLocation, "--export", "ED444FF07D8D0BF6").Output()

	if err != nil {
		log.Fatalf("%s[x]%s Failed exporting repository key: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Exported repository key -> %s\n", ColorGreen, ColorNone, KeyFileLocation)
}

func UpdateRepositoryList() {
	cmd := exec.Command("apt-get", "update", "-o", "Dir::Etc::sourcelist=sources.list.d/kali.list", "-o", "Dir::Etc::sourceparts=\"-\"", "-o", "apt::Get::List-Cleanup=\"0\"")

	if err := cmd.Run(); err != nil {
		log.Fatalf("%s[x]%s Failed updating sources.list: %s\n", ColorRed, ColorNone, err)
	}

	fmt.Printf("%s[+]%s Updated sources.list \n", ColorGreen, ColorNone)
}
