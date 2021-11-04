# Start [YAML][]

[YAML]: https://github.com/go-yaml/yaml

- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3)

## Install

```bash
cd start-golang/_/start-yaml
go mod init github.com/ikuokuo/start-golang/_/start-yaml

go get -u gopkg.in/yaml.v3
```

Example: [example.go](example.go)

## Struct

[struct.go](common/struct.go):

```go
package common

import "time"

type Config struct {
	V         string `yaml:"version,omitempty"`
	CreatedAt time.Time
	Labels    []string `yaml:",flow"`
	Server    struct {
		Addr string
		Port int
	}
}

type ServerConfig struct {
	Addr string
	Port int
}
```

### Struct fields

The field tag format accepted is:

`(...) yaml:"[<key>][,<flag1>[,<flag2>]]" (...)`

The following flags are currently supported:

```txt
omitempty    Only include the field if it's not set to the zero
             value for the type or to empty slices or maps.
             Zero valued structs will be omitted if all their public
             fields are zero, unless they implement an IsZero
             method (see the IsZeroer interface type), in which
             case the field will be excluded if IsZero returns true.

flow         Marshal using a flow style (useful for structs,
             sequences and maps).

inline       Inline the field, which must be a struct or a map,
             causing all of its fields or keys to be processed as if
             they were part of the outer struct. For maps, keys must
             not conflict with the yaml keys of other struct fields.
```

In addition, if the key is "-", the field is ignored.

## [Marshal](https://pkg.go.dev/gopkg.in/yaml.v3#Marshal)

[marshal.go](marshal.go):

```go
cfg := common.Config{
	V:         "1.0.0",
	CreatedAt: time.Now(),
	Labels:    []string{"go", "coding"},
	Server: common.ServerConfig{
		Addr: "0.0.0.0",
		Port: 8000,
	},
}

d, err := yaml.Marshal(&cfg)
if err != nil {
	log.Fatalf("error: %v", err)
}
log.Print(string(d))
```

## [Unmarshal](https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal)

[unmarshal.go](unmarshal.go):

```go
var data = `
version: 0.0.0
createdat: 2021-10-30T10:00:00+08:00
labels:
  - go
  - coding
server:
  addr: 0.0.0.0
  port: 8000
`

cfg := common.Config{}

err := yaml.Unmarshal([]byte(data), &cfg)
if err != nil {
	log.Fatalf("error: %v", err)
}
log.Println(cfg)
```

## [Encoder](https://pkg.go.dev/gopkg.in/yaml.v3#Encoder)

[encoder.go](encoder.go):

```go
cfg := common.Config{
	V:         "1.0.0",
	CreatedAt: time.Now(),
	Labels:    []string{"go", "coding"},
	Server: common.ServerConfig{
		Addr: "0.0.0.0",
		Port: 8000,
	},
}

filename := "config.yaml"
log.Printf("open file: %s\n", filename)
f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
if err != nil {
	log.Fatal(err)
}
defer f.Close() // ignore error

enc := yaml.NewEncoder(f)
if err := enc.Encode(cfg); err != nil {
	log.Fatal(err)
}
log.Println("encode success! 🍺")
```

## [Decoder](https://pkg.go.dev/gopkg.in/yaml.v3#Decoder)

[decoder.go](decoder.go):

```go
filename := "config.yaml"
log.Printf("open file: %s\n", filename)
f, err := os.Open(filename)
if err != nil {
	log.Fatal(err)
}
defer f.Close() // ignore error

dec := yaml.NewDecoder(f)

var cfg common.Config
if err := dec.Decode(&cfg); err != nil {
	log.Fatal(err)
}
log.Printf("decode success! 🍺\n  %v\n", cfg)
```
