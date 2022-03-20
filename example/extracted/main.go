package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/zipizap/EmbeddedSource"
)

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

func main() {
	if len(os.Args) > 1 &&
		os.Args[1] == "--source-files-extract" {
		// FilesExtract all files into subdir "extracted"
		err := EmbSrc.FilesExtract("extracted")
		if err != nil {
			log.Fatalln(err)
		}
	} else if len(os.Args) > 1 &&
		os.Args[1] == "--source-files-info" {
		filenames, err := EmbSrc.FilesList()
		if err != nil {
			log.Fatalln(err)
		}
		// List all files names, sha512-hash and byte-size
		for _, a_filename := range filenames {
			// FileSha512
			a_sha512, err := EmbSrc.FileSha512(a_filename)
			if err != nil {
				log.Fatalln(err)
			}
			// FileInfo: .Name(), .Size(), .Mode(), ModTime(), IsDir()
			a_file_info, err := EmbSrc.FileInfo(a_filename)
			if err != nil {
				log.Fatalln(err)
			}
			a_file_size := a_file_info.Size()
			log.Println(a_filename, "\t", a_file_size, "\t", a_sha512)
		}
	} else {
		fmt.Printf("Usage: %s  --source-files-extract|--source-files-info \n", os.Args[0])
	}

}
