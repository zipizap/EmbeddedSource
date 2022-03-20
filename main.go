package EmbeddedSource

import (
	"crypto/sha512"
	"embed"
	"encoding/hex"
	"io/fs"
	"os"
	"path/filepath"
)

type EmbeddedSource struct {
	Fs              embed.FS
	cachedlistfiles []string
}

func (es *EmbeddedSource) FilesList() (filenames []string, err error) {
	// filenames := []string
	// err := nil
	if len(es.cachedlistfiles) > 0 {
		// es.cachedlistfiles is already defined from a previous function-call
		filenames = es.cachedlistfiles
		return filenames, err
	}
	walk_err := fs.WalkDir(es.Fs, ".", func(a_file_path string, a_file_info fs.DirEntry, parent_err error) error {
		if parent_err != nil {
			return parent_err
		}
		// skip any directories entries, we just want to add files
		if a_file_info.IsDir() {
			return nil
		}

		filenames = append(filenames, a_file_path)
		return nil
	})
	if walk_err != nil {
		filenames = []string{}
		err = walk_err
		return filenames, err
	}
	// save filenames into es.cachedlistfiles to optimize future function-calls
	es.cachedlistfiles = filenames
	return filenames, err
}
func (es *EmbeddedSource) FileSha512(filename string) (sha512_as_hex string, err error) {
	file_contents, err := es.Fs.ReadFile(filename)
	if err != nil {
		sha512_as_hex = ""
		return sha512_as_hex, err
	}
	hasher := sha512.New()
	hasher.Write(file_contents)
	sha512_as_hex = hex.EncodeToString(hasher.Sum(nil))
	return sha512_as_hex, err
}
func (es *EmbeddedSource) FileRead(filename string) (file_contents []byte, err error) {
	file_contents, err = es.Fs.ReadFile(filename)
	return file_contents, err
}
func (es *EmbeddedSource) FilesExtract(destination_dir string) (err error) {
	// create destination_dir
	if err := os.MkdirAll(destination_dir, 0777); err != nil {
		return err
	}
	// create files inside destination_dir
	filenames, err := es.FilesList()
	if err != nil {
		return err
	}
	for _, a_filename := range filenames {
		a_dest_filepath := filepath.Join(destination_dir, a_filename)
		a_file_content, err := es.FileRead(a_filename)
		if err != nil {
			return err
		}
		err = os.WriteFile(a_dest_filepath, a_file_content, 0644)
		if err != nil {
			return err
		}
	}
	return err // err is nil
}
func (es *EmbeddedSource) FileInfo(filename string) (fileInfo fs.FileInfo, err error) {
	fileInfo, err = fs.Stat(es.Fs, filename)
	return fileInfo, err
}
