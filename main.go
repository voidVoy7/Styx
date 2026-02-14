package main

import (
	"fmt"
	"os"

	"github.com/Bowser/modules"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bowser <command>")
		fmt.Println("Commands: update")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "update":
		modules.UpdateCommand(os.Args[2:])
	case "upgrade":
		modules.UpgradeCommmand(os.Args[2:])
	case "install":
		modules.InstallCommand(os.Args[2:])
	case "search":
		modules.SearchCommand(os.Args[2:])
	case "remove":
		modules.RemoveCommand(os.Args[2:])
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
}
