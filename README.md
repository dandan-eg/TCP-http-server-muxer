# TCP Server HTTP muxer

TCP Server HTTP is a simple HTTP server implemented in Go. It listens for incoming TCP connections on port 8080 and
handles HTTP requests by routing them to the appropriate handlers.

## Purpose

The purpose of this project is to serve as a learning exercise focused on understanding TCP and HTTP protocols. It was
developed with the intention of gaining a deeper understanding of networking concepts and exploring the implementation
of an HTTP server using the TCP protocol.

## Usage

To use this server, follow the steps below:

1. Clone the repository:

```shell
git clone https://github.com/dandan-eg/TCP-http-server-muxer.git
```

2. Navigate to the project directory:

```shell
cd TCP-server-http
```

3. Build the project:

```shell
go build
```

4. Run the executable:

```shell
./TCP-http-server-muxer
```

The server will start listening on port 8080. You can access it by opening a web browser and
entering http://localhost:8080 in the address bar. You should see a basic HTML page "Home"

## Dependencies

This server use the following Go packages:

- net: Provides basic networking operations, including TCP/IP socket communication.

The package is part of the Go standard library and is available by default.

## Project Structure

The project is structured as follows:

- `main.go:` The entry point of the application. It sets up the TCP server. Incoming connections are accepted, and each
  connection is handled concurrently by launching separate goroutines. The server reads the HTTP request from the
  client, sends it to the muxer for routing, and receives the corresponding response template. Finally, the server
  writes the response back to the client, ensuring the inclusion of appropriate headers and content.
- `request/request.go:` Provides functions for parsing HTTP requests and extracting request parameters.
- `muxer/muxer.go:` Implements a request multiplexer (mux) that matches incoming requests to registered routes and
  invokes
  the corresponding template.
- `muxer/route.go:` Defines the Route struct, which represents a registered route in the muxer.
- `template/template.go:` Contains template functions that generate HTML content for different pages.

## Customizing Routes

To customize the routes handled by the server, modify the routes function in the routes.go file. The routes function
uses
the muxer package to define the routes and associate them with the appropriate template functions.

```go
package main

import (
	"TCP-server-http/muxer"
	"TCP-server-http/template"
	"fmt"
)

func routes() muxer.Muxer {
	m := muxer.New()

	//routes goes here
	m.Add("GET", "/", template.Home)
	m.Add("GET", "/account", template.Account)
	m.Add("POST", "/account", template.Delete)
	m.Add("GET", "/register", template.Register)
	

	fmt.Println(m)
	return m
}
```

## Conclusion

In conclusion, this project serves as a personal learning endeavor aimed at gaining a better understanding of TCP and
HTTP protocols. It provides a TCP server that handles HTTP requests and demonstrates how routing and template rendering
can be implemented.

It is important to note that this project is not intended for actual usage in production
environments.