# Godi
![Go](https://github.com/riyadhalnur/godi/workflows/Go/badge.svg?branch=master)  

> 'Godi' - pillow in Bengali, is a simple but highly opinionated Go server template that you can
> use to bootstrap your Go web applications and/or use as use the modules in your project as you see fit.

### Requirements  
1. [Go](https://golang.org) >= 1.13  
2. Direnv (optional)  

### Structure
```
.
|-- cmd
|   `-- api
|       `-- main.go
|-- go.mod
|-- go.sum
|-- pkg
|   |-- middleware
|   |   |-- request.go
|   |   `-- request_test.go
|   `-- server
|       |-- config.go
|       |-- server.go
|       |-- server_test.go
|       `-- util
|           |-- event.go
|           |-- response.go
|           |-- response_test.go
|           |-- route.go
|           |-- type.go
|           `-- type_test.go
|-- README.md
`-- static
    |-- css
    |   `-- main.css
    `-- index.html
```  

### Developing  
Run tests using
```
go test -v ./...
```  

To run the server, in `cmd/api/`,
```
go run main.go
```

Build a binary in `cmd/api` using  
```
go build -o <binary-name>
```  

### Healthcheck
The server package exposes a health endpoint by default at `/healthz`. This is compatible with Kubernetes integration.  

### Adding new services and middlewares
**Middlewares**  
By default, the default server will mount a request ID middleware that adds an `X-Request-ID` header to all requests. To define new middlewares, define it inside `pkg/middleware` and then mount/register it with the server instance,  
```go
srv := &Server{}
srv.AddMiddlewares(middleware.MyMiddlewareFunc)
...
``` 
*P.S.* Middleware order matters. You can also use any middleware that implements the basic middleware   

**Services**  
To add and register a service, create a new folder under `pkg/service`. Export the list of routes it will use and then register them with the server instance.    
```go
// in pkg/service/user/routes.go
Routes := []util.Route{
  &Route{
    "createUser",
    http.MethodPost,
    "/user",
    CreateUser,
  },
}

// then in your main.go
srv := &Server{}
srv.AddRoutes(user.Routes...)
...
```  

### Environment variables  
Environment variables are never read by the `pkg/server` package; it uses the `Configuration` struct passed in. Use environment variables when implementing it. Refer to `cmd/api/main.go` for usage.  
```
PORT=<port-server-listens-on> // defaults to 3001
STATIC_DIR=<static-file-directory>
TIMEOUT=<server-timeout-in-seconds> // write/read/idle timeouts
DEBUG=<true-or-false>
```  

### Contributing  
Read the [CONTRIBUTING](CONTRIBUTING.md) guide for information.  

### License  
Licensed under MIT. See [LICENSE](LICENSE) for more information.  

### Issues  
Report a bug in [issues](https://github.com/riyadhalnur/godi/issues).   

Made with love in Kuala Lumpur, Malaysia by [Riyadh Al Nur](https://verticalaxisbd.com)
