# Start Go

Getting started:

- [A Tour of Go](https://tour.golang.org/), [_tour](_tour)

[Standard library](https://golang.org/pkg/):

- [crypto](https://pkg.go.dev/crypto)
  - [aes](https://pkg.go.dev/crypto/aes)
    - [aes_encrypt.go](_crypto/aes_encrypt.go)
    - [aes_decrypt.go](_crypto/aes_decrypt.go)
- [encoding](https://pkg.go.dev/encoding)
  - [json](https://pkg.go.dev/encoding/json)
    - [json_marshal.go](_encoding/json_marshal.go)
    - [json_unmarshal.go](_encoding/json_unmarshal.go)
  - [xml](https://pkg.go.dev/encoding/xml)
    - [xml_marshal.go](_encoding/xml_marshal.go)
    - [xml_unmarshal.go](_encoding/xml_unmarshal.go)
- [io](https://pkg.go.dev/io)
  - [ioutil](https://pkg.go.dev/io/ioutil)
    - [ioutil_readfile.go](_io/ioutil_readfile.go)
    - [ioutil_writefile.go](_io/ioutil_writefile.go)
- [net](https://pkg.go.dev/net)
  - [http](https://pkg.go.dev/net/http)
    - [http_server.go](_net/http_server.go)
  - [interface.go](_net/interface.go)

Popular library:

- Command Line
  - [cobra](https://github.com/spf13/cobra), [_/start-cobra](_/start-cobra)
  - [urfave/cli](https://github.com/urfave/cli), [_/start-cli](_/start-cli)
- ORM
  - [ent](https://github.com/ent/ent), [_/start-ent](_/start-ent)
- Web Frameworks
  - [Gin](https://github.com/gin-gonic/gin), [_/start-gin](_/start-gin)
- YAML
  - [go-yaml](https://github.com/go-yaml/yaml), [_/start-yaml](_/start-yaml)

## References

- [Go Doc](https://golang.org/doc/)
- [Awesome Go](https://github.com/avelino/awesome-go)
