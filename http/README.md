## HTTP Servers and Clients in Go

Examples:

**simple_http_server.go** - A simple HTTP server with basic logging

Run the HTTP server
```
$  go run simple_http_server.go
2019/01/14 18:33:16 Listening for HTTP connections on port: 8080
2019/01/14 18:33:17 /?message=Hello
```
```
$ curl "http://localhost:8080/?message=Hello"
{"Message":"Hello"}
```

**simple_https_server.go** - A simple HTTPS server with basic logging

First generate a self-signed TLS cert with the openssl tool
```
$ openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout selfsigned.key.pem -out selfsigned-x509.crt -subj "/C=US/ST=WA/L=Seattle/CN=example.com/emailAddress=someEmail@gmail.com"

Output files: selfsigned.key.pem  selfsigned-x509.crt
```

Run the HTTPS server with generated SSL certs
```
$ go run simple_https_server.go
2019/01/15 00:08:47 Listening for HTTPS connections on port: 8080
2019/01/15 00:09:34 /?message=Sup
```
```
$ curl -k "https://localhost:8080/?message=Sup"
{"Message":"Sup"}
```

