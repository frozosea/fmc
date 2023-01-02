package oney

type InfoAboutMovingListEntity struct {
	MaxRows     int             `json:"maxRows"`
	Models      []interface{}   `json:"models"`
	VslCd       string          `json:"vslCd"`
	No          string          `json:"no"`
	CopNo       string          `json:"copNo"`
	EventDt     string          `json:"eventDt"`
	VslEngNm    string          `json:"vslEngNm"`
	PlaceNm     string          `json:"placeNm"`
	SkdVoyNo    string          `json:"skdVoyNo"`
	YardNm      string          `json:"yardNm"`
	CopDtlSeq   string          `json:"copDtlSeq"`
	SkdDirCd    string          `json:"skdDirCd"`
	ActTpCd     string          `json:"actTpCd"`
	StatusNm    string          `json:"statusNm"`
	StatusCd    string          `json:"statusCd"`
	NodCd       string          `json:"nodCd"`
	Vvd         string          `json:"vvd"`
	LloydNo     string          `json:"lloydNo"`
	HashColumns [][]interface{} `json:"hashColumns"`
	HashFields  []interface{}   `json:"hashFields"`
}

type BaseApiResponseEntity struct {
	TRANSRESULTKEY string `json:"TRANS_RESULT_KEY"`
	Exception      string `json:"Exception"`
	Count          string `json:"count"`
}

type InfoAboutMovingApiResponseSchema struct {
	*BaseApiResponseEntity
	List []*InfoAboutMovingListEntity `json:"list"`
}

type ContainerSizeListEntity struct {
	MaxRows     int             `json:"maxRows"`
	Models      []interface{}   `json:"models"`
	Weight      string          `json:"weight"`
	CopNo       string          `json:"copNo"`
	BlNo        string          `json:"blNo"`
	EventDt     string          `json:"eventDt"`
	CntrTpszCd  string          `json:"cntrTpszCd"`
	CopStsCd    string          `json:"copStsCd"`
	Piece       string          `json:"piece"`
	SealNo      string          `json:"sealNo"`
	PlaceNm     string          `json:"placeNm"`
	YardNm      string          `json:"yardNm"`
	StatusNm    string          `json:"statusNm"`
	BkgNo       string          `json:"bkgNo"`
	PoNo        string          `json:"poNo"`
	YardCd      string          `json:"yardCd"`
	StatusCd    string          `json:"statusCd"`
	CntrNo      string          `json:"cntrNo"`
	CntrTpszNm  string          `json:"cntrTpszNm"`
	DspbkgNo    string          `json:"dspbkgNo"`
	SocFlg      string          `json:"socFlg"`
	MvmtStsCd   string          `json:"mvmtStsCd"`
	HashColumns [][]interface{} `json:"hashColumns"`
	HashFields  []interface{}   `json:"hashFields"`
}

type ContainerSizeApiResponseSchema struct {
	*BaseApiResponseEntity
	List []*ContainerSizeListEntity `json:"list"`
}

type BkgAndCopNosApiResponseSchema struct {
	*BaseApiResponseEntity
	List []*struct {
		MaxRows     int             `json:"maxRows"`
		Models      []interface{}   `json:"models"`
		CopNo       string          `json:"copNo"`
		BkgNo       string          `json:"bkgNo"`
		CntrNo      string          `json:"cntrNo"`
		EnblFlag    string          `json:"enblFlag"`
		HashColumns [][]interface{} `json:"hashColumns"`
		HashFields  []interface{}   `json:"hashFields"`
	}
}
