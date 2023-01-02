package sklu

import "time"

type ApiResponse struct {
	BKNO string      `json:"BKNO,omitempty"`
	CNTR string      `json:"CNTR,omitempty"`
	POR  interface{} `json:"POR,omitempty"`
	POL  string      `json:"POL,omitempty"`
	POD  string      `json:"POD,omitempty"`
	DLV  string      `json:"DLV,omitempty"`
	VSL  string      `json:"VSL,omitempty"`
	VYG  string      `json:"VYG,omitempty"`
	ETD  string      `json:"ETD,omitempty"`
	ETA  string      `json:"ETA,omitempty"`
}

type ContainerInfo struct {
	BillNo        string
	Eta           time.Time
	ContainerSize string
	Unlocode      string
}
