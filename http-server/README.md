## Simple HTTP server in go

#### Run server

`go run main.go`

#### Connect to server using telnet

```console
   $ telnet localhost 7292
   GET /foo HTTP/1.1
   Host: localhost

   HTTP/1.1 200 OK
   Date: Sat, 29 Jun 2024 14:38:54 GMT
   Content-Length: 0

```
