# EmbeddedSource
GoLang module to easily embed-and-latter-extract main.go source-code into the compiled-binary 

See the [example](example/README.md)

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
  // When the program is executing, you can use EmbSrc to extract the source-files into subdir "extracted":
	if err = EmbSrc.FilesExtract("extracted"); err != nil {
		log.Fatalln(err)
	}
}

```

And import it with `go get -u github.com/zipizap/EmbeddedSource`

Now, when you latter `go run/build` the code, the golang compiler will embed the files `*.go  go.mod  go.sum` inside the compiled binary, and the program can then use `EmbSrc` to extract the embeded source-files to a subdirectory (ex: "extracted")

This can be usefull to store the source-code of small-go-programs (with only a few .go files), and be able to latter "extract" the source-code from the binary itself, by using for example a "--flag"





  
