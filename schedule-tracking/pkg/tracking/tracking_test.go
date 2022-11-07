package tracking

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSkluArrivedChecker(t *testing.T) {
	testTable := []struct {
		infoAboutMoving []BaseInfoAboutMoving
		isArrived       bool
	}{
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Arrival (Scheduled)", Vessel: ""}}, isArrived: false},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Arrival", Vessel: ""}}, isArrived: true},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Arrival (15/15)", Vessel: ""}}, isArrived: true},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Pickup (1/1)", Vessel: ""}, {Time: time.Now(), Location: "", OperationName: "Return (1/1)", Vessel: ""}}, isArrived: false},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Pickup (1/1)", Vessel: ""}, {Time: time.Now(), Location: "", OperationName: "Departure(T/S)", Vessel: ""}}, isArrived: false},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), Location: "", OperationName: "Pickup (1/1)", Vessel: ""}, {Time: time.Now(), Location: "", OperationName: "Arrival(T/S)", Vessel: ""}}, isArrived: false},
	}
	a := newSkluArrivedChecker()
	for _, v := range testTable {
		res := a.checkInfoAboutMoving(v.infoAboutMoving)
		if v.isArrived {
			assert.Equal(t, IsArrived(true), res)
		} else {
			assert.Equal(t, IsArrived(false), res)
		}
	}
}
func TestFesoArrivedChecker(t *testing.T) {
	testTable := []struct {
		infoAboutMoving []BaseInfoAboutMoving
		isArrived       bool
	}{
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), OperationName: "Gate out empty for loading", Location: "MAGISTRAL", Vessel: ""}, {Time: time.Now(), OperationName: "Gate in empty from consignee", Location: "ZAPSIBCONT", Vessel: ""}}, isArrived: true},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), OperationName: "Gate out empty for loading", Location: "XIN BA DA", Vessel: ""}, {Time: time.Now(), OperationName: "Gate in empty from consignee", Location: "Honghai", Vessel: ""}, {Time: time.Now(), OperationName: "ETS", Location: "Shanghai", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "Load on board", Location: "Honghai", Vessel: "FESCO DALNEGORSK"}}, isArrived: true},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), OperationName: "Gate out empty for loading", Location: "XIN BA DA", Vessel: ""}, {Time: time.Now(), OperationName: "Gate in empty from consignee", Location: "Honghai", Vessel: ""}, {Time: time.Now(), OperationName: "ETS", Location: "Shanghai", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "Load on board", Location: "Honghai", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "ETA", Location: "Vostochny", Vessel: "FESCO DALNEGORSK"}}, isArrived: false},
		{infoAboutMoving: []BaseInfoAboutMoving{{Time: time.Now(), OperationName: "Gate out empty for loading", Location: "XIN BA DA", Vessel: ""}, {Time: time.Now(), OperationName: "Gate in empty from consignee", Location: "Honghai", Vessel: ""}, {Time: time.Now(), OperationName: "ETS", Location: "Shanghai", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "Load on board", Location: "Honghai", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "ETA", Location: "Vostochny", Vessel: "FESCO DALNEGORSK"}, {Time: time.Now(), OperationName: "Discharge", Location: "jhkghjh", Vessel: "bvvhg"}}, isArrived: true},
	}
	a := newFesoArrivedChecker()
	for _, v := range testTable {
		res := a.checkArrivedByInfoAboutMoving(v.infoAboutMoving)
		if v.isArrived {
			assert.Equal(t, IsArrived(true), res)
		} else {
			assert.Equal(t, IsArrived(false), res)
		}
	}
}
func TestMscuArrivedChecker(t *testing.T) {
	testTable := []struct {
		infoAboutMoving []BaseInfoAboutMoving
		isArrived       bool
	}{
		{infoAboutMoving: []BaseInfoAboutMoving{{
			Time:          time.Now(),
			OperationName: "Empty to Shipper",
			Location:      "CHONGQING, CN",
			Vessel:        "",
		},
			{
				Time:          time.Now(),
				OperationName: "Export at barge yard",
				Location:      "CHONGQING, CN",
				Vessel:        "",
			}}, isArrived: true},

		{infoAboutMoving: []BaseInfoAboutMoving{
			{
				Time:          time.Now(),
				OperationName: "Export at barge yard",
				Location:      "CHONGQING, CN",
				Vessel:        "",
			}}, isArrived: false},
	}
	a := newMscuArrivedChecker()
	for _, v := range testTable {
		res := a.checkInfoAboutMoving(v.infoAboutMoving)
		if v.isArrived {
			assert.Equal(t, IsArrived(true), res)
		} else {
			assert.Equal(t, IsArrived(false), res)
		}
	}
}
func TestOneyArrivedChecker(t *testing.T) {
	testTable := []struct {
		InfoAboutMoving []BaseInfoAboutMoving
		isArrived       bool
	}{
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Location:      "PUSAN, KOREA REPUBLIC OF",
			OperationName: "'HYUNDAI SINGAPORE 126E' Arrival at Port of Discharging",
			Time:          time.Now(),
			Vessel:        "HYUNDAI SINGAPORE",
		}}, isArrived: true},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Location:      "PUSAN, KOREA REPUBLIC OF",
			OperationName: "Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading",
			Time:          time.Now(),
			Vessel:        "HYUNDAI SINGAPORE",
		}}, isArrived: false},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Location:      "PUSAN, KOREA REPUBLIC OF",
			OperationName: "Gate In to Outbound Terminal",
			Time:          time.Now(),
			Vessel:        "HYUNDAI SINGAPORE",
		}}, isArrived: false},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Location:      "PUSAN, KOREA REPUBLIC OF",
			OperationName: "Empty Container Release to Shipper",
			Time:          time.Now(),
			Vessel:        "HYUNDAI SINGAPORE",
		}}, isArrived: true},
	}
	a := newOneyArrivedChecker()
	for _, v := range testTable {
		res := a.checkInfoAboutMoving(v.InfoAboutMoving)
		if v.isArrived {
			assert.Equal(t, IsArrived(true), res)
		} else {
			assert.Equal(t, IsArrived(false), res)
		}
	}
}

type maeuRequestMoch struct {
}

func (m *maeuRequestMoch) Get(_ string) (*maeuResponse, error) {
	var s maeuResponse
	if err := json.Unmarshal([]byte(`{
    "isContainerSearch": true,
    "origin": {
        "terminal": "Laem Chabang Terminal PORT D1",
        "geo_site": "9QYVWASPIUQBZ",
        "city": "Laem Chabang",
        "state": "",
        "country": "Thailand",
        "country_code": "TH",
        "geoid_city": "0NTCCB4ENSFX8",
        "site_type": "TERMINAL"
    },
    "destination": {
        "terminal": "",
        "geo_site": "Unknown",
        "city": "Spartanburg",
        "state": "South Carolina",
        "country": "United States",
        "country_code": "US",
        "geoid_city": "09L62FUTQWHHB",
        "site_type": ""
    },
    "containers": [
        {
            "container_num": "MSKU6874333",
            "container_size": "40",
            "container_type": "Dry",
            "iso_code": "42G0",
            "operator": "MAEU",
            "locations": [
                {
                    "terminal": "Win Win Container Depot",
                    "geo_site": "1GP8G188VBR2G",
                    "city": "Laem Chabang",
                    "state": "",
                    "country": "Thailand",
                    "country_code": "TH",
                    "geoid_city": "0NTCCB4ENSFX8",
                    "site_type": "DEPOT",
                    "events": [
                        {
                            "activity": "GATE-OUT-EMPTY",
                            "stempty": true,
                            "actfor": "EXP",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "actual_time": "2022-03-29T16:42:00.000",
                            "rkem_move": "GATE-OUT",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "Laem Chabang Terminal PORT D1",
                    "geo_site": "9QYVWASPIUQBZ",
                    "city": "Laem Chabang",
                    "state": "",
                    "country": "Thailand",
                    "country_code": "TH",
                    "geoid_city": "0NTCCB4ENSFX8",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "GATE-IN",
                            "stempty": false,
                            "actfor": "EXP",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-16T14:00:00.000",
                            "actual_time": "2022-03-30T10:02:00.000",
                            "rkem_move": "GATE-IN",
                            "is_cancelled": false,
                            "is_current": false
                        },
                        {
                            "activity": "LOAD",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-16T14:00:00.000",
                            "actual_time": "2022-04-14T00:15:00.000",
                            "rkem_move": "LOAD",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "YANGSHAN SGH GUANDONG TERMINAL",
                    "geo_site": "37O5HQ17XCL3X",
                    "city": "Shanghai",
                    "state": "Shanghai",
                    "country": "China",
                    "country_code": "CN",
                    "geoid_city": "2IW9P6J7XAW72",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "DISCHARG",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-25T08:00:00.000",
                            "actual_time": "2022-04-25T12:49:00.000",
                            "rkem_move": "DISCHARG",
                            "is_cancelled": false,
                            "is_current": false
                        },
                        {
                            "activity": "LOAD",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "ZIM WILMINGTON",
                            "voyage_num": "006E",
                            "vessel_num": "U5T",
                            "expected_time": "2022-05-02T18:30:00.000",
                            "actual_time": "2022-05-02T01:11:00.000",
                            "rkem_move": "LOAD",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "Charleston Wando Welch terminal N59",
                    "geo_site": "1ML38I7Q8BBKU",
                    "city": "Charleston",
                    "state": "South Carolina",
                    "country": "United States",
                    "country_code": "US",
                    "geoid_city": "3RSB4DDP23AM7",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "DISCHARG",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "ZIM WILMINGTON",
                            "voyage_num": "006E",
                            "vessel_num": "U5T",
                            "expected_time": "2022-06-11T07:00:00.000",
                            "actual_time": "2022-06-11T13:54:00.000",
                            "rkem_move": "DISCHARG",
                            "is_cancelled": false,
                            "is_current": false
                        },
                        {
                            "activity": "GATE-OUT",
                            "stempty": false,
                            "actfor": "DEL",
                            "vessel_name": "ZIM WILMINGTON",
                            "voyage_num": "006E",
                            "vessel_num": "U5T",
                            "expected_time": "2022-06-14T10:00:00.000",
                            "actual_time": "2022-06-30T12:46:00.000",
                            "rkem_move": "GATE-OUT",
                            "is_cancelled": false,
                            "is_current": true
                        }
                    ]
                }
            ],
            "eta_final_delivery": "2022-06-11T13:54:00.000",
            "latest": {
                "actual_time": "2022-06-30T12:46:00.000",
                "activity": "GATE-OUT",
                "stempty": false,
                "actfor": "DEL",
                "geo_site": "1ML38I7Q8BBKU",
                "city": "Charleston",
                "state": "South Carolina",
                "country": "United States",
                "country_code": "US"
            },
            "status": "IN-PROGRESS"
        }
    ]
}`), &s); err != nil {
		return &maeuResponse{}, err
	}
	return &s, nil
}
func TestMaeuArrivedChecker(t *testing.T) {
	moch := &maeuRequestMoch{}
	a := NewMaeuArrivedChecker(moch)
	r, err := moch.Get("")
	assert.NoError(t, err)
	assert.Equal(t, IsArrived(false), a.checkStatus(r))
}
func TestCosuArrivedChecker(t *testing.T) {
	testTable := []struct {
		InfoAboutMoving []BaseInfoAboutMoving
		isArrived       bool
	}{
		{InfoAboutMoving: []BaseInfoAboutMoving{
			{
				Time:          time.Now(),
				OperationName: "Discharged at Last POD",
				Location:      "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
				Vessel:        "Vessel",
			}},
			isArrived: true},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Time:          time.Now(),
			OperationName: "Empty Equipment Returned",
			Location:      "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
			Vessel:        "Vessel",
		}},
			isArrived: true},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Time:          time.Now(),
			OperationName: "Loaded at First POL",
			Location:      "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
			Vessel:        "Vessel",
		}},
			isArrived: false},
		{InfoAboutMoving: []BaseInfoAboutMoving{{
			Time:          time.Now(),
			OperationName: "Gate-In at First POL",
			Location:      "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
			Vessel:        "Vessel",
		}},
			isArrived: false},
	}
	a := newCosuArrivedChecker()
	for _, v := range testTable {
		res := a.checkInfoAboutMoving(v.InfoAboutMoving)
		if v.isArrived {
			assert.Equal(t, IsArrived(true), res)
		} else {
			assert.Equal(t, IsArrived(false), res)
		}
	}
}
