package htmlTemplateReader

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	reader "user-api/pkg/fileReader"
)

type HTMLReader struct {
}

func New() *HTMLReader {
	return &HTMLReader{}
}

func (h *HTMLReader) ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
func (h *HTMLReader) GetStringHtml(dir, filename string, data interface{}) (string, error) {
	f := reader.New()
	fmt.Println(filename)
	fileByteBody, err := f.ReadFile(f.GetFileNameByDirNameAndFilename(dir, filename))
	if err != nil {
		return "", err
	}
	t, err := template.New(filename).Parse(string(fileByteBody))
	if err != nil {
		return "", err
	}
	var body bytes.Buffer
	t = t.Lookup(filename)
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}
	return body.String(), nil
}
