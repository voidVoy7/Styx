package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewUpdateCommand(args []string) {
	updateFlags := pflag.NewFlagSet("update", pflag.ExitOnError)

	verbose := updateFlags.BoolP("verbose", "v", false, "show extra output")

	updateFlags.Parse(args)

	if *verbose {
		fmt.Println("Running: brew update")
	}

	cmd := exec.Command("brew", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("brew update failed:", err)
		os.Exit(1)
	}
}
