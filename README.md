# simple-go-server
> Simple but highly opinionated Go server template

### Requirements  
1. [Go](https://golang.org) >= 1.13  
2. Direnv (optional)  

### Structure
```
|-- go.mod
|-- go.sum
|-- middleware
|   |-- auth.go
|   `-- logging.go
|-- README.md
|-- router.go
|-- routes.go
|-- server.go
|-- static
|   |-- css
|   |   `-- main.css
|   `-- index.html
|-- user.go
|-- user_test.go
`-- utils
    `-- response.go
```  

### Developing  
Run tests using
```
go test -v
```  

This package uses [Gin](https://github.com/codegangsta/gin) for automatic server reload for any file changes. To run the server,
```
gin run server.go
```

Build a binary using  
```
go build -o <binary-name>
```  

### Adding new routes and controllers  
To add and register a new route, update the `routes.go` file to define the methods, name and handler for your controllers. The new routes will be automatically mounted under `/api` when the server is initialised.  

For example, in `routes.go`  
```go
&Route{
  "createUser",
  http.MethodPost,
  "/user",
  CreateUser,
},
```  

### Environment variables  
There are only 2 environment variables used by default,
```
PORT=<port-server-listens-on> // defaults to 3001
STATIC=<static-file-directory> // defaults to /static
```  

### Contributing  
Read the [CONTRIBUTING](CONTRIBUTING.md) guide for information.  

### License  
Licensed under MIT. See [LICENSE](LICENSE) for more information.  

### Issues  
Report a bug in [issues](https://github.com/riyadhalnur/simple-go-server/issues).   

Made with love in Kuala Lumpur, Malaysia by [Riyadh Al Nur](https://verticalaxisbd.com)
