package dnyg

type NumberInfoResponse struct {
	DltResultBlList []struct {
		OUTPOD   string `json:"OUTPOD"`
		OUTVDS   string `json:"OUTVDS"`
		OUTBKN   string `json:"OUTBKN"`
		OUTETD   string `json:"OUTETD"`
		OUTBNO   string `json:"OUTBNO"`
		OUTETA   string `json:"OUTETA"`
		OUTCNT   string `json:"OUTCNT"`
		OUTPOL   string `json:"OUTPOL"`
		SIZETYPE string `json:"SIZETYPE"`
	} `json:"dlt_resultBlList"`
}

type InfoAboutMovingResponse struct {
	DltResultMovementList []struct {
		OUTPOR    string `json:"OUTPOR"`
		OUTPVY    string `json:"OUTPVY"`
		OUTDTD    string `json:"OUTDTD"`
		CMVDPT    string `json:"CMVDPT"`
		OUTDES    string `json:"OUTDES"`
		OUTPODK   string `json:"OUTPODK"`
		OUTAREA   string `json:"OUTAREA"`
		CMVLCN    string `json:"CMVLCN"`
		OUTLOC    string `json:"OUTLOC"`
		OUTPOLK   string `json:"OUTPOLK"`
		RNUM      int    `json:"RNUM"`
		OUTPOD    string `json:"OUTPOD"`
		OUTCNO    string `json:"OUTCNO"`
		OUTSEQ    int    `json:"OUTSEQ"`
		OUTSTA    string `json:"OUTSTA"`
		OUTPOL    string `json:"OUTPOL"`
		OUTPORK   string `json:"OUTPORK"`
		OUTPVYK   string `json:"OUTPVYK"`
		OUTDIK    string `json:"OUTDIK"`
		MOUTSEQ   int    `json:"MOUTSEQ"`
		OUTDEK    string `json:"OUTDEK"`
		OUTVSSVOY string `json:"OUTVSSVOY"`
		OUTATD    string `json:"OUTATD,omitempty"`
		OUTVVD    string `json:"OUTVVD,omitempty"`
		OUTBNO    string `json:"OUTBNO,omitempty"`
		VSLNAME   string `json:"VSLNAME,omitempty"`
		LTDVVD    string `json:"LTDVVD,omitempty"`
		VSLCOD    string `json:"VSLCOD,omitempty"`
	} `json:"dlt_resultMovementList"`
}
