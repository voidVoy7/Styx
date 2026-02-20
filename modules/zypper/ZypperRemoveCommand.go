package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperRemoveCommand(args []string) {
	removeFlag := pflag.NewFlagSet("remove", pflag.ExitOnError)
	removeFlag.Usage = func() {
		fmt.Println("Usage: styx remove [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		removeFlag.PrintDefaults()
	}

	help := removeFlag.BoolP("help", "h", false, "show helpful information")
	verbose := removeFlag.BoolP("verbose", "v", false, "show extra output")
	useDetails := removeFlag.BoolP("details", "d", false, "show more details of the remove process")

	removeFlag.Parse(args)

	if *help {
		removeFlag.Usage()
		return
	}

	pkgs := removeFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: styx remove requires a package name")
		fmt.Println()
		removeFlag.Usage()
		return
	}

	zypperArgs := []string{"remove"}

	if *useDetails {
		zypperArgs = append(zypperArgs, "--details")
	}

	zypperArgs = append(zypperArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: zypper remove")
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper remove failed:", err)
		os.Exit(1)
	}
}
