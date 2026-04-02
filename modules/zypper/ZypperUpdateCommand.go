package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperUpdateCommand(args []string) {
	updateFlag := pflag.NewFlagSet("update", pflag.ExitOnError)
	updateFlag.Usage = func() {
		fmt.Println("Usage: ow update [option]")
		fmt.Println()
		fmt.Println("Options:")
		updateFlag.PrintDefaults()
	}

	help := updateFlag.BoolP("help", "h", false, "show helpful information")
	verbose := updateFlag.BoolP("verbose", "v", false, "show extra output")
	useDetails := updateFlag.BoolP("details", "d", false, "show more details of the update process")

	updateFlag.Parse(args)

	zypperArgs := []string{"update"}

	if *useDetails {
		zypperArgs = append(zypperArgs, "--details")
	}

	if *verbose {
		fmt.Println("Running: zypper update")
	}

	if *help {
		updateFlag.Usage()
		return
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper update failed:", err)
		os.Exit(1)
	}
}
