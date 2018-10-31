package validator

import (
	"testing"
	"time"

	"github.com/khigor777/date-server/services"
)

func TestValidatorValidateTime(t *testing.T) {

	testData := []struct {
		in    float64
		error bool
	}{
		{in: 181029.230059, error: false},
		{in: 181029.250059, error: true},
		{in: 181029.246100, error: true},
		{in: 181029.240061, error: true},
		{in: 181032.240059, error: true},
		{in: 1311032.240059, error: true},
		{in: 181430.240059, error: true},
	}
	for _, val := range testData {
		v := Validator(val.in)
		err := v.ValidateTime(services.DefaultFormat)

		if _, ok := err.(*time.ParseError); ok != val.error {
			t.Error(err)
		}

	}

}

func TestValidatorValidateDelta(t *testing.T) {
	testData := []struct {
		in    float64
		error bool
	}{
		{99, false},
		{99.99, false},
		{99.0199, false},
		{99.0199, false},
		{99.019901, false},
		{9, true},
		{9., true},
		{9.1, true},
		{9.00009999, true},
		{0, true},
	}

	for _, val := range testData {
		v := Validator(val.in)
		res := v.ValidateDelta()
		if val.error != (res == ValidateError) {
			t.Errorf("%v, in:%v", res, val.in)
		}
	}
}
