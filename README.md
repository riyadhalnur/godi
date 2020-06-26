# Godi
![Go](https://github.com/riyadhalnur/godi/workflows/Go/badge.svg?branch=master)  

> 'Godi' - pillow in Bengali, is a simple but highly opinionated Go server template that you can
> use to bootstrap your Go web applications and/or use as use the modules in your project as you see fit.

### Requirements  
1. [Go](https://golang.org) >= 1.13  
2. Docker (optional)  
3. Kubernetes (optional)  
4. Kustomize (optional)  
5. Direnv (optional)  

### Structure
```
.
|-- cmd
|   `-- api
|       `-- main.go
|-- deploy
|   |-- base
|   |   |-- deployment.yml
|   |   |-- kustomization.yml
|   |   `-- service.yml
|   `-- overlays
|       `-- dev
|           |-- config-map.yml
|           `-- kustomization.yml
|-- Dockerfile
|-- go.mod
|-- go.sum
|-- Makefile
|-- pkg
|   |-- godierr
|   |   |-- error.go
|   |   |-- error_test.go
|   |   |-- sentinel.go
|   |   `-- sentinel_test.go
|   |-- logger
|   |   |-- logger.go
|   |   `-- logger_test.go
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
```shell  
make test
```  

To run the server, in `cmd/api/`,
```shell  
make run
```

Build a binary in `cmd/api` using  
```shell  
make build
```  

Build a Docker image  
```shell  
docker build -t godi .  
```  

Deploy to Kubernetes  
```shell  
kubectl apply -k deploy/overlay/dev  
```  

### Healthcheck
The server package exposes a health endpoint by default at `/health`. This is compatible with Kubernetes integration.  

### Logging
The logger package is modeled after the standard `log` package in Go to expose a global logger that is configured to use for a uniform logging experience across the application. It wraps around `zap` with custom configuration that plays nice with Docker, Kubernetes and Stackdriver.  

Methods not ending with `f` are aliases for `zap.<level>w` methods that accept loosely typed key-value pairs, e.g.  
```go
logger.Info("Failed to fetch URL.",
    "url", url,
)
```  

If `DEBUG` mode is `true` in the environment, debug level messages are logged to `stdout`. Otherwise, everything less than error level but not debug level is routed to `stdout`. The benefits of this are being able to turn on debug mode without having to redeploy your application. Just a simple restart will do.    

Errors and anything above are routed to `stderr` by default.  

### Adding new services and middlewares
**Middlewares**  
By default, the default server will mount a request ID middleware that adds an `X-Request-ID` header to all requests. To define new middlewares, define it inside `pkg/middleware` and then mount/register it with the server instance,  
```go
srv := &Server{}
srv.AddMiddlewares(middleware.MyMiddlewareFunc)
...
``` 
*P.S.* Middleware order matters. You can also use any middleware that matches the `http.HandlerFunc` signature     

**Services**  
To add and register a service, create a new folder under `pkg/service/<service-name>` or `pkg/api/<api-name>` or `api/<api-name>`. How you wish to structure your routes and their respective controllers is upto you. Make sure to export the list of routes it will use and then register them with the server instance to mount them when it starts up.    
```go
// in ../user/routes.go
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

### Static files  
The boilerplate comes with a basic HTML page and a rudimentary stylesheet inside the `/static` folder. By default, the server will not serve any static files. You have to explicitly pass in the path to the folder when configuring the server instance. The static files though are always served at the `/static` path of the listening server.  

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
