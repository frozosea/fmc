package sklu

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"regexp"
	"strings"
	"time"
)

type ApiParser struct {
	datetime datetime.IDatetime
}

func NewApiParser(datetime datetime.IDatetime) *ApiParser {
	return &ApiParser{datetime: datetime}
}

func (a *ApiParser) Get(response *ApiResponse) *ContainerInfo {
	//2022-11-06
	eta, err := a.datetime.Strptime(response.ETA, "%Y-%m-%d")
	if err != nil {
		return nil
	}
	return &ContainerInfo{
		BillNo:        response.BKNO,
		Eta:           eta,
		ContainerSize: strings.ToUpper(response.CNTR),
		Unlocode:      strings.ToUpper(response.POD),
	}

}

type TableParser struct {
}

func NewTableParser() *TableParser {
	return &TableParser{}
}
func (t *TableParser) Get(doc *goquery.Document) ([]string, error) {
	var events []string
	doc.Find("#wrapper > div > div:nth-child(6) > div.panel-body > div > table").Find("td").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "" {
			events = append(events, text)
		}
	})
	if len(events) == 0 {
		return nil, errors.New("couldn't parse table in this html")
	}
	return events, nil
}

type OperationsParser struct {
}

func NewOperationsParser() *OperationsParser {
	return &OperationsParser{}
}

func (o *OperationsParser) Get(doc *goquery.Document) ([]string, error) {
	var events []string
	doc.Find(".splitTable").Find(".firstTh").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text != "" {
			events = append(events, strings.ToTitle(text))
		}
	})
	if len(events) == 0 {
		return nil, errors.New("couldn't parse operations in this html")
	}
	return events, nil
}

type timeParser struct {
	datetime datetime.IDatetime
}

func newTimeParser(datetime datetime.IDatetime) *timeParser {
	return &timeParser{datetime: datetime}
}
func (t *timeParser) get(timeString string) (time.Time, error) {
	//2022-09-08 THU 21:30
	splitTime := strings.Split(timeString, " ")
	return t.datetime.Strptime(fmt.Sprintf(`%s %s`, splitTime[0], splitTime[2]), `%Y-%m-%d %H:%M`)
}

type InfoAboutMovingParser struct {
	tableParser      *TableParser
	operationsParser *OperationsParser
	timeParser       *timeParser
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{
		timeParser:       newTimeParser(datetime),
		operationsParser: NewOperationsParser(),
		tableParser:      NewTableParser(),
	}
}

func (i *InfoAboutMovingParser) Get(document *goquery.Document, containerNo string) ([]*tracking.Event, error) {
	var events []*tracking.Event
	table, err := i.tableParser.Get(document)
	if err != nil {
		return nil, err
	}
	lastIndex := len(table) - 1
	times := tracking.StringSliceWithStep(table, 2, lastIndex, 3)
	operations, err := i.operationsParser.Get(document)
	if err != nil {
		return nil, err
	}
	locations := tracking.StringSliceWithStep(table, 1, lastIndex, 3)
	vessels := tracking.StringSliceWithStep(table, 0, lastIndex, 3)
	if len(times) != len(operations) || len(locations) != len(operations) || len(vessels) != len(operations) {
		return nil, &LensNotEqualError{}
	}
	for index, v := range vessels {
		if v == containerNo {
			vessels[index] = ""
		}
	}

	for index := 0; index < len(times); index++ {
		eventTime, err := i.timeParser.get(times[index])
		if err != nil {
			continue
		}
		events = append(events, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToTitle(operations[index]),
			Location:      strings.ToUpper(locations[index]),
			Vessel:        strings.ToUpper(vessels[index]),
		})
	}
	return events, nil
}

type ContainerNumberParser struct {
}

func NewContainerNumberParser() *ContainerNumberParser {
	return &ContainerNumberParser{}
}

func (c *ContainerNumberParser) Get(doc *goquery.Document) (string, error) {
	elem := doc.Find("#wrapper > div > div:nth-child(5) > div.panel-body > div > div.form-group.table-responsive > div:nth-child(2) > table > tbody > tr:nth-child(1) > td:nth-child(1)")
	text := elem.Text()
	if text == "" {
		return "", errors.New("cannot find container number")
	}
	return text, nil
}

type EtaParser struct {
	timeParser *timeParser
	dt         datetime.IDatetime
}

func NewEtaParser(dt datetime.IDatetime) *EtaParser {
	return &EtaParser{timeParser: newTimeParser(dt), dt: dt}
}
func (e *EtaParser) clearRawText(rawText string) (string, error) {
	re, err := regexp.Compile("[\\t\\n\\r\\f\\v]")
	if err != nil {
		return "", err
	}
	return strings.Join(strings.Fields(strings.Trim(re.ReplaceAllString(rawText, ""), " ")), " "), nil
}
func (e *EtaParser) GetEta(doc *goquery.Document) (time.Time, error) {
	firstDivLength := doc.Find("#wrapper > div > div:nth-child(4) > div.panel-body > div > div.form-both.form-group > div").Children().Length()
	toUlSelector := fmt.Sprintf("#wrapper > div > div:nth-child(4) > div.panel-body > div > div.form-both.form-group > div:nth-child(%d) > ul > li.col-sm-12.col-md-8", firstDivLength)
	lastDivSelectorLength := doc.Find(toUlSelector).Children().Length()
	selector := fmt.Sprintf(`%s > div:nth-child(%d)`, toUlSelector, lastDivSelectorLength)
	rawText := doc.Find(selector).Text()
	re, err := regexp.Compile("\\d{4}-\\d{1,2}-\\d{1,2} \\d{1,2}:\\d{2}")
	if err != nil {
		return time.Time{}, err
	}
	text, err := e.clearRawText(rawText)
	if err != nil {
		return time.Time{}, err
	}
	textTime := strings.TrimSpace(re.FindAllString(text, 1)[0])
	return e.dt.Strptime(textTime, "%Y-%m-%d %H:%M")
}
