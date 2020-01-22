package main

import (
	"flag"
	"github.com/weakish/gosugar"
	"os"
	"strings"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -r instead of -a")
	var compress *bool = flag.Bool("compress", false, "enable compression")
	flag.Parse()

	runRsync(*simple, *compress)
}

func runRsync(isSimple bool, enableCompress bool) {
	var simple int = gosugar.Btoi(isSimple)

	var opt string = [2]string{"-a", "-r"}[simple]
	var args []string = []string{"rsync", opt, "--partial", "--delete"}
	if enableCompress {
		args = append(args, "-z");
	}

	if len(os.Args) != 3 + simple + gosugar.Btoi(enableCompress) {
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
	os.Stderr.WriteString("Usage: r3c [-simple] [-compress] directory_1 directory_2\n")
	os.Exit(64)
}
