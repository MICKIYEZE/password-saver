package main

import (
	"flag"
	"fmt"
	"os"
	"password-saver/pkg/db"
	"password-saver/pkg/ui"
)

func main() {
	add := flag.Bool("add", false, "Add a new password entry")
	file := flag.String("file", "passwords.json", "Path to the JSON database file")

	flag.Parse()

	d := db.New(*file)
	if err := d.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load database:", err)
		os.Exit(1)
	}

	// Show existing entries each launch
	ui.Display(d)

	if *add {
		if err := ui.PromptAdd(d); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to add entry:", err)
			os.Exit(1)
		}
	}
}
