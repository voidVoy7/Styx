package main

import (
	"fmt"
	"log"
	"os"

	homebrew "github.com/Styx/modules/homebrew"
	zypper "github.com/Styx/modules/zypper"
	"github.com/hairyhenderson/go-which"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: styx <command>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("---homebrew---")
		fmt.Println("  update")
		fmt.Println("  upgrade")
		fmt.Println("  install")
		fmt.Println("  search")
		fmt.Println("  remove")
		fmt.Println("---zypper---")
		fmt.Println("  update")
		fmt.Println("  upgrade")
		fmt.Println("  patch")
		fmt.Println("  search")
		fmt.Println("  install")
		fmt.Println("  remove")
		os.Exit(1)
	}

	pm, ok := PackageManagerCheck()
	if !ok {
		log.Fatal("no supported package manager found")
	}

	switch pm {

	case "brew":
		switch os.Args[1] {
		case "update":
			homebrew.BrewUpdateCommand(os.Args[2:])
		case "upgrade":
			homebrew.BrewUpgradeCommmand(os.Args[2:])
		case "install":
			homebrew.BrewInstallCommand(os.Args[2:])
		case "search":
			homebrew.BrewSearchCommand(os.Args[2:])
		case "remove":
			homebrew.BrewRemoveCommand(os.Args[2:])
		default:
			fmt.Println("Unknown command:", os.Args[1])
			fmt.Println("Commands:")
			fmt.Println("  update")
			fmt.Println("  upgrade")
			fmt.Println("  install")
			fmt.Println("  search")
			fmt.Println("  remove")
			os.Exit(1)
		}

	case "zypper":

		switch os.Args[1] {
		case "update":
			zypper.ZypperUpdateCommand(os.Args[2:])
		case "patch":
			zypper.ZypperPatchCommand(os.Args[2:])
		case "search":
			zypper.ZypperSearchCommand(os.Args[2:])
		case "upgrade":
			zypper.ZypperDistUpgradeCommand(os.Args[2:])
		case "install":
			zypper.ZypperInstallCommand(os.Args[2:])
		case "remove":
			zypper.ZypperRemoveCommand(os.Args[2:])

		default:
			fmt.Println("Unknown command:", os.Args[1])
			fmt.Println("Commands:")
			fmt.Println("  update")
			fmt.Println("  upgrade")
			fmt.Println("  patch")
			fmt.Println("  search")
			fmt.Println("  install")
			fmt.Println("  remove")
		}
	}

}

func PackageManagerCheck() (string, bool) {
	supported := []string{"brew", "zypper"}

	for _, candidate := range supported {
		if which.Found(candidate) {
			return candidate, true
		}
	}

	return "", false
}
