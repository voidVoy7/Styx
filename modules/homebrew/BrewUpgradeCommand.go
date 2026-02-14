package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewUpgradeCommmand(args []string) {
	upgradeflags := pflag.NewFlagSet("upgrade", pflag.ExitOnError)

	verbose := upgradeflags.BoolP("verbose", "v", false, "show extra output")

	upgradeflags.Parse(args)

	if *verbose {
		fmt.Println("Running: brew upgrade")
	}

	cmd := exec.Command("brew", "upgrade")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("brew upgrade failed!", err)
		os.Exit(1)
	}
}
