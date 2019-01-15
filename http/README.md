## HTTP Servers and Clients in Go

**simple_http_server.go** - A Simple Web Server with basic logging


```
$  go run simple_http_server.go
2019/01/14 18:33:16 Listening on port: 8080
2019/01/14 18:33:17 /?message=Hello
```
```
$ curl http://localhost:8080/?message=Hello
{"Message":"Hello"}
```


