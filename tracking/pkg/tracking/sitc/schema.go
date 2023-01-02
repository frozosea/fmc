package sitc

type ContainerListEntity struct {
	ContainerNo    string `json:"containerNo"`
	MovementName   string `json:"movementName"`
	MovementCode   string `json:"movementCode"`
	MovementNameEn string `json:"movementNameEn"`
	EventPort      string `json:"eventPort"`
	EventDate      string `json:"eventDate"`
	VesselCode     string `json:"vesselCode"`
	VoyageNo       string `json:"voyageNo"`
}

type ContainerApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []*ContainerListEntity `json:"list"`
	} `json:"data"`
}

type BillNumberList1Entity struct {
	BlNo  string `json:"blNo"`
	Polen string `json:"polen"`
	Del   string `json:"del"`
	Delen string `json:"delen"`
	Pol   string `json:"pol"`
}

type BillNumberList2Entity struct {
	VesselName   string      `json:"vesselName"`
	VoyageNo     string      `json:"voyageNo"`
	VoyageLeg    string      `json:"voyageLeg"`
	PortFrom     string      `json:"portFrom"`
	PortFromName string      `json:"portFromName"`
	PortTo       string      `json:"portTo"`
	PortToName   string      `json:"portToName"`
	Eta          string      `json:"eta"`
	Etd          string      `json:"etd"`
	Atd          string      `json:"atd"`
	Agtb         string      `json:"agtb"`
	Cctd         interface{} `json:"cctd"`
	Cta          interface{} `json:"cta"`
	Ata          string      `json:"ata"`
	Ctd          interface{} `json:"ctd"`
	Agta         string      `json:"agta"`
	Ccta         string      `json:"ccta"`
	Atb          string      `json:"atb"`
	Agtd         string      `json:"agtd"`
	Ctb          interface{} `json:"ctb"`
	Cctb         string      `json:"cctb"`
}

type BillNumberList3Entity struct {
	RowNo          string `json:"rowNo"`
	TotalCount     string `json:"totalCount"`
	ContainerNo    string `json:"containerNo"`
	SealNo         string `json:"sealNo"`
	VoyageNo       string `json:"voyageNo"`
	CntrType       string `json:"cntrType"`
	Quantity       string `json:"quantity"`
	CntrSize       string `json:"cntrSize"`
	Weight         string `json:"weight"`
	Currentport    string `json:"currentport"`
	Movementname   string `json:"movementname"`
	Movementnameen string `json:"movementnameen"`
}

type BillNumberApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List1        []*BillNumberList1Entity `json:"list1"`
		ContainerNo  string                   `json:"containerNo"`
		BlNo         string                   `json:"blNo"`
		List3        []*BillNumberList3Entity `json:"list3"`
		List2        []*BillNumberList2Entity `json:"list2"`
		Movementcode []struct {
			MovementStatus string `json:"movementStatus"`
		} `json:"movementcode"`
	} `json:"data"`
}

type BillNumberInfoAboutContainerApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			Movementid     string `json:"movementid"`
			BlNo           string `json:"blNo"`
			ContainerNo    string `json:"containerNo"`
			Eventdate      string `json:"eventdate"`
			Portname       string `json:"portname"`
			Movementcode   string `json:"movementcode"`
			Movementname   string `json:"movementname"`
			Movementnameen string `json:"movementnameen"`
		} `json:"list"`
	} `json:"data"`
}
