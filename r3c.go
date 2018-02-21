package main

import (
	"flag"
	"github.com/weakish/gosugar"
	"os"
	"strings"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -r instead of -a")
	flag.Parse()

	runRsync(*simple)
}

func runRsync(isSimple bool) {
	var simple int = gosugar.Btoi(isSimple)

	var opt string = [2]string{"-a", "-r"}[simple]
	var args []string = []string{"rsync", opt, "-z", "--partial", "--delete"}

	if len(os.Args) != 3 + simple {
		printUsage()
	} else {
		var sourceDir string = os.Args[1 + simple]
		if !strings.HasSuffix(sourceDir, "/") {
			sourceDir += "/"
		}
		var destDir string = os.Args[2 + simple]
		args = append(args, sourceDir, destDir)
		gosugar.Exec(args)
	}
}

func printUsage() {
	os.Stderr.WriteString("Usage: r3c [--simple] directory_1 directory_2\n")
	os.Exit(64)
}
