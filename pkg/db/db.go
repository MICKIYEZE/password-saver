package db

import (
	"encoding/json"
	"errors"
	"os"
	"password-saver/pkg/entry"
	"sync"
)

// Database handles JSON persistence of entries.
type Database struct {
	FilePath string
	mu       sync.RWMutex
	Entries  []entry.Entry
}

// New creates a new Database instance pointing to the provided file path.
func New(filePath string) *Database {
	return &Database{FilePath: filePath, Entries: []entry.Entry{}}
}

// Load reads entries from the JSON file if it exists.
func (d *Database) Load() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	f, err := os.Open(d.FilePath)
	if errors.Is(err, os.ErrNotExist) {
		// nothing to load
		d.Entries = []entry.Entry{}
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	return dec.Decode(&d.Entries)
}

// Save writes the current entries to the JSON file (atomic-ish by truncating).
func (d *Database) Save() error {
	d.mu.RLock()
	entries := d.Entries
	d.mu.RUnlock()

	f, err := os.Create(d.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(entries)
}

// Add appends an entry to the database.
func (d *Database) Add(e entry.Entry) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, e)
}

// List returns a copy of the entries for safe concurrent usage.
func (d *Database) List() []entry.Entry {
	d.mu.RLock()
	defer d.mu.RUnlock()
	// return a shallow copy
	copyEntries := make([]entry.Entry, len(d.Entries))
	copy(copyEntries, d.Entries)
	return copyEntries
}
