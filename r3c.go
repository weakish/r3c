package main

import (
	"flag"
	"github.com/weakish/gosugar"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -r instead of -a")
	var compress *bool = flag.Bool("compress", false, "enable compression")
	var progress *bool = flag.Bool("progress", false, "show progress")
	flag.Parse()

	runRsync(*simple, *compress, *progress)
}

func runRsync(isSimple, enableCompress, showProgress bool) {
	var simple int = gosugar.Btoi(isSimple)

	var opt string = [2]string{"-a", "-r"}[simple]
	var args []string = []string{opt, "--partial", "--delete"}
	if enableCompress {
		args = append(args, "-z")
	}
	if showProgress {
		args = append(args, "--progress")
	}

	if len(os.Args) != 3+simple+gosugar.Btoi(enableCompress) {
		printUsage()
	} else {
		var sourceDir string = os.Args[1+simple]
		if !strings.HasSuffix(sourceDir, "/") {
			info, err := os.Stat(sourceDir)
			if err == nil {
				if info.IsDir() {
					sourceDir += "/"
				} else if info.Mode().IsRegular() {
					// pass
				} else {
					panic("r3c only supports regular file or directory: " + sourceDir)
				}
			} else {
				if os.IsNotExist(err) {
					gosugar.Expect(err, sourceDir+" not exist!")
				} else {
					panic(err)
				}
			}
		}
		var destDir string = os.Args[2+simple]
		args = append(args, sourceDir, destDir)
		cmd := exec.Command("rsync", args...)
		_ = cmd.Run()
	}
}

func printUsage() {
	_, _ = os.Stderr.WriteString("Usage: r3c [-simple] [-compress] directory_1 directory_2\n")
	os.Exit(64)
}
