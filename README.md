# go-shellhook

shell hook library for go programs inspired by direnv

## usage

get the latest version

```bash
go get -u github.com/patrickhuber/shellhook
```

```go
func run() error{
  
  executable, err := os.Executable()
  if err != nil {
    return err
  }

  result, err := shellhook.Hook(
    sh, 
    &shellhook.Metadata{
		Executable: executable,
		Name:       "test",
		Args:       []string{"export", sh.Name()},
    })
  
  if err != nil{
    return err
  }

  fmt.Println(result)
  return nil
}
```