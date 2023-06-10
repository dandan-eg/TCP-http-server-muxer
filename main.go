package main

import (
	"TCP-server-http/muxer"
	"TCP-server-http/request"
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer func() { log.Println(li.Close()) }()

	mux := routes()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn, mux)
	}
}

func handle(conn net.Conn, mux muxer.Muxer) {
	defer conn.Close()

	req := read(conn)
	write(conn, mux.Serve(req))
}

const requestIdx = 0

func read(conn net.Conn) request.Request {
	var (
		i = 0
		p request.Request
	)
	sc := bufio.NewScanner(conn)

	for sc.Scan() {
		txt := sc.Text()

		if i == requestIdx {
			p = request.New(txt)
			log.Printf("[%s] %s", p.URI, p.Method)
		}

		if txt == "" {
			break
		}

		i++
	}

	return p
}

func write(conn net.Conn, body string) {

	html := fmt.Sprintf(
		`
<!DOCTYPE html>
<html lang="en">
	<body>
		%s
	</body>
</html>`, body)

	//headers
	if body == "404" {
		fmt.Fprintf(conn, "HTTP/1.1 404 NOTFOUND\r\n")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")

	}
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(html))
	fmt.Fprintf(conn, "Content-Type : text/html\r\n")

	fmt.Fprintf(conn, "\r\n")

	//body
	fmt.Fprintf(conn, html)
}
