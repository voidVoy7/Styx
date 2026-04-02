package xbps

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func XbpsRemoveCommand(args []string) {
	removeFlag := pflag.NewFlagSet("remove", pflag.ExitOnError)
	removeFlag.Usage = func() {
		fmt.Println("Usage: ow remove [option] <package>")
		fmt.Println()
		fmt.Println("Options")
		removeFlag.PrintDefaults()
	}

	help := removeFlag.BoolP("help", "h", false, "show helpful information")
	verbose := removeFlag.BoolP("verbose", "v", false, "show extra output")
	forceRevdeps := removeFlag.BoolP("force-revdeps", "F", false, "Force package removal even with revdeps or unresolved shared libraries")
	force := removeFlag.BoolP("force", "f", false, "Force package files removal")
	dryRun := removeFlag.BoolP("dry-run", "n", false, "Dry-run mode")
	cleanCache := removeFlag.BoolP("clean-cache", "O", false, "Remove outdated packages from the cache, If specified twice, also remove uninstalled packages")
	removeOrphans := removeFlag.BoolP("remove-orphans", "o", false, "Remove package orphans")
	recursive := removeFlag.BoolP("recursive", "R", false, "Recursively remove dependencies")
	yes := removeFlag.BoolP("yes", "y", false, "Assume yes to all questions")

	if *help {
		removeFlag.Usage()
		return
	}
	removeFlag.Parse(args)
	pkgs := removeFlag.Args()

	if len(pkgs) == 0 || len(args) == 0 {
		fmt.Println("Invalid usage: ow remove requires a package name")
		fmt.Println()
		removeFlag.Usage()
		return
	}
	xbpsArgs := []string{}

	if *verbose {
		fmt.Println("Running: xbps-remove")
	}

	if *yes {
		xbpsArgs = append(xbpsArgs, "-y")
	}

	if *forceRevdeps {
		xbpsArgs = append(xbpsArgs, "-F")
	}

	if *force {
		xbpsArgs = append(xbpsArgs, "-f")
	}

	if *dryRun {
		xbpsArgs = append(xbpsArgs, "-n")
	}

	if *cleanCache {
		xbpsArgs = append(xbpsArgs, "-O")
	}

	if *removeOrphans {
		xbpsArgs = append(xbpsArgs, "-o")
	}

	if *recursive {
		xbpsArgs = append(xbpsArgs, "-R")
	}

	xbpsArgs = append(xbpsArgs, pkgs...)

	cmd := exec.Command("xbps-remove", xbpsArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("xbps-remove failed:", err)
		os.Exit(1)
	}
}
