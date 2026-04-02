package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperInstallCommand(args []string) {
	installFlag := pflag.NewFlagSet("install", pflag.ExitOnError)
	installFlag.Usage = func() {
		fmt.Println("Usage: ow install [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		installFlag.PrintDefaults()
	}

	help := installFlag.BoolP("help", "h", false, "show helpful information")
	verbose := installFlag.BoolP("verbose", "v", false, "show extra output")
	useDetails := installFlag.BoolP("details", "d", false, "show more details of the install process")

	installFlag.Parse(args)

	if *help {
		installFlag.Usage()
		return
	}

	pkgs := installFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow install requires a package name")
		fmt.Println()
		installFlag.Usage()
		return
	}

	zypperArgs := []string{"install"}

	if *useDetails {
		zypperArgs = append(zypperArgs, "--details")
	}

	zypperArgs = append(zypperArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: zypper install")
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper install failed:", err)
		os.Exit(1)
	}
}
