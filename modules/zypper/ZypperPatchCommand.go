package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperPatchCommand(args []string) {
	patchFlag := pflag.NewFlagSet("patch", pflag.ExitOnError)

	verbose := patchFlag.BoolP("verbose", "v", false, "show extra output")
	useDetails := patchFlag.BoolP("details", "d", false, "show more details of the patch process")
	useReplaceFiles := patchFlag.BoolP("replacefiles", "r", false, "install the packages even if they would replace other files")

	patchFlag.Parse(args)

	zypperArgs := []string{"patch"}

	if *useDetails {
		zypperArgs = append(zypperArgs, "--details")
	}

	if *useReplaceFiles {
		zypperArgs = append(zypperArgs, "--replacefiles")
	}

	if *verbose {
		fmt.Println("Running: zypper patch")
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper patch failed:", err)
		os.Exit(1)
	}
}
