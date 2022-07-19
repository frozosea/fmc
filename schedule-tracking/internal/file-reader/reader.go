package file_reader

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type FileReader struct{}

func New() *FileReader {
	return &FileReader{}
}

func (f *FileReader) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
func (f *FileReader) GetFileName(filePath string) (string, error) {
	goos := runtime.GOOS
	var splitFilePath []string
	switch goos {
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
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf(`%s\%s`, dirName, filename)
	case "darwin":
		return fmt.Sprintf(`%s/%s`, dirName, filename)
	case "linux":
		return fmt.Sprintf(`%s/%s`, dirName, filename)
	default:
		return ""
	}
}
