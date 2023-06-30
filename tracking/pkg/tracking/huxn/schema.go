package huxn

type ContainerTrackingOneEvent struct {
	CONNO        string `json:"CON_NO"`
	TGCODE       string `json:"TG_CODE"`
	DYNTYPE      string `json:"DYN_TYPE"`
	DYNTIME      string `json:"DYN_TIME"`
	DYNTIMENAME  string `json:"DYN_TIME_NAME"`
	PORTFULLNAME string `json:"PORT_FULL_NAME"`
	PLACENAME    string `json:"PLACE_NAME"`
	VESSELVOYAGE string `json:"VESSEL_VOYAGE"`
}

type ContainerTrackingResponse struct {
	Status       string                       `json:"status"`
	ListDynamics []*ContainerTrackingOneEvent `json:"listDynamics"`
}
