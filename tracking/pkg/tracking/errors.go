package tracking

import "fmt"

type NotThisLineException struct {
}

func NewNotThisLineException() *NotThisLineException {
	return &NotThisLineException{}
}
func (n *NotThisLineException) Error() string {
	return "Not this shipping line!"
}

type NoScacException struct {
}

func NewNoScacException() *NoScacException {
	return &NoScacException{}
}
func (n *NoScacException) Error() string {
	return "no scac"
}

type NumberNotFoundException struct {
	number string
}

func NewNumberNotFoundException(number string) *NumberNotFoundException {
	return &NumberNotFoundException{number: number}
}

func (n *NumberNotFoundException) Error() string {
	return fmt.Sprintf(`cannot find data with %s`, n.number)
}
