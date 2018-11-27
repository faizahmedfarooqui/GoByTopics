package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// Execute a command
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Here, we pipe data to the external process on its
	// stdin and collect the results from its stdout.
	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start the process,
	// write some input to it, read the resulting output, and finally
	// wait for the process to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// for an explicitly delineated command & arguement array
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
