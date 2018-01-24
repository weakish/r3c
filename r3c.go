package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -r instead of -a")
	flag.Parse()

	var sourceDir string = os.Args[1]
	var destDir string = os.Args[2]
	if !strings.HasSuffix(sourceDir, "/") {
		sourceDir += "/"
	}

	defaultArgs := []string{"-a", "-z", "--partial", "--delete"}
	simpleArgs := []string{"-r", "-z", "--partial", "--delete"}
	var args []string
	if *simple {
		if len(os.Args) != 4 {
			printUsage()
		} else {
			args = append(simpleArgs, sourceDir, destDir)
		}
	} else {
		if len(os.Args) != 3 {
			printUsage()
		} else {
			args = append(defaultArgs, sourceDir, destDir)
		}
	}

	err := syscall.Exec(which("echo"), args, os.Environ())
	panicIf(err)

}
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
func which(command string) string {
	var commandPath string
	commandPath, err := exec.LookPath(command)
	panicIf(err)

	return commandPath
}
func printUsage() {
	os.Stderr.WriteString("Usage: r3c [--simple] directory_1 directory_2")
	os.Exit(64)
}
