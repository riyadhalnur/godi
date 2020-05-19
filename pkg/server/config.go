package server

// Config specifies the parameters
// that can be passed in to a Server instance
//
// Port (required) - tcp port the server will listen on
// Timeout (required) - the write/read/idle timeout in seconds
// StaticDir - the server from static files will be served
type Config struct {
	Port      string
	Timeout   int
	StaticDir string
}
