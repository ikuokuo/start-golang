# Start [Gin][]

[Gin]: https://github.com/gin-gonic/gin

## Install

```zsh
# clone this module and init
git clone --depth 1 https://github.com/ikuokuo/start-golang.git
cd start-golang/_/start-gin
go mod tidy

# create this module from scratch
mkdir -p start-golang/_/start-gin
cd start-golang/_/start-gin
go mod init github.com/ikuokuo/start-golang/_/start-gin
go get -u github.com/gin-gonic/gin
go get -u github.com/stretchr/testify/assert
go get -u github.com/google/uuid
```

## Basic Example

`examples/quickstart.go`:

```go
package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

Run:

```zsh
❯ go run examples/quickstart.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.setupRouter.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Test:

```zsh
❯ curl http://localhost:8080/ping
{"message":"pong"}

# examples/quickstart_test.go
❯ go clean -testcache && go test ./examples
ok      github.com/ikuokuo/start-golang/_/start-gin/examples    0.019s
```

## [API Examples](https://github.com/gin-gonic/gin#api-examples)

Go through Gin examples to learn api usage.

- [API Examples](https://github.com/gin-gonic/gin#api-examples)
- [Doc Examples](https://gin-gonic.com/docs/examples/)

## Getting Started

### Design RESTful API

- /albums
  - GET - Get all albums
  - POST - Create a new album
- /albums/:id
  - GET - Get an album by its id
  - PUT - Update an album by its id
  - DELETE - Delete an album by its id
- /files/upload
  - POST - Upload a file

How to design API, please see section [References](#references).

### Project Structure

```zsh
❯ tree start-gin -aF --dirsfirst
start-gin
├── app/
│   ├── api/
│   │   ├── albums.go
│   │   └── files.go
│   ├── config/
│   │   └── version.go
│   ├── entity/
│   │   └── album.go
│   ├── router/
│   │   ├── middleware/
│   │   │   ├── cors.go
│   │   │   └── version.go
│   │   └── router.go
│   ├── service/
│   │   └── db.go
│   └── main.go
├── examples/
│   ├── quickstart.go
│   └── quickstart_test.go
├── .gitignore
├── README.md
├── go.mod
└── go.sum
```

How to organize the project, we could learn from [projects using Gin](https://gin-gonic.com/docs/users/), or search it in GitHub, for example, [Mindinventory/Golang-Project-Structure](https://github.com/Mindinventory/Golang-Project-Structure).

### Write the code

#### Create data

Define `Album` struct in `entity/album.go`:

```go
type Album struct {
	ID        string    `form:"id"         json:"id"         binding:"-"`
	Title     string    `form:"title"      json:"title"      binding:"required"`
	Artist    string    `form:"artist"     json:"artist"     binding:"required"`
	Price     float64   `form:"price"      json:"price"      binding:"required"`
	CreatedAt time.Time `form:"created_at" json:"created_at" binding:"-"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" binding:"-"`
}
```

About struct field tags, please see [Model binding and validation](https://github.com/gin-gonic/gin#model-binding-and-validation):

- `form:"<key>"`, `json:"<key>"`: set its key name
- `binding:"required"`: required fields, will error if not given
- `binding:"-"`: ignored fields, won't error if not given

Define `map` as database in `service/db.go`:

```go
var DB = make(map[string]interface{})
```

#### Write handler

Write our handlers in `api/` directory. Such as `api/albums.go`:

```go
// Get all albums
func GetAlbums(c *gin.Context) {
	albums := make([]entity.Album, 0, len(service.DB))
	for _, a := range service.DB {
		albums = append(albums, a.(entity.Album))
	}
	if len(albums) > 0 {
		c.JSON(http.StatusOK, albums)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "albums are empty"})
	}
}

// ...
```

#### Write router

Write the router in `router/router.go`:

```go
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/albums", api.GetAlbums)
	r.POST("/albums", api.PostAlbums)

	r.GET("/albums/:id", api.GetAlbumByID)
	r.PUT("/albums/:id", api.UpdateAlbumByID)
	r.DELETE("/albums/:id", api.DeleteAlbumByID)

	r.POST("/files/upload", api.PostFiles)

	return r
}
```

We associate HTTP method and path with the handlers here.

Finnally, attach the router to the server in `main.go`:

```go
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
```

### Run the code

```zsh
❯ go run app/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /albums                   --> github.com/ikuokuo/start-golang/_/start-gin/app/api.GetAlbums (5 handlers)
[GIN-debug] POST   /albums                   --> github.com/ikuokuo/start-golang/_/start-gin/app/api.PostAlbums (5 handlers)
[GIN-debug] GET    /albums/:id               --> github.com/ikuokuo/start-golang/_/start-gin/app/api.GetAlbumByID (5 handlers)
[GIN-debug] PUT    /albums/:id               --> github.com/ikuokuo/start-golang/_/start-gin/app/api.UpdateAlbumByID (5 handlers)
[GIN-debug] DELETE /albums/:id               --> github.com/ikuokuo/start-golang/_/start-gin/app/api.DeleteAlbumByID (5 handlers)
[GIN-debug] POST   /files/upload             --> github.com/ikuokuo/start-golang/_/start-gin/app/api.PostFiles (5 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Use `curl` to make some requests (`jq` to pretty-print json):

```zsh
# create an album
❯ curl -s -X POST http://localhost:8080/albums \
-H 'Content-Type: application/json' \
-d '{
  "title": "Start Gin",
  "artist": "GoCoding",
  "price": 0.99
}' | jq
{
  "id": "73e3a9d5-c689-4002-aa42-c96df82b178c",
  "title": "Start Gin",
  "artist": "GoCoding",
  "price": 0.99,
  "created_at": "2021-11-04T20:30:00.057259+08:00",
  "updated_at": "2021-11-04T20:30:00.057259+08:00"
}

# get all albums
❯ curl -s http://localhost:8080/albums | jq
[
  {
    "id": "73e3a9d5-c689-4002-aa42-c96df82b178c",
    "title": "Start Gin",
    "artist": "GoCoding",
    "price": 0.99,
    "created_at": "2021-11-04T20:30:00.057259+08:00",
    "updated_at": "2021-11-04T20:30:00.057259+08:00"
  }
]
```

```zsh
# get an album
curl -s http://localhost:8080/albums/73e3a9d5-c689-4002-aa42-c96df82b178c | jq

# update an album, using form
curl -s -X PUT http://localhost:8080/albums/73e3a9d5-c689-4002-aa42-c96df82b178c \
-H 'Content-Type: multipart/form-data' \
-F 'title=Start Gin 2' \
-F 'artist=GoCoding' \
-F 'price=9.99' | jq

# delete an album
curl -s -X DELETE http://localhost:8080/albums/73e3a9d5-c689-4002-aa42-c96df82b178c | jq

# upload a file
curl -X POST http://localhost:8080/files/upload \
-H "Content-Type: multipart/form-data" \
-F "file=@$HOME/Downloads/service.png"
```

If wanna see response headers attached by middlewares:

```zsh
❯ curl -v -s http://localhost:8080/albums | jq
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /albums HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Access-Control-Allow-Credentials: true
< Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max
< Access-Control-Allow-Methods: POST, GET, PUT, PATCH, DELETE
< Access-Control-Allow-Origin: *
< Access-Control-Max-Age: 86400
< Content-Type: application/json
< X-Start-Gin-Version: 0.1.0
< Date: Thu, 04 Nov 2021 20:30:00 GMT
< Content-Length: 196
<
{ [196 bytes data]
* Connection #0 to host localhost left intact
* Closing connection 0
[
  {
    "id": "73e3a9d5-c689-4002-aa42-c96df82b178c",
    "title": "Start Gin",
    "artist": "GoCoding",
    "price": 0.99,
    "created_at": "2021-11-04T20:30:00.057259+08:00",
    "updated_at": "2021-11-04T20:30:00.057259+08:00"
  }
]
```

## References

- [Gin Doc](https://gin-gonic.com/docs/)
- [Tutorial: Developing a RESTful API with Go and Gin](https://golang.org/doc/tutorial/web-service-gin)

API Design:

- [HTTP API Design Guide](https://github.com/interagent/http-api-design)
- [RESTful API design](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)
- [RESTful API 最佳实践](http://www.ruanyifeng.com/blog/2018/10/restful-api-best-practices.html)
- [GitHub REST API](https://docs.github.com/en/rest)

API Idempotent:

- [Idempotent REST APIs](https://restfulapi.net/idempotent-rest-apis/)
- 中文，请搜：幂等性
