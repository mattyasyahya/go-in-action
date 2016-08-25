package execute

import (
	"bufio"
	"fmt"
	"os/exec"
)

// RunJavaVersion run java -version
func RunJavaVersion() {
	path, err := exec.LookPath("java")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	command := exec.Command(path, "-jar", "demo-0.0.1-SNAPSHOT.jar")

	cmdReader, err := command.StdoutPipe()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println("=============")
			fmt.Println(scanner.Text())
			fmt.Println("=============")
		}
	}()

	err = command.Start()
	if err != nil {
		fmt.Println("error", err)
	}

	err = command.Wait()
	if err != nil {
		fmt.Println("error", err)
	}
}
