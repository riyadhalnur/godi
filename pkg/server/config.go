package server

// Config holds the required configurations
// to pass into a server instance
type Config struct {
	// Port to listen on. Defaults to 3000
	Port string

	// Timeout (required) for write/read/idle in seconds
	Timeout int

	// Static path to the directory
	// from which to server static files
	// Leave blank to disable
	StaticDir string

	// Debug turn on debug mode
	Debug bool
}
