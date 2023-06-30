package unlocodesParser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(doc *goquery.Document, portLoadUnlocode string) ([]string, error) {
	var mu sync.Mutex
	var unlocodes []string
	doc.Find("#cmbPortLoad").Children().Each(func(i int, selection *goquery.Selection) {
		value, exists := selection.Attr("value")
		if !exists || value == "" || value == " " || value == portLoadUnlocode {
			return
		}
		defer mu.Unlock()
		mu.Lock()
		unlocodes = append(unlocodes, strings.Trim(strings.ToUpper(value), " "))
	})

	if len(unlocodes) == 0 {
		return nil, errors.New("no unlocodes")
	}

	return unlocodes, nil
}
