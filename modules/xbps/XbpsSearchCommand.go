package xbps

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func XbpsSearchCommand(args []string) {
	searchFlag := pflag.NewFlagSet("query", pflag.ExitOnError)
	searchFlag.Usage = func() {
		fmt.Println("Usage: ow query [option] <package>")
		fmt.Println()
		fmt.Println("Options")
		searchFlag.PrintDefaults()
	}

	help := searchFlag.BoolP("help", "h", false, "show helpful information")
	verbose := searchFlag.BoolP("verbose", "v", false, "show extra output")
	ignoreConfRepos := searchFlag.BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	memorySync := searchFlag.BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives")
	repositoryMode := searchFlag.BoolP("repository-mode", "R", false, "Enable repository mode. This mode explicitly looks for packages in repositories")
	// listPkgs := searchFlag.BoolP("list-pkgs", "-l", false, "List installed packages")
	// listRepos := searchFlag.BoolP("list-repos", "-L", false, "List registered repositories")
	// listHoldPkgs := searchFlag.BoolP("list-hold-pkgs", "-H", false, "List packages on hold")
	// listManualPkgs := searchFlag.BoolP("list-manual-pkgs", "", false, "List packages installed explicitly")
	// listOrphans := searchFlag.BoolP("list-orphans", "-O", false, "List package orphans")
	search := searchFlag.BoolP("search", "s", false, "Search for packages")

	if *help {
		searchFlag.Usage()
		return
	}
	searchFlag.Parse(args)
	pkgs := searchFlag.Args()

	if len(pkgs) == 0 || len(args) == 0 {
		fmt.Println("Invalid usage: ow search requires a package name")
		fmt.Println()
		searchFlag.Usage()
		return
	}
	xbpsArgs := []string{}

	if *ignoreConfRepos {
		xbpsArgs = append(xbpsArgs, "-i")
	}

	if *memorySync {
		xbpsArgs = append(xbpsArgs, "-M")
	}

	if *repositoryMode {
		xbpsArgs = append(xbpsArgs, "-R")
	}

	if *search {
		xbpsArgs = append(xbpsArgs, "-s")
	}

	if *verbose {
		fmt.Println("Running: xbps-query")
	}

	xbpsArgs = append(xbpsArgs, pkgs...)

	cmd := exec.Command("xbps-query", xbpsArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("xbps-query failed:", err)
		os.Exit(1)
	}
}
