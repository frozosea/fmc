package tracking

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ru string
type Other string
type ipApiResponse struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}
type util struct {
}

func (u *util) sendRequest(ip string) (*ipApiResponse, error) {
	url := fmt.Sprintf(`http://ip-api.com/json/%s`, ip)
	client := &http.Client{}
	response, err := client.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var s ipApiResponse
	if err := json.Unmarshal(body, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
func (u *util) getCountry(ip string) Ru {
	//response, err := u.sendRequest(ip)
	//if err != nil {
	//	return "OTHER"
	//}
	//countryCode := strings.ToUpper(response.CountryCode)
	//if countryCode == "RU" {
	//	return "RU"
	//}
	return "OTHER"
}
