package zhgu

type GetEtaException struct {
}

func NewGetEtaException() *GetEtaException {
	return &GetEtaException{}
}

func (g *GetEtaException) Error() string {
	return "Cannot get ETA!"
}

type GetEtdException struct {
}

func NewGetEtdException() *GetEtdException {
	return &GetEtdException{}
}
func (g *GetEtdException) Error() string {
	return "Cannot find ETD!"
}
