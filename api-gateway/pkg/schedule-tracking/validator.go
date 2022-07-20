package schedule_tracking

import (
	"errors"
	"regexp"
)

type Validator struct {
}

func (v *Validator) ValidateTime(s string) error {
	ok, err := regexp.MatchString("\\d{1,2}:\\d{2}", s)
	if !ok || err != nil {
		return errors.New("cannot validate time")
	}
	return nil
}
func (v *Validator) ValidateBill(oneBill string) error {
	if len(oneBill) > 30 {
		return errors.New("cannot validate bill number")
	}
	ok, err := regexp.MatchString(`[a-zA-Z]{3,20}\w(.+)`, oneBill)
	if !ok || err != nil {
		return errors.New("cannot validate container")
	}
	return nil
}
func (v *Validator) ValidateBills(bills []string) error {
	for _, oneBill := range bills {
		if err := v.ValidateBill(oneBill); err != nil {
			return err
		}
	}
	return nil
}
func (v *Validator) ValidateContainers(containers []string) error {
	for _, b := range containers {
		if len(b) > 12 {
			return errors.New("cannot validate container")
		}
		ok, err := regexp.MatchString("[a-zA-Z]{3,4}[0-9]{5,7}", b)
		if !ok {
			return errors.New("cannot validate container")
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (v *Validator) ValidateEmails(emails []string) error {
	for _, email := range emails {
		ok, err := regexp.MatchString("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])", email)
		if err != nil || !ok {
			return errors.New("cannot validate email address")
		}
	}
	return nil
}
