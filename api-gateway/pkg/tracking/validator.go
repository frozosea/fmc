package tracking

import (
	"errors"
	"regexp"
)

type Validator struct{}

func (v *Validator) ValidateContainer(s string) error {
	if len(s) > 12 {
		return errors.New("cannot validate container")
	}
	ok, err := regexp.MatchString("[a-zA-Z]{3,4}[0-9]{5,7}", s)
	if !ok {
		return errors.New("cannot validate container")
	}
	if err != nil {
		return err
	}
	return nil
}
func (v *Validator) ValidateBillNumber(bill string) error {
	if len(bill) > 30 || len(bill) < 9 {
		return errors.New("cannot validate bill number")
	}
	ok, err := regexp.MatchString(`[a-zA-Z]{3,8}[0-9]{5,22}`, bill)
	if !ok || err != nil {
		return errors.New("cannot validate container")
	}
	return nil
}
func (v *Validator) ValidateScac(s string) error {
	if len(s) > 4 {
		return errors.New("cannot validate scac")
	}
	ok, err := regexp.MatchString("[a-zA-Z]{4}", s)
	if !ok || err != nil {
		return errors.New("cannot validate scac")
	}
	return nil
}
