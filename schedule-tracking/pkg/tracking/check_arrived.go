package tracking

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/corpix/uarand"
	"schedule-tracking/pkg/requests"
	"strings"
)

type IsArrived bool

type IArrivedChecker interface {
	CheckContainerArrived(result ContainerNumberResponse) IsArrived
	CheckBillNoArrived(result BillNumberResponse) IsArrived
}

type skluArrivedChecker struct{}

func newSkluArrivedChecker() *skluArrivedChecker {
	return &skluArrivedChecker{}
}

func (s *skluArrivedChecker) checkInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, item := range result {
		if strings.Contains(strings.ToLower(item.OperationName), "arrival") &&
			!strings.Contains(strings.ToLower(item.OperationName), strings.ToLower("Arrival(T/S)")) &&
			!strings.Contains(strings.ToLower(item.OperationName), "scheduled") {
			return true
		}
	}
	return false
}
func (s *skluArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return s.checkInfoAboutMoving(result.InfoAboutMoving)
}
func (s *skluArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return s.checkInfoAboutMoving(result.InfoAboutMoving)
}

type fesoArrivedChecker struct {
}

func newFesoArrivedChecker() *fesoArrivedChecker {
	return &fesoArrivedChecker{}
}
func (f *fesoArrivedChecker) getIndex(item interface{}, s []BaseInfoAboutMoving) int {
	for index, v := range s {
		if v == item {
			return index
		}
	}
	return -1
}
func (f *fesoArrivedChecker) contains(s []BaseInfoAboutMoving, operation string) bool {
	for _, a := range s {
		if strings.EqualFold(a.OperationName, operation) {
			return true
		}
	}
	return false
}
func (f *fesoArrivedChecker) checkArrivedByInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, v := range result {
		if strings.EqualFold(v.OperationName, "ETA") {
			index := f.getIndex(v, result)
			if index != -1 {
				return index != len(result)-1
			}
		}

		if strings.EqualFold(v.Location, "MAGISTRAL") {
			return true
		}
	}
	if f.contains(result, "ETS") && !f.contains(result, "ETA") {
		return true
	}
	return false
}
func (f *fesoArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return f.checkArrivedByInfoAboutMoving(result.InfoAboutMoving)
}
func (f *fesoArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return f.checkArrivedByInfoAboutMoving(result.InfoAboutMoving)
}

type mscuArrivedChecker struct{}

func newMscuArrivedChecker() *mscuArrivedChecker {
	return &mscuArrivedChecker{}
}

func (m *mscuArrivedChecker) checkInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, v := range result {
		if strings.EqualFold(v.OperationName, "Empty to Shipper") {
			return true
		}
	}
	return false
}
func (m *mscuArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return m.checkInfoAboutMoving(result.InfoAboutMoving)
}
func (m *mscuArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return m.checkInfoAboutMoving(result.InfoAboutMoving)
}

type oneyArrivedChecker struct{}

func newOneyArrivedChecker() *oneyArrivedChecker {
	return &oneyArrivedChecker{}
}

func (o *oneyArrivedChecker) checkInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, v := range result {
		if strings.Contains(strings.ToLower(v.OperationName), strings.ToLower("Arrival at Port of Discharging")) ||
			strings.Contains(strings.ToLower(v.OperationName), strings.ToLower("Empty Container Returned from Customer")) ||
			strings.Contains(strings.ToLower(v.OperationName), strings.ToLower("at Port of Discharging")) ||
			strings.Contains(strings.ToLower(v.OperationName), strings.ToLower("POD")) ||
			strings.Contains(strings.ToLower(v.OperationName), strings.ToLower("Empty Container Release to Shipper")) {
			return true
		}
	}
	return false
}
func (o *oneyArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return o.checkInfoAboutMoving(result.InfoAboutMoving)
}
func (o *oneyArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return o.checkInfoAboutMoving(result.InfoAboutMoving)
}

type cosuArrivedChecker struct{}

func newCosuArrivedChecker() *cosuArrivedChecker {
	return &cosuArrivedChecker{}
}

func (c *cosuArrivedChecker) checkInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, v := range result {
		if strings.EqualFold(v.OperationName, "Discharged at Last POD") || strings.EqualFold(v.OperationName, "Empty Equipment Returned") {
			return true
		}
	}
	return false
}
func (c *cosuArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return c.checkInfoAboutMoving(result.InfoAboutMoving)
}
func (c *cosuArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return c.checkInfoAboutMoving(result.InfoAboutMoving)
}

type maeuResponse struct {
	IsContainerSearch bool `json:"isContainerSearch"`
	Origin            struct {
		Terminal    string `json:"terminal"`
		GeoSite     string `json:"geo_site"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		GeoidCity   string `json:"geoid_city"`
		SiteType    string `json:"site_type"`
	} `json:"origin"`
	Destination struct {
		Terminal    string `json:"terminal"`
		GeoSite     string `json:"geo_site"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		GeoidCity   string `json:"geoid_city"`
		SiteType    string `json:"site_type"`
	} `json:"destination"`
	Containers []struct {
		ContainerNum  string `json:"container_num"`
		ContainerSize string `json:"container_size"`
		ContainerType string `json:"container_type"`
		IsoCode       string `json:"iso_code"`
		Operator      string `json:"operator"`
		Locations     []struct {
			Terminal    string `json:"terminal"`
			GeoSite     string `json:"geo_site"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			GeoidCity   string `json:"geoid_city"`
			SiteType    string `json:"site_type"`
			Events      []struct {
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
			} `json:"events"`
		} `json:"locations"`
		EtaFinalDelivery string `json:"eta_final_delivery"`
		Latest           struct {
			ActualTime  string `json:"actual_time"`
			Activity    string `json:"activity"`
			Stempty     bool   `json:"stempty"`
			Actfor      string `json:"actfor"`
			GeoSite     string `json:"geo_site"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
		} `json:"latest"`
		Status string `json:"status"`
	} `json:"containers"`
}

type IMaeuRequest interface {
	Get(number string) (*maeuResponse, error)
}
type maeuRequest struct {
	r *requests.Request
}

func newMaeuRequest() *maeuRequest {
	return &maeuRequest{r: requests.New()}
}

func (m *maeuRequest) Get(number string) (*maeuResponse, error) {
	headers := map[string]string{
		"authority":          "api.maersk.com",
		"accept":             "application/json",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"origin":             "https://www.maersk.com.cn",
		"referer":            "https://www.maersk.com.cn/",
		"sec-ch-ua":          "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"user-agent":         uarand.GetRandom(),
	}
	url := fmt.Sprintf(`https://api.maersk.com/track/%s?operator=MAEU`, number)
	response, err := m.r.Url(url).Method("GET").Headers(headers).Do(context.Background())
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *maeuResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	if len(s.Containers) != 0 {
		return s, nil
	}
	return nil, errors.New("no len")
}

type maeuArrivedChecker struct {
	r IMaeuRequest
}

func NewMaeuArrivedChecker(maeuRequest IMaeuRequest) *maeuArrivedChecker {
	return &maeuArrivedChecker{r: maeuRequest}
}
func (m *maeuArrivedChecker) checkStatus(response *maeuResponse) IsArrived {
	fmt.Println(response.Containers[0].Status)
	return IsArrived(strings.EqualFold(response.Containers[0].Status, "COMPLETE"))
}
func (m *maeuArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	resp, err := m.r.Get(result.Container)
	fmt.Println(err)
	if err != nil {
		return true
	}
	return m.checkStatus(resp)
}
func (m *maeuArrivedChecker) checkBillNoArrived(_ BillNumberResponse) IsArrived {
	return true
}

type sitcArrivedChecker struct{}

func newSitcArrivedChecker() *sitcArrivedChecker {
	return &sitcArrivedChecker{}
}
func (s *sitcArrivedChecker) checkInfoAboutMoving(result []BaseInfoAboutMoving) IsArrived {
	for _, item := range result {
		if strings.Contains(strings.ToUpper(item.OperationName), strings.ToUpper("Inbound Delivery")) ||
			strings.Contains(strings.ToUpper(item.OperationName), strings.ToUpper("Empty Container")) {
			return true
		}
	}
	return false
}
func (s *sitcArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	return s.checkInfoAboutMoving(result.InfoAboutMoving)
}
func (s *sitcArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	return s.checkInfoAboutMoving(result.InfoAboutMoving)

}

type zhguArrivedChecker struct{}

func newZhguArrivedChecker() *zhguArrivedChecker {
	return &zhguArrivedChecker{}
}

func (z *zhguArrivedChecker) checkContainerArrived(result ContainerNumberResponse) IsArrived {
	for _, v := range result.InfoAboutMoving {
		if strings.EqualFold(v.OperationName, "ATA") {
			return true
		}
	}
	return false
}
func (z *zhguArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	for _, v := range result.InfoAboutMoving {
		if strings.EqualFold(v.OperationName, "ATA") {
			return true
		}
	}
	return false

}

type reelArrivedChecker struct {
}

func newReelArrivedChecker() *reelArrivedChecker {
	return &reelArrivedChecker{}

}
func (r *reelArrivedChecker) checkBillNoArrived(result BillNumberResponse) IsArrived {
	for _, v := range result.InfoAboutMoving {
		if strings.EqualFold(strings.ToUpper(v.OperationName), "PICKUP BY CONSIGNEE") {
			return true
		}
	}
	return false

}

type ArrivedChecker struct {
	*skluArrivedChecker
	*fesoArrivedChecker
	*mscuArrivedChecker
	*oneyArrivedChecker
	*cosuArrivedChecker
	*maeuArrivedChecker
	*sitcArrivedChecker
	*zhguArrivedChecker
	*reelArrivedChecker
}

func NewArrivedChecker() *ArrivedChecker {
	return &ArrivedChecker{skluArrivedChecker: newSkluArrivedChecker(), fesoArrivedChecker: newFesoArrivedChecker(),
		mscuArrivedChecker: newMscuArrivedChecker(),
		oneyArrivedChecker: newOneyArrivedChecker(),
		cosuArrivedChecker: newCosuArrivedChecker(),
		maeuArrivedChecker: NewMaeuArrivedChecker(newMaeuRequest()),
		sitcArrivedChecker: newSitcArrivedChecker(),
		zhguArrivedChecker: newZhguArrivedChecker(),
		reelArrivedChecker: newReelArrivedChecker(),
	}
}

func (a *ArrivedChecker) CheckContainerArrived(result ContainerNumberResponse) IsArrived {
	switch strings.ToUpper(result.Scac) {
	case "HALU":
		return a.skluArrivedChecker.checkContainerArrived(result)
	case "FESO":
		return a.fesoArrivedChecker.checkContainerArrived(result)
	case "SKLU":
		return a.skluArrivedChecker.checkContainerArrived(result)
	case "MSCU":
		return a.mscuArrivedChecker.checkContainerArrived(result)
	case "ONEY":
		return a.oneyArrivedChecker.checkContainerArrived(result)
	case "COSU":
		return a.cosuArrivedChecker.checkContainerArrived(result)
	case "MAEU":
		fmt.Println("MAEU")
		return a.maeuArrivedChecker.checkContainerArrived(result)
	case "SITC":
		return a.sitcArrivedChecker.checkContainerArrived(result)
	case "ZHGU":
		return a.zhguArrivedChecker.checkContainerArrived(result)
	default:
		return false
	}
}

// CheckBillNoArrived
func (a *ArrivedChecker) CheckBillNoArrived(result BillNumberResponse) IsArrived {
	switch strings.ToUpper(result.Scac) {
	case "HALU":
		return a.skluArrivedChecker.checkBillNoArrived(result)
	case "FESO":
		return a.fesoArrivedChecker.checkBillNoArrived(result)
	case "SKLU":
		return a.skluArrivedChecker.checkBillNoArrived(result)
	case "MSCU":
		return a.mscuArrivedChecker.checkBillNoArrived(result)
	case "ONEY":
		return a.oneyArrivedChecker.checkBillNoArrived(result)
	case "COSU":
		return a.cosuArrivedChecker.checkBillNoArrived(result)
	case "MAEU":
		return a.maeuArrivedChecker.checkBillNoArrived(result)
	case "SITC":
		return a.sitcArrivedChecker.checkBillNoArrived(result)
	case "ZHGU":
		return a.zhguArrivedChecker.checkBillNoArrived(result)
	default:
		return false
	}
}
