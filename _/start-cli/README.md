# Start [cli][]

[cli]: https://github.com/urfave/cli

- [Usage v2](https://github.com/urfave/cli/blob/master/docs/v2/manual.md)

## Install

```zsh
cd start-golang/_/start-cli
go mod init github.com/ikuokuo/start-golang/_/start-cli

go get -u github.com/urfave/cli/v2
```

## Getting Started

- [main.go](main.go)

```zsh
go build .
./start-cli -h
./start-cli --meeting 2019-08-12T15:04:05 world
```
