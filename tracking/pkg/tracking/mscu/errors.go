package mscu

type GetEtaException struct {
}

func NewGetEtaException() *GetEtaException {
	return &GetEtaException{}
}

func (g *GetEtaException) Error() string {
	return "Cannot get ETA!"
}
