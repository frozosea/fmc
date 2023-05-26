package sklu

type LensNotEqualError struct {
}

func (l *LensNotEqualError) Error() string {
	return "lens of arrays (times,operations,locations,vessels) are not equal"
}
