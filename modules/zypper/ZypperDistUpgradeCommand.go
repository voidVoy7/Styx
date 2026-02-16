package zypper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func ZypperDistUpgradeCommand(args []string) {
	distFlag := pflag.NewFlagSet("upgrade", pflag.ExitOnError)
	distFlag.Usage = func() {
		fmt.Println("Usage: bowser upgrade [options]")
		fmt.Println()
		fmt.Println("Options:")
		distFlag.PrintDefaults()
	}

	help := distFlag.BoolP("help", "h", false, "show helpful information")
	verbose := distFlag.BoolP("verbose", "v", false, "show extra output")
	useDetails := distFlag.BoolP("details", "d", false, "show more details of the upgrade process")

	distFlag.Parse(args)

	zypperArgs := []string{"dist-upgrade"}

	if *useDetails {
		zypperArgs = append(zypperArgs, "--details")
	}

	if *verbose {
		fmt.Println("Running: zypper dist-upgrade")
	}

	if *help {
		distFlag.Usage()
		return
	}

	cmd := exec.Command("zypper", zypperArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("zypper dist-upgrade failed:", err)
		os.Exit(1)
	}
}
