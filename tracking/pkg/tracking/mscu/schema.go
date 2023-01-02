package mscu

type Event struct {
	Order       int      `json:"Order"`
	Date        string   `json:"Date"`
	Location    string   `json:"Location"`
	Description string   `json:"Description"`
	Detail      []string `json:"Detail"`
}
type ContainerInfo struct {
	ContainerNumber string   `json:"ContainerNumber"`
	PodEtaDate      string   `json:"PodEtaDate"`
	ContainerType   string   `json:"ContainerType"`
	LatestMove      string   `json:"LatestMove"`
	Events          []*Event `json:"Events"`
}

type GeneralTrackingInfo struct {
	ShippedFrom          string        `json:"ShippedFrom"`
	ShippedTo            string        `json:"ShippedTo"`
	PortOfLoad           string        `json:"PortOfLoad"`
	PortOfDischarge      string        `json:"PortOfDischarge"`
	Transshipments       []interface{} `json:"Transshipments"`
	PriceCalculationDate string        `json:"PriceCalculationDate"`
	FinalPodEtaDate      string        `json:"FinalPodEtaDate"`
}

type BillOfLading struct {
	BillOfLadingNumber  string               `json:"BillOfLadingNumber"`
	NumberOfContainers  int                  `json:"NumberOfContainers"`
	GeneralTrackingInfo *GeneralTrackingInfo `json:"GeneralTrackingInfo"`
	ContainersInfo      []*ContainerInfo     `json:"ContainersInfo"`
}

type ApiResponse struct {
	IsSuccess bool `json:"IsSuccess"`
	Data      struct {
		TrackingType          string          `json:"TrackingType"`
		TrackingTitle         string          `json:"TrackingTitle"`
		TrackingNumber        string          `json:"TrackingNumber"`
		CurrentDate           string          `json:"CurrentDate"`
		PriceCalculationLabel string          `json:"PriceCalculationLabel"`
		TrackingResultsLabel  string          `json:"TrackingResultsLabel"`
		BillOfLadings         []*BillOfLading `json:"BillOfLadings"`
	} `json:"Data"`
}
