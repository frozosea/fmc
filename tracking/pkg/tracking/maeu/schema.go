package maeu

type Origin struct {
	Terminal    string `json:"terminal"`
	GeoSite     string `json:"geo_site"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	GeoidCity   string `json:"geoid_city"`
	SiteType    string `json:"site_type"`
}
type Destination struct {
	Terminal    string `json:"terminal"`
	GeoSite     string `json:"geo_site"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	GeoidCity   string `json:"geoid_city"`
	SiteType    string `json:"site_type"`
}
type Event struct {
	Activity     string `json:"activity"`
	Stempty      bool   `json:"stempty"`
	Actfor       string `json:"actfor"`
	VesselName   string `json:"vessel_name"`
	VoyageNum    string `json:"voyage_num"`
	VesselNum    string `json:"vessel_num"`
	ActualTime   string `json:"actual_time,omitempty"`
	RkemMove     string `json:"rkem_move,omitempty"`
	IsCancelled  bool   `json:"is_cancelled,omitempty"`
	IsCurrent    bool   `json:"is_current"`
	ExpectedTime string `json:"expected_time,omitempty"`
}

type Location struct {
	Terminal    string   `json:"terminal"`
	GeoSite     string   `json:"geo_site"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	Country     string   `json:"country"`
	CountryCode string   `json:"country_code"`
	GeoidCity   string   `json:"geoid_city"`
	SiteType    string   `json:"site_type"`
	Events      []*Event `json:"events"`
}
type Latest struct {
	ActualTime  string `json:"actual_time"`
	Activity    string `json:"activity"`
	Stempty     bool   `json:"stempty"`
	Actfor      string `json:"actfor"`
	GeoSite     string `json:"geo_site"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}
type Container struct {
	ContainerNum     string      `json:"container_num"`
	ContainerSize    string      `json:"container_size"`
	ContainerType    string      `json:"container_type"`
	IsoCode          string      `json:"iso_code"`
	Operator         string      `json:"operator"`
	Locations        []*Location `json:"locations"`
	EtaFinalDelivery string      `json:"eta_final_delivery"`
	Latest           *Latest     `json:"latest"`
	Status           string      `json:"status"`
}

type ApiResponse struct {
	IsContainerSearch bool         `json:"isContainerSearch"`
	Origin            *Origin      `json:"origin"`
	Destination       *Destination `json:"destination"`
	Containers        []*Container `json:"containers"`
}
