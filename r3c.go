package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	"github.com/weakish/gosugar"
)

func main() {
	var simple *bool = flag.Bool("simple", false, "use -rt instead of -a")
	var compress *bool = flag.Bool("compress", false, "enable compression")
	var progress *bool = flag.Bool("progress", false, "show progress")
	var dry *bool = flag.Bool("dry", false, "dry run")
	var nodel *bool = flag.Bool("nodel", false, "do not delete extraneous files from dest dirs")
	flag.Parse()

	runRsync(*simple, *compress, *progress, *dry, !*nodel)
}

func runRsync(isSimple bool, enableCompress bool, showProgress bool, dryRun bool, delete bool) {
	var simple int = gosugar.Btoi(isSimple)

  var opt string = [2]string{"-a", "-rt"}[simple]
	var args []string = []string{opt, "--partial", "--sparse"}

	if delete {
		args = append(args, "--delete")
	}
	if enableCompress {
		args = append(args, "-z")
	}
	if showProgress {
		args = append(args, "--progress")
	}

	var optLen = simple + gosugar.Btoi(enableCompress) + gosugar.Btoi(showProgress) + gosugar.Btoi(dryRun) + gosugar.Btoi(!delete)
	if len(os.Args) != 1+optLen+2 { // r3c + options + source + dest
		printUsage()
	} else {
		var sourceDir string = os.Args[1+optLen] // following r3c + options
		if !strings.HasSuffix(sourceDir, "/") {
			info, err := os.Stat(sourceDir)
			if err == nil {
				if info.IsDir() {
					sourceDir += "/"
				} else if info.Mode().IsRegular() {
					// r3c is designed to be used with directories.
					// But it also tolerates a local file as the source.
				} else {
					panic("r3c only supports regular file or directory: " + sourceDir)
				}
			} else {
				if os.IsNotExist(err) {
					// Directory does not exist.
					// r3c assumes that the source is a remote directory and feels lucky.
					// If the source points to a remote file or if the source does not exist,
					// no sync will be performed, which is safe.
					sourceDir += "/"
				} else {
					panic(err)
				}
			}
		}
		var destDir string = os.Args[1+optLen+1] // following r3c + options + source
		args = append(args, sourceDir, destDir)
		cmd := exec.Command("rsync", args...)
		if dryRun {
			_, _ = os.Stderr.WriteString(cmd.String())
		} else {
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}
	}
}

func printUsage() {
	_, _ = os.Stderr.WriteString("Usage: r3c [-simple] [-compress] [-progress] [-nodel] [-dry] directory_1 directory_2\n")
	os.Exit(64)
}
