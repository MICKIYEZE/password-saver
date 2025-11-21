package entry

// Entry represents a saved credential.
type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Notes    string `json:"notes"`
}
