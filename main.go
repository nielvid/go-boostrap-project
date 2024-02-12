package main

import (
	// "flag"

	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"os/exec"

)


func ShellCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
var projectName string
func main() {

	var name string
	args := os.Args
	name = args[len(args)-1]
	module := strings.Split(name, "/")
	if len(module) == 1 {
		projectName = module[0]
	} else {
		projectName = module[len(module)-1]
	}

	//create src directory inside the projectName directory
	err := os.MkdirAll(projectName+"/src", 0750)

	if err != nil && os.IsExist(err) {
		log.Fatal("folder already exist")

	}
	//create main.go file inside the projectName directory
	mainFile := fmt.Sprintf("%s%s", projectName, "/src/main.go")
	err = os.WriteFile(mainFile, []byte("package main\n import \"fmt\"\nfunc main(){\nfmt.Println(\"hello word\")\n}"), 0660)
	if err != nil {
		log.Fatal(err)
	}
    command := fmt.Sprintf("cd %v && git init && go mod init %v", projectName, projectName)
	out, errout, err := ShellCommand(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	fmt.Println(out)
	fmt.Println("--- stderr ---")
	fmt.Println(errout)

	fmt.Println("project created")
	fmt.Printf("cd into %v/src\n", projectName)
	fmt.Println("go run main.go")

}
