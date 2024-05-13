For very simple golang programs - without arguments/flags parsing - this can be very usefull.


1. Copy file `embeddedSourceStub.go` 
2. Call `embeddedSourceStub()` in  main.go:
```
func main() {
	embeddedSourceStub()
  ...
```

Now your compiled program will react to arguments `-h|--help` `--source-files-extract` `--source-files-info`
