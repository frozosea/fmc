package unlocodesParser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	data, err := os.ReadFile("test_data/page.txt")
	if err != nil {
		panic(err)
	}

	var expectedData = []string{"CNCGO", "CNCGU", "CNCKG", "CNCNA", "CNCOZ", "CNCTU", "CNDLC", "CNFOC", "CNFUG", "CNHAK", "CNHUA", "CNJIA", "CNJIU", "CNLYG", "CNNGB", "CNNKG", "CNNSA", "CNNTG", "CNQZH", "CNQZJ", "CNRZH", "CNSFE", "CNSHA", "CNSHK", "CNSIA", "CNSWA", "CNSZX", "CNTAC", "CNTAO", "CNTJN", "CNWHI", "CNWNZ", "CNWUH", "CNXMN", "CNXNG", "CNYIC", "CNYIK", "CNYIU", "CNYPG", "CNZAP", "CNZHA", "CNZHE", "CNZJG", "CNZSN", "JPNGO", "JPOSA", "JPSBS", "JPSHI", "JPTYO", "JPUKB", "JPYOK", "KRINC", "KRPTK", "KRPUS", "RUKRS", "RULED", "RUMOS", "RUNAK", "RUOVB", "RUVVO", "RUVYP", "RUYEK"}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	p := NewParser()
	unlocodes, err := p.Parse(doc, "CNBAY")
	assert.NoError(t, err)
	assert.Equal(t, expectedData, unlocodes)
}
