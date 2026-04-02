package pacman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func PacmanQueryCommand(args []string) {
	queryFlag := pflag.NewFlagSet("query", pflag.ExitOnError)
	queryFlag.Usage = func() {
		fmt.Println("Usage: ow query [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		queryFlag.PrintDefaults()
	}

	help := queryFlag.BoolP("help", "h", false, "show helpful information")
	verbose := queryFlag.BoolP("verbose", "v", false, "show extra output")
	useChangelog := queryFlag.BoolP("changelog", "c", false, "view the changelog of a package")
	useDeps := queryFlag.BoolP("deps", "d", false, "list packages installed as dependencies")
	useExplicit := queryFlag.BoolP("explicit", "e", false, "list packages explicitly installed")
	useGroups := queryFlag.BoolP("groups", "g", false, "view all members of a package group")
	useInfo := queryFlag.BoolP("info", "i", false, "view package information")
	useCheck := queryFlag.BoolP("check", "k", false, "check that package files exist")
	useList := queryFlag.BoolP("list", "l", false, "list the files owned by the queried package")

	queryFlag.Parse(args)

	if *help {
		queryFlag.Usage()
		return
	}

	pkgs := queryFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow query requires a package name")
		fmt.Println()
		queryFlag.Usage()
		return
	}

	pacmanArgs := []string{"-Q"}

	if *useChangelog {
		pacmanArgs = append(pacmanArgs, "--changelog")
	}

	if *useDeps {
		pacmanArgs = append(pacmanArgs, "--deps")
	}

	if *useExplicit {
		pacmanArgs = append(pacmanArgs, "--explicit")
	}

	if *useGroups {
		pacmanArgs = append(pacmanArgs, "--groups")
	}

	if *useInfo {
		pacmanArgs = append(pacmanArgs, "--info")
	}

	if *useCheck {
		pacmanArgs = append(pacmanArgs, "--check")
	}

	if *useList {
		pacmanArgs = append(pacmanArgs, "--list")
	}

	pacmanArgs = append(pacmanArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: pacman -Q")
	}

	cmd := exec.Command("pacman", pacmanArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("pacman -Q failed", err)
		os.Exit(1)
	}
}
