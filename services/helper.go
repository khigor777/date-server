package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Delta struct {
	Day, Hour, Min, Second int
}

func ParseDelta(delta float64) *Delta {
	d := strings.Replace(fmt.Sprintf("%09.6f", delta), ".", "", -1)
	day, _ := strToInt(d[0:2])
	hour, _ := strToInt(d[2:4])
	min, _ := strToInt(d[4:6])
	sec, _ := strToInt(d[6:8])
	return &Delta{
		Day:    day,
		Hour:   hour,
		Min:    min,
		Second: sec,
	}
}

func ParseTime(timeF float64) (*time.Time, error) {
	t := strconv.FormatFloat(timeF, 'f', 6, 64)
	currentTime, err := time.Parse(DefaultFormat, t)
	if err != nil {
		return nil, err
	}
	return &currentTime, nil
}

func strToInt(in string) (int, error) {
	return strconv.Atoi(in)
}
