# Start [Cobra][]

[Cobra]: https://github.com/spf13/cobra

- [User Guide](https://github.com/spf13/cobra/blob/master/user_guide.md)

## Install

Install the `cobra` generator executable along with the library and its dependencies:

```zsh
cd start-golang/_/start-cobra
go get -u github.com/spf13/cobra/cobra
```

## Getting Started

Cobra-based app's organization:

```zsh
❯ tree start-cobra -aF --dirsfirst
start-cobra
├── cmd/
│   ├── config.go
│   ├── create.go
│   ├── root.go
│   └── serve.go
└── main.go
```

### Using the Cobra Generator

#### `cobra init`

```zsh
cobra init --pkg-name github.com/ikuokuo/start-golang/_/start-cobra
```

```zsh
❯ go build
❯ ./start-cobra
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
```

#### `cobra add`

Add commands:

- app serve
- app config
- app config create

```zsh
cobra add serve
cobra add config
cobra add create -p 'configCmd'
```

```zsh
❯ go build
❯ ./start-cobra
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  start-cobra [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  config      A brief description of your command
  help        Help about any command
  serve       A brief description of your command

Flags:
      --config string   config file (default is $HOME/.start-cobra.yaml)
  -h, --help            help for start-cobra
  -t, --toggle          Help message for toggle

Use "start-cobra [command] --help" for more information about a command.
```

#### Configuring the cobra generator

```bash
cat <<-EOF > ~/.cobra.yaml
author: ikuokuo <ikuokuo@hotmail.com>
license: MIT
EOF
```

### Using the Cobra Library

To manually implement Cobra [see here](https://cobra.dev/#using-the-cobra-library).
