package fileReader

import (
	"bufio"
	"encoding/base64"
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
func (f *FileReader) GetSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return `\`
	case "darwin":
		return `/`
	case "linux":
		return `/`
	default:
		return ""
	}
}
func (f *FileReader) ReadImageToBase64(imageFilePath string) (string, error) {

	imgFile, err := os.Open(imageFilePath) // a QR code image

	if err != nil {
		return "", err
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	if _, err := fReader.Read(buf); err != nil {
		return "", err
	}

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return imgBase64Str, nil
}
