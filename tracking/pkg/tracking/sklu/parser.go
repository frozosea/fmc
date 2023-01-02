package sklu

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
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
		return nil, errors.New("lens of arrays (times,operations,locations,vessels) are not equal")
	}

	for index, v := range vessels {
		if v == containerNo {
			vessels[index] = ""
		}
	}

	for index := 0; index < len(operations); index++ {
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
