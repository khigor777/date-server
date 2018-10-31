package validator

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var ValidateError = errors.New("dosen't fit the mask 00.000000")
var reDelta = regexp.MustCompile(`^[0-9]{2}([.][0-9]{2,6})?$`) //why global? speed optimization!

type timeValidator interface {
	ValidateTime() error
	ValidateDelta() error
}

type Validator float64

func (v *Validator) ValidateTime(format string) error {
	_, err := time.Parse(format, v.getString())
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) ValidateDelta() error {
	str := v.getString()
	str = strings.TrimRight(strings.TrimRight(str, "0"), ".")
	if reDelta.Match([]byte(str)) {
		return nil
	}
	return ValidateError
}

func (v *Validator) getString() string {
	return strconv.FormatFloat(float64(*v), 'f', 6, 64)
}
