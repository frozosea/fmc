package sklu

type LensNotEqualError struct {
}

func (l *LensNotEqualError) Error() string {
	return "lens not equal"
}
