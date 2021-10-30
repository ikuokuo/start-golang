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
