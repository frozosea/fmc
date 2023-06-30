package huaxin_schedule

import "time"

type Schedule struct {
	VESSEL        string      `json:"VESSEL"`
	VOYAGE        string      `json:"VOYAGE"`
	BARGEVESSEL   interface{} `json:"BARGE_VESSEL"`
	BARGEVOYAGE   interface{} `json:"BARGE_VOYAGE"`
	BARGEDISCNAME interface{} `json:"BARGE_DISC_NAME"`
	PORTLOADNAME  string      `json:"PORT_LOAD_NAME"`
	LOADPIERNAME  string      `json:"LOAD_PIER_NAME"`
	LOADETD       string      `json:"LOAD_ETD"`
	LOADETDNAME   string      `json:"LOAD_ETD_NAME"`
	PORTDISCNAME  string      `json:"PORT_DISC_NAME"`
	DISCPIERNAME  string      `json:"DISC_PIER_NAME"`
	DISCETA       string      `json:"DISC_ETA"`
	DISCETANAME   string      `json:"DISC_ETA_NAME"`
	TRANSITTIME   string      `json:"TRANSIT_TIME"`
}

type ServerResponse struct {
	Status        string      `json:"status"`
	ListSchedules []*Schedule `json:"listSchedules"`
}

type DataForScheduleRequest struct {
	lastVoyage       string
	lastPortUnlocode string
	etd              time.Time
}
