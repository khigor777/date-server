package validator

import "testing"

func TestValidator(t *testing.T) {
	v := Validator(181029.080046)
	v.Time()
}
