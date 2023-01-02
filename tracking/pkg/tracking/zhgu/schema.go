package zhgu

type BookingApiResponseSchema struct {
	BackCode    string `json:"backCode"`
	BackMessage string `json:"backMessage"`
	Object      []interface{}
}

type ApiResponseSchema struct {
	BackCode    string `json:"backCode"`
	BackMessage string `json:"backMessage"`
	Object      []struct {
		CreatedByUser      interface{} `json:"createdByUser"`
		CreatedOffice      interface{} `json:"createdOffice"`
		CreatedDtmLoc      interface{} `json:"createdDtmLoc"`
		CreatedTimeZone    interface{} `json:"createdTimeZone"`
		UpdatedByUser      interface{} `json:"updatedByUser"`
		UpdatedOffice      interface{} `json:"updatedOffice"`
		UpdatedDtmLoc      interface{} `json:"updatedDtmLoc"`
		UpdatedTimeZone    interface{} `json:"updatedTimeZone"`
		RecordVersion      interface{} `json:"recordVersion"`
		RowStatus          int         `json:"rowStatus"`
		PrincipalGroupCode interface{} `json:"principalGroupCode"`
		BlNo               string      `json:"blNo"`
		TripNumber         string      `json:"tripNumber"`
		LineType           string      `json:"lineType"`
		VesselName         string      `json:"vesselName"`
		Voyage             string      `json:"voyage"`
		PortFrom           string      `json:"portFrom"`
		PortFromName       string      `json:"portFromName"`
		PortTo             string      `json:"portTo"`
		PortToName         string      `json:"portToName"`
		Etd                string      `json:"etd"`
		Atd                string      `json:"atd"`
		Eta                string      `json:"eta"`
		Ata                string      `json:"ata"`
		DepartureTime      string      `json:"departureTime"`
		ArrivalTime        string      `json:"arrivalTime"`
	} `json:"object"`
}
