package feso

import "time"

type fesoLastEventSchema struct {
	Time               string      `json:"time"`
	Operation          string      `json:"operation,omitempty"`
	OperationName      string      `json:"operationName"`
	OperationNameLatin string      `json:"operationNameLatin"`
	LocId              interface{} `json:"locId"`
	LocName            string      `json:"locName"`
	LocNameLatin       string      `json:"locNameLatin"`
	LocIdTo            int         `json:"locIdTo,omitempty"`
	LocNameTo          string      `json:"locNameTo,omitempty"`
	LocNameLatinTo     string      `json:"locNameLatinTo,omitempty"`
	EtCode             *string     `json:"etCode,omitempty"`
	TransportType      *string     `json:"transportType,omitempty"`
	Vessel             string      `json:"vessel,omitempty"`
}

type location struct {
	Id              int         `json:"id"`
	Text            string      `json:"text"`
	TextLatin       string      `json:"textLatin"`
	ParentText      interface{} `json:"parentText"`
	ParentTextLatin interface{} `json:"parentTextLatin"`
	Country         string      `json:"country"`
	CountryLatin    string      `json:"countryLatin"`
	LtCode          string      `json:"ltCode"`
	SoftshipCode    string      `json:"softshipCode"`
	Code            interface{} `json:"code"`
	Type            string      `json:"type"`
	Info            []struct {
		Type          string    `json:"type"`
		TransportType string    `json:"transportType"`
		Date          time.Time `json:"date"`
	} `json:"info"`
	To     string `json:"to,omitempty"`
	Passed bool   `json:"passed,omitempty"`
	From   string `json:"from,omitempty"`
	Here   bool   `json:"here,omitempty"`
}

type T struct {
	Container       string    `json:"container"`
	Time            time.Time `json:"time"`
	ContainerCTCode string    `json:"containerCTCode"`
	ContainerOwner  string    `json:"containerOwner"`
	RequestDate     time.Time `json:"requestDate"`
	LatLng          struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"latLng"`
	Locations  []*location            `json:"locations"`
	LastEvents []*fesoLastEventSchema `json:"lastEvents"`
}
type ResponseSchema struct {
	Container       string    `json:"container"`
	Time            time.Time `json:"time"`
	ContainerCTCode string    `json:"containerCTCode"`
	ContainerOwner  string    `json:"containerOwner"`
	RequestDate     time.Time `json:"requestDate"`
	LatLng          struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"latLng"`
	Locations  []*location            `json:"locations"`
	LastEvents []*fesoLastEventSchema `json:"lastEvents"`
}

type FullApiResponseSchema struct {
	RequestKey string
	Containers []string
	Missing    []string
}
