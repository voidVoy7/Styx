package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperSearchCommand(args []string) {
	searchFlag := pflag.NewFlagSet("search", pflag.ExitOnError)
	searchFlag.Usage = func() {
		fmt.Println("Usage: bowser search [options] <package>")
		fmt.Println()
		fmt.Println("Options:")
		searchFlag.PrintDefaults()
	}

	help := searchFlag.BoolP("help", "h", false, "show helpful information")
	verbose := searchFlag.BoolP("verbose", "v", false, "show extra output")
	useMatchSubstrings := searchFlag.BoolP("match-substrings", "", false, "search for a match to partial words")
	useMatchWords := searchFlag.BoolP("match-words", "", false, "search for a match to only whole words")
	useMatchExact := searchFlag.BoolP("match-exact", "", false, "search for an exact match to the provided strings")
	useProvides := searchFlag.BoolP("provides", "p", false, "search for packages that provide the searched strings")
	useRequires := searchFlag.BoolP("requires", "r", false, "search for packages which require the searched strings")
	useRecommends := searchFlag.BoolP("recommends", "R", false, "search for packages which recommend the searched strings")
	useSupplements := searchFlag.BoolP("supplements", "s", false, "search for packages which supplement the searched strings")
	useConflicts := searchFlag.BoolP("conflicts", "c", false, "search for packages conflicting with the searched strings")
	useObsoletes := searchFlag.BoolP("obsoletes", "o", false, "search for packages which obsolete the searched strings")
	useSuggests := searchFlag.BoolP("suggests", "S", false, "search for packages which suggest the searched strings")
	useEnhances := searchFlag.BoolP("enhances", "e", false, "search for packages which enhance the searched strings")

	searchFlag.Parse(args)

	pkgs := searchFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: bowser search requires a package name")
		fmt.Println()
		searchFlag.Usage()
		return
	}

	zypperArgs := []string{"search"}

	if *useMatchSubstrings {
		zypperArgs = append(zypperArgs, "--match-substrings")
	}

	if *useMatchExact {
		zypperArgs = append(zypperArgs, "--match-exact")
	}

	if *useMatchWords {
		zypperArgs = append(zypperArgs, "--match-words")
	}

	if *useProvides {
		zypperArgs = append(zypperArgs, "--provides")
	}

	if *useRequires {
		zypperArgs = append(zypperArgs, "--requires")
	}

	if *useRecommends {
		zypperArgs = append(zypperArgs, "--recommends")
	}

	if *useSupplements {
		zypperArgs = append(zypperArgs, "--supplements")
	}

	if *useConflicts {
		zypperArgs = append(zypperArgs, "--conflicts")
	}

	if *useObsoletes {
		zypperArgs = append(zypperArgs, "--obsoletes")
	}

	if *useSuggests {
		zypperArgs = append(zypperArgs, "--suggests")
	}

	if *useEnhances {
		zypperArgs = append(zypperArgs, "--enhances")
	}

	zypperArgs = append(zypperArgs, pkgs...)

	if *help {
		searchFlag.Usage()
		return
	}

	if *verbose {
		fmt.Println("Running: zypper search")
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper install failed", err)
		os.Exit(1)
	}
}
