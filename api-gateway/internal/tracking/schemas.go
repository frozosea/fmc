package tracking

type HttpRequestSchema struct {
	Number  string `json:"number" validate:"min=10,max=28,regexp=[a-zA-Z]{2,}\d{5,}"`
	Scac    string `json:"scac"`
	Country string `json:"country"`
}
