# go-shellhook

shell hook library for go programs inspired by direnv

## usage

get the latest version

```bash
go get -u github.com/patrickhuber/go-shellhook
```

For a full example see the [example](cmd/example/README.md)

## note

This library executes in the shell prompt and does an export of variables in the current shell session. These shell variables will persist between commands. 