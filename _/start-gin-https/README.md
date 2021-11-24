# Start [Gin][] with HTTPS

[Gin]: https://github.com/gin-gonic/gin

## Install

```zsh
# clone this module and init
git clone --depth 1 https://github.com/ikuokuo/start-golang.git
cd start-golang/_/start-gin-https
go mod tidy

# create this module from scratch
mkdir -p start-golang/_/start-gin-https
cd start-golang/_/start-gin-https
go mod init github.com/ikuokuo/start-golang/_/start-gin-https
go get -u github.com/gin-gonic/gin
# gin-gonic/autotls
go get -u github.com/gin-gonic/autotls
# unrolled/secure
go get -u github.com/unrolled/secure
```

## Use [gin-gonic/autotls](https://github.com/gin-gonic/autotls)

### 1-line LetsEncrypt HTTPS server

- [Let's Encrypt](https://letsencrypt.org/)

[example1.go](example1.go):

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
```

### Custom autocert manager

- [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)

[example2.go](example2.go):

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("./www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
```

## Use [unrolled/secure](https://github.com/unrolled/secure)

### [Integration: Gin](https://github.com/unrolled/secure#gin)

[secure_gin.go](secure_gin.go):

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure" // or "gopkg.in/unrolled/secure.v1"
)

func main() {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	router := gin.Default()
	router.Use(secureFunc)

	router.GET("/", func(c *gin.Context) {
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
		c.String(200, "X-Frame-Options header is now `DENY`.")
	})

	router.RunTLS(":3000", "cert.pem", "key.pem")
}
```

Run:

```zsh
go run secure_gin.go
```

Test:

```zsh
❯ curl http://localhost:3000
Client sent an HTTP request to an HTTPS server.

❯ curl -k --cert cert.pem --key key.pem https://localhost:3000
X-Frame-Options header is now `DENY`.
```

### [Redirecting HTTP to HTTPS](https://github.com/unrolled/secure#redirecting-http-to-https)

[secure_redirect.go](secure_redirect.go):

```go
package main

import (
	"log"
	"net/http"

	"github.com/unrolled/secure" // or "gopkg.in/unrolled/secure.v1"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
})

func main() {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     "localhost:8443", // This is optional in production. The default behavior is to just redirect the request to the HTTPS protocol. Example: http://github.com/some_page would be redirected to https://github.com/some_page.
	})

	app := secureMiddleware.Handler(myHandler)

	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(":8080", app))
	}()

	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", app))
}
```

Run:

```zsh
go run secure_redirect.go
```

Test:

```zsh
❯ curl http://localhost:8080
<a href="https://localhost:8443/">Moved Permanently</a>.

❯ curl -k --cert cert.pem --key key.pem https://localhost:8443
hello world
```
