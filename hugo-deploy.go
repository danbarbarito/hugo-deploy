package main

import "fmt"
import "os/exec"
import "os"
import "log"

func main() {
	// Changes directory to example hugo project
	os.Chdir("/home/barbz/Projects/barbarito.me")
	// Checks if there is a config.toml
	_, err := os.Stat("config.toml")
	if os.IsNotExist(err) {
		log.Fatal("config.toml not found - You must run 'hugo-deploy' from the root of a hugo project")
	}
	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
