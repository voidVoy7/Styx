package main

import (
	"fmt"
	"log"
	"os"

	homebrew "github.com/Omniwrap/modules/homebrew"
	pacman "github.com/Omniwrap/modules/pacman"
	xbps "github.com/Omniwrap/modules/xbps"
	zypper "github.com/Omniwrap/modules/zypper"
	"github.com/hairyhenderson/go-which"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ow <command>")
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
		fmt.Println("---pacman---")
		fmt.Println("  install")
		fmt.Println("  remove")
		fmt.Println("  search")
		fmt.Println("  update")
		fmt.Println("  query")
		fmt.Println("---xbps---")
		fmt.Println("  install")
		fmt.Println("  remove")
		fmt.Println("  search")
		fmt.Println("  update")
		os.Exit(1)
	}

	// Super sneaky easter egg handling :P
	cmd := os.Args[1]
	switch cmd {
	case "sao":
		fmt.Println("I'd rather trust and regret than doubt and regret. - Kirigaya Kazuto 'Kirito' Sword Art Online")
		return
	case "konami":
		fmt.Println("↑ ↑ ↓ ↓ ← → ← → B A")
		return
	}

	pm, ok := PackageManagerCheck()
	if !ok {
		log.Fatal("no supported package manager found")
	}

	switch pm {

	case "pacman":
		switch os.Args[1] {
		case "install":
			pacman.PacmanInstallCommand(os.Args[2:])
		case "search":
			pacman.PacmanSearchCommand(os.Args[2:])
		case "remove":
			pacman.PacmanRemoveCommand(os.Args[2:])
		case "update":
			pacman.PacmanUpdateCommand(os.Args[2:])
		case "query":
			pacman.PacmanQueryCommand(os.Args[2:])
		default:
			fmt.Println("Unknown command:", os.Args[1])
			fmt.Println("Commands:")
			fmt.Println("  install")
			fmt.Println("  remove")
			fmt.Println("  search")
			fmt.Println("  update")
			fmt.Println("  query")
		}

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

	case "xbps-install":

		switch os.Args[1] {
		case "update":
			xbps.XbpsUpdateCommand(os.Args[2:])
		case "search":
			xbps.XbpsSearchCommand(os.Args[2:])
		case "install":
			xbps.XbpsInstallCommand(os.Args[2:])
		case "remove":
			xbps.XbpsRemoveCommand(os.Args[2:])

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
	supported := []string{"brew", "zypper", "pacman", "xbps-install"}

	for _, candidate := range supported {
		if which.Found(candidate) {
			return candidate, true
		}
	}

	return "", false
}
