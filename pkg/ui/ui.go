package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"password-saver/pkg/db"
	"password-saver/pkg/entry"
)

// Display prints all saved entries
func Display(d *db.Database) {
	entries := d.List()
	if len(entries) == 0 {
		fmt.Println("No saved passwords yet.")
		return
	}
	fmt.Println("Saved passwords:")
	fmt.Println("----------------")
	for i, e := range entries {
		fmt.Printf("%d) Username: %s\n Password: %s\n Notes: %s\n\n", i+1, e.Username, e.Password, e.Notes)
	}
}

// PromptAdd interacts with the user to create a new entry and saves it.
func PromptAdd(d *db.Database) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Enter notes: ")
	notes, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	e := entry.Entry{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
		Notes:    strings.TrimSpace(notes),
	}

	d.Add(e)
	if err := d.Save(); err != nil {
		return err
	}
	fmt.Println("Entry saved!")
	return nil
}
