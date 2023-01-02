package cosu

type Container struct {
	ContainerUuid       string   `json:"containerUuid"`
	ContainerNumber     string   `json:"containerNumber"`
	ContainerType       string   `json:"containerType"`
	GrossWeight         string   `json:"grossWeight"`
	PiecesNumber        int      `json:"piecesNumber"`
	Label               string   `json:"label"`
	SealNumber          string   `json:"sealNumber"`
	Location            string   `json:"location"`
	LocationDateTime    string   `json:"locationDateTime"`
	Transportation      string   `json:"transportation"`
	Flag                string   `json:"flag"`
	RailRef             string   `json:"railRef"`
	InlandMvId          string   `json:"inlandMvId"`
	ContainerLocation   string   `json:"containerLocation"`
	IsShow              bool     `json:"isShow"`
	PolEtd              string   `json:"polEtd"`
	PolAtd              string   `json:"polAtd"`
	PodEta              string   `json:"podEta"`
	PodAta              string   `json:"podAta"`
	TransportId         string   `json:"transportId"`
	Pol                 string   `json:"pol"`
	Pod                 string   `json:"pod"`
	HsCode              []string `json:"hsCode"`
	IsNorthAmericaRails bool     `json:"isNorthAmericaRails"`
}

type CircleStatus struct {
	Uuid                  string `json:"uuid"`
	ContainerNumber       string `json:"containerNumber"`
	ContainerNumberStatus string `json:"containerNumberStatus"`
	Location              string `json:"location"`
	TimeOfIssue           string `json:"timeOfIssue"`
	Transportation        string `json:"transportation"`
	PolEtd                string `json:"polEtd"`
	PolAtd                string `json:"polAtd"`
	PodEta                string `json:"podEta"`
	PodAta                string `json:"podAta"`
	TransportId           string `json:"transportId"`
	Pol                   string `json:"pol"`
	Pod                   string `json:"pod"`
}

type ContainerHistory struct {
	Uuid                  string `json:"uuid"`
	ContainerNumber       string `json:"containerNumber"`
	ContainerNumberStatus string `json:"containerNumberStatus"`
	Location              string `json:"location"`
	TimeOfIssue           string `json:"timeOfIssue"`
	Transportation        string `json:"transportation"`
	PolEtd                string `json:"polEtd"`
	PolAtd                string `json:"polAtd"`
	PodEta                string `json:"podEta"`
	PodAta                string `json:"podAta"`
	TransportId           string `json:"transportId"`
	Pol                   string `json:"pol"`
	Pod                   string `json:"pod"`
}

type ApiResponseSchema struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content struct {
			Containers []struct {
				Container             *Container          `json:"container"`
				ContainerCircleStatus []*CircleStatus     `json:"containerCircleStatus"`
				ContainerHistorys     []*ContainerHistory `json:"containerHistorys"`
			} `json:"containers"`
			NotFound string `json:"notFound"`
		} `json:"content"`
	} `json:"data"`
}

type EtaApiResponseSchema struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content string `json:"content"`
	} `json:"data"`
}
