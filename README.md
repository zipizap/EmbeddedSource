# EmbeddedSource
GoLang module to easily embed-and-latter-extract main.go source-code into the compiled-binary 



# How to use

## Add into your main.go:
```
import (
  ...
	"embed"
	"github.com/zipizap/EmbeddedSource"
)

...
// Embed files: ./*.go ./go.sum ./go.mod
// Does not embed subdirs:
//     ./a.go              is embedded
//     ./mysubdir/a.go     is not embedded
//
//go:embed go.mod go.sum *.go
var embfs embed.FS
var EmbSrc = &EmbeddedSource.EmbeddedSource{
	Fs: embfs,
}

...
func main() {
  ...
  // When compiling, golang will automatically embed the *.go,go.mod,go.sum files into binary
  // When the program is executing, you can use EmbSrc to extract the source-files into subdir "extracted":
	err = EmbSrc.FilesExtract("extracted")
	if err != nil {
		log.Fatalln(err)
	}
}

```

And run `go get -u github.com/zipizap/EmbeddedSource`

Now, when you latter compile your code, the golang compiler will embed the files `*.go  go.mod  go.sum` inside the compiled binary, and the program can then use `EmbSrc` to extract the embeded source-files to a a subdirectory.

This can be usefull to store the source-code of small-go-programs with only a few .go files, and latter "extract" that source-code from the binary itself by using for example a "--flag"





  
