package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/canerakdas/potato/color"
	"github.com/canerakdas/potato/new"
)

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(color.Bright, "USAGE  ", color.Reset)
		fmt.Println("    potato new [appname]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])
		new.Create(newCmd.Args()[0])
	case "generate":
		generateCmd.Parse(os.Args[2:])
	case "version":
		fmt.Println("0.0.1")
	default:
		fmt.Println(color.Bright, "USAGE  ", color.Reset)
		fmt.Println("    potato new [appname]")
		os.Exit(1)
	}
}
