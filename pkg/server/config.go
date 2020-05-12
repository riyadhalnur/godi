package server

// Config specifies the parameters
// that can be passed in to a Server instance
//
// Port - tcp port the server will listen on. Defaults to 3000
// Timeout (required) - the write/read/idle timeout in seconds
// StaticDir - the server from static files will be served
type Config struct {
	Port      string
	Timeout   int
	StaticDir string
}
