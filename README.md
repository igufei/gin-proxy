# Proxy Middleware for Gin Framework

This is a middleware for [Gin](https://github.com/gin-gonic/gin) framework.

## Usage

Download and install it:

```sh
$ go get github.com/igufei/gin-proxy
```

Import it in your code:

```go
import "github.com/igufei/gin-proxy"
```

## Example

```go
engine := gin.New()
engine.Use(ginproxy.HandleProxy("/api",ginproxy.ProxyOptions{
    Target:"http://www.google.com",
    PathRewrite:"/api",
}))