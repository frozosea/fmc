package cosu

type GetEtaError struct {
}

func NewGetEtaError() *GetEtaError {
	return &GetEtaError{}
}

func (g *GetEtaError) Error() string {
	return "Cannot get ETA, because COSU server response is empty!"
}
