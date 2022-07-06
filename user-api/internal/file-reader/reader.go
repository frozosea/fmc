package file_reader

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (f *FileReader) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
func (f *FileReader) GetFileName(filePath string) (string, error) {
	os := runtime.GOOS
	var splitFilePath []string
	switch os {
	case "windows":
		splitFilePath = strings.Split(filePath, `\`)
	case "darwin":
		splitFilePath = strings.Split(filePath, `/`)
	case "linux":
		splitFilePath = strings.Split(filePath, `/`)
	default:
		return "", errors.New("unsupported system")
	}
	return splitFilePath[len(splitFilePath)-1], nil
}
func (f *FileReader) GetFileNameByDirNameAndFilename(dirName string, filename string) string {
	os := runtime.GOOS
	var fullPath string
	switch os {
	case "windows":
		fullPath = fmt.Sprintf(`%s\%s`, dirName, filename)
	case "darwin":
		fullPath = fmt.Sprintf(`%s/%s`, dirName, filename)
	case "linux":
		fullPath = fmt.Sprintf(`%s/%s`, dirName, filename)
	default:
		return ""
	}
	return fullPath
}
