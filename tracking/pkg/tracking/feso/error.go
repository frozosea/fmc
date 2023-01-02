package feso

type GetEtaException struct {
}

func NewGetEtaException() *GetEtaException {
	return &GetEtaException{}
}

func (g *GetEtaException) Error() string {
	return "cannot get eta"
}
