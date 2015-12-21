package main

//import "fmt"
import "os/exec"
import "os"
import "github.com/fatih/color"
import "time"

// Checks if the working directory is a Hugo project
func isHugoProject() bool {
	// Checks if there is a config.toml
	_, err := os.Stat("config.toml")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func updateBranch(branch string) {
	out, err := exec.Command("git", "pull", "origin", branch).Output()
	if err != nil {
		color.Red("Error: " + err.Error())
		os.Exit(1)
	}
	color.Yellow(string(out))
}

func commitBranch(branch string) {
	out, err := exec.Command("git", "add", "-A").Output()
	if err != nil {
		color.Red(string(out))
		os.Exit(1)
	}
	commitMsg := "HugoDeploy - " + time.Now().String()
	out, err = exec.Command("git", "commit", "-m", commitMsg).Output()
	if err != nil {
		color.Yellow(string(out))
	} else {
		color.Yellow("Successfully pushed changes to " + branch + " branch")
	}
	out, err = exec.Command("git", "push", "origin", branch).Output()
	if err != nil {
		color.Yellow(string(out))
	}
}

func runHugo() {
	out, err := exec.Command("hugo").Output()
	if err != nil {
		color.Red("Unable to locate Config file. Make sure you are in the root of a Hugo project")
		os.Exit(0)
	}
	color.Yellow(string(out))
}

func main() {
	// Runs hugo
	color.Green("Building Hugo Site...")	
	runHugo()	
	// Updates source branch
	color.Green("Updating Your Source Branch...")	
	updateBranch("source")
	// Commits changes to source branch
	color.Green("Commiting Changes To Your Source Branch...")
	commitBranch("source")
	os.Chdir("public")
	// Commits changes to source branch
	color.Green("Commiting Changes To Your Master Branch...")
	commitBranch("master")
	color.Green("Hugo site successfully deployed")
}
