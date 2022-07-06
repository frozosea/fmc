package line

import (
	"io"
)

type BaseLine struct {
	Scac     string
	FullName string
}
type WithByteImage struct {
	BaseLine
	Image io.Reader
}
type AddRepoLine struct {
	BaseLine
	ImageUrl string
}
type Line struct {
	Id int
	BaseLine
	ImageUrl string
}
