# How to cross compile in Go

[go build]: https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
[go tool]: https://pkg.go.dev/cmd/dist

[go build][] cross compile on `linux/amd64`:

```zsh
env GOOS=linux GOARCH=amd64 go build [packages]
```

`$GOOS` and `$GOARCH` [described here](https://go.dev/doc/install/source#environment), or list them by [go tool][]:

```zsh
❯ go tool dist list
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/mips64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
windows/386
windows/amd64
windows/arm
windows/arm64
```
