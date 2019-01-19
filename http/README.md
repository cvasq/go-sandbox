## HTTP Servers and Clients in Go

Examples:

**http_client.go** - A basic HTTP client executing a GET request to google.com

```
go run http_client.go
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en"><head><meta content="Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for." name="description"><meta content="noodp" name="robots"><meta content="text/html; 
...
```

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

**gorilla_mux_url_router.go** - Using the gorilla/mux http url router

Running the server
```
 go run gorilla_mux_routing.go
2019/01/18 21:01:48 Listening for HTTP connections on port: 8080
2019/01/18 21:01:48 /
2019/01/18 21:02:05 /list-users
...
```

Testing the endpoints
```
$ curl localhost:8080/
Endpoints:

// List all users
curl localhost:8080/list-users

// Get user name by ID
curl localhost:8080/get-user/id/bh0kkf47uh3blslo4t80

// Add new user
curl -X POST "http://localhost:8080/add-user/?first=Bill&last=Nye&age=35"

```

