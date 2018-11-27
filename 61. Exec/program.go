/**
 * Implementation of classic exec function
 *
 * Note that Go does not offer a classic Unix fork function.
 * Usually this isn’t an issue though, since starting goroutines,
 * spawning processes, and exec’ing processes covers most use
 * cases for fork.
 */
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Go requires an absolute path to the binary we want to
	// execute, so we’ll use exec.LookPath to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use.
	// Here we just provide our current environment.
	env := os.Environ()

	// actual exec call
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
