package sklu

type LensNotEqualError struct {
}

func (l *LensNotEqualError) Error() string {
	return "lens not equal"
}

type ParseTableWithEventsError struct {
}

func (p *ParseTableWithEventsError) Error() string {
	return "couldn't parse table in this html"
}
