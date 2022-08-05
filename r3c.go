package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	"github.com/weakish/gosugar"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -r instead of -a")
	var compress *bool = flag.Bool("compress", false, "enable compression")
	var progress *bool = flag.Bool("progress", false, "show progress")
	var dry *bool = flag.Bool("dry", false, "dry run")
	var nodel *bool = flag.Bool("nodel", false, "do not delete extraneous files from dest dirs")
	flag.Parse()

	runRsync(*simple, *compress, *progress, *dry, !*nodel)
}

func runRsync(isSimple bool, enableCompress bool, showProgress bool, dryRun bool, delete bool) {
	var simple int = gosugar.Btoi(isSimple)

	var opt string = [2]string{"-a", "-r"}[simple]
	var args []string = []string{opt, "--partial"}
	if delete {
		args = append(args, "--delete")
	}
	if enableCompress {
		args = append(args, "-z")
	}
	if showProgress {
		args = append(args, "--progress")
	}

	var optLen = simple + gosugar.Btoi(enableCompress) + gosugar.Btoi(dryRun) + gosugar.Btoi(!delete)
	if len(os.Args) != 1+optLen+2 {
		printUsage()
	} else {
		var sourceDir string = os.Args[1+optLen]
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
		var destDir string = os.Args[2+optLen]
		args = append(args, sourceDir, destDir)
		cmd := exec.Command("rsync", args...)
		if dryRun {
			_, _ = os.Stderr.WriteString(cmd.String())
		} else {
			_ = cmd.Run()
		}
	}
}

func printUsage() {
	_, _ = os.Stderr.WriteString("Usage: r3c [-simple] [-compress] [-progress] [-nodel] [-dry] directory_1 directory_2\n")
	os.Exit(64)
}
