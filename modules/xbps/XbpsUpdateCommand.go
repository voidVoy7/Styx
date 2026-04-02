package xbps

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func XbpsUpdateCommand(args []string) {
	updateFlag := pflag.NewFlagSet("update", pflag.ExitOnError)
	updateFlag.Usage = func() {
		fmt.Println("Usage: ow update [option]")
		fmt.Println()
		fmt.Println("Options")
		updateFlag.PrintDefaults()
	}

	help := updateFlag.BoolP("help", "h", false, "show helpful information")
	verbose := updateFlag.BoolP("verbose", "v", false, "show extra output")
	automatic := updateFlag.BoolP("automatic", "A", false, "Set automatic installation mode")
	downloadOnly := updateFlag.BoolP("download-only", "D", false, "Download packages and check integrity, nothing else")
	force := updateFlag.BoolP("force", "f", false, "Force package re-installation, If specified twice, all files will be overwritten.")
	ignoreConfRepos := updateFlag.BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	ignoreFileConflicts := updateFlag.BoolP("ignore-file-conflicts", "I", false, "Ignore detected file conflicts")
	unpackOnly := updateFlag.BoolP("unpack-only", "U", false, "Unpack packages in transaction, do not configure them")
	memorySync := updateFlag.BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives")
	dryRun := updateFlag.BoolP("dry-run", "n", false, "Dry-run mode")

	updateFlag.Parse(args)

	if *help {
		updateFlag.Usage()
		return
	}

	pkgs := updateFlag.Args()

	xbpsArgs := []string{"-Su"}

	if *automatic {
		xbpsArgs = append(xbpsArgs, "-A")
	}

	if *downloadOnly {
		xbpsArgs = append(xbpsArgs, "-D")
	}

	if *force {
		xbpsArgs = append(xbpsArgs, "-f")
	}

	if *ignoreConfRepos {
		xbpsArgs = append(xbpsArgs, "-i")
	}

	if *ignoreFileConflicts {
		xbpsArgs = append(xbpsArgs, "-i")
	}

	if *unpackOnly {
		xbpsArgs = append(xbpsArgs, "-U")
	}

	if *memorySync {
		xbpsArgs = append(xbpsArgs, "-M")
	}

	if *dryRun {
		xbpsArgs = append(xbpsArgs, "-n")
	}

	if *ignoreFileConflicts {
		xbpsArgs = append(xbpsArgs, "-i")
	}

	if *verbose {
		fmt.Println("Running: xbps-install -Su")
	}

	xbpsArgs = append(xbpsArgs, pkgs...)

	cmd := exec.Command("xbps-install", xbpsArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("xbps-install failed:", err)
		os.Exit(1)
	}
}
