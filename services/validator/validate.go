package validator

import (
	"fmt"
	"strconv"
	"strings"
)

type Validator float64

func (v *Validator) Time() error {
	time := v.getString()
}

func (v *Validator) getMapTime() map[string]string {
	t := v.getString()
	return map[string]string{
		"year": t[0:2],
	}
}

func (v *Validator) Delta() error {
	str := strings.Split(v.getString(), ".")
	if len(str) != 1 {
		return fmt.Errorf("wrong float format must have one point now %d", len(str))
	}

	if err := v.checkDelta(str[0], "day"); err != nil {
		return err
	}
	return v.checkDelta(str[1], "hour")

}

func (Validator) checkDelta(str, timeType string) error {
	res, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	if res > 0 && res <= 99 {
		return nil
	}
	return fmt.Errorf("wrong delta [%s] must be from 1 to 99 now: %d", timeType, res)
}

func (v *Validator) getString() string {
	return strconv.FormatFloat(float64(*v), 'f', 6, 64)
}
