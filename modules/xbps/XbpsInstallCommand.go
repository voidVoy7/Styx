package xbps

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func XbpsInstallCommand(args []string) {
	installFlag := pflag.NewFlagSet("install", pflag.ExitOnError)
	installFlag.Usage = func() {
		fmt.Println("Usage: ow install [option] <package>")
		fmt.Println()
		fmt.Println("Options")
		installFlag.PrintDefaults()
	}

	help := installFlag.BoolP("help", "h", false, "show helpful information")
	verbose := installFlag.BoolP("verbose", "v", false, "show extra output")
	automatic := installFlag.BoolP("automatic", "A", false, "Set automatic installation mode")
	downloadOnly := installFlag.BoolP("download-only", "D", false, "Download packages and check integrity, nothing else")
	force := installFlag.BoolP("force", "f", false, "Force package re-installation, If specified twice, all files will be overwritten.")
	ignoreConfRepos := installFlag.BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	ignoreFileConflicts := installFlag.BoolP("ignore-file-conflicts", "I", false, "Ignore detected file conflicts")
	unpackOnly := installFlag.BoolP("unpack-only", "U", false, "Unpack packages in transaction, do not configure them")
	memorySync := installFlag.BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives")
	dryRun := installFlag.BoolP("dry-run", "n", false, "Dry-run mode")
	yes := installFlag.BoolP("yes", "y", false, "Assume yes to all questions")

	installFlag.Parse(args)

	if *help {
		installFlag.Usage()
		return
	}

	pkgs := installFlag.Args()

	xbpsArgs := []string{}

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow install requires a package name")
		fmt.Println()
		installFlag.Usage()
		return
	}

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

	if *yes {
		xbpsArgs = append(xbpsArgs, "-y")
	}

	if *dryRun {
		xbpsArgs = append(xbpsArgs, "-n")
	}

	if *ignoreFileConflicts {
		xbpsArgs = append(xbpsArgs, "-i")
	}

	if *verbose {
		fmt.Println("Running: xbps-install")
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
