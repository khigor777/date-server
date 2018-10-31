package timeservice

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/khigor777/date-server/services"
)

type DbService interface {
	SetDelta(time float64) error
	GetDelta() (float64, error)
}

type Delta struct {
	Day, Hour, Min, Second int
}

type TimeService struct {
	db DbService
}

func NewTimeService(db DbService) *TimeService {
	return &TimeService{
		db: db,
	}
}

func (ts *TimeService) GetFloat64() (float64, error) {
	return strconv.ParseFloat(ts.getTime(), 64)
}

func (ts *TimeService) GetString(timeF float64) (string, error) {
	t := strconv.FormatFloat(timeF, 'f', 6, 64)
	str, err := time.Parse(services.DefaultFormat, t)
	if err != nil {
		return "", err
	}
	return str.String(), nil
}

func (ts *TimeService) Add(timeF float64, delta float64) (float64, error) {
	d := ts.parseDelta(delta)
	t, err := ts.parseTime(timeF)
	if err != nil {
		return 0, err
	}
	dayInHour := d.Day * 24
	timeStr := t.Add(time.Hour*time.Duration(d.Hour+dayInHour) + time.Minute*time.Duration(d.Min) +
		time.Second*time.Duration(d.Second)).Format(services.DefaultFormat)
	return strconv.ParseFloat(timeStr, 64)
}

func (ts *TimeService) getTime() string {
	return time.Now().UTC().Format(services.DefaultFormat)
}

func (ts *TimeService) parseDelta(delta float64) *Delta {
	d := strings.Replace(fmt.Sprintf("%09.6f", delta), ".", "", -1)
	day, _ := ts.strToInt(d[0:2])
	hour, _ := ts.strToInt(d[2:4])
	min, _ := ts.strToInt(d[4:6])
	sec, _ := ts.strToInt(d[6:8])
	return &Delta{
		Day:    day,
		Hour:   hour,
		Min:    min,
		Second: sec,
	}
}

func (ts *TimeService) parseTime(timeF float64) (*time.Time, error) {
	t := strconv.FormatFloat(timeF, 'f', 6, 64)
	currentTime, err := time.Parse(services.DefaultFormat, t)
	if err != nil {
		return nil, err
	}
	return &currentTime, nil
}

func (ts *TimeService) strToInt(in string) (int, error) {
	return strconv.Atoi(in)
}
