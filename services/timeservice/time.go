package timeservice

import (
	"fmt"
	"strconv"
	"time"

	"github.com/khigor777/date-server/services"
)

type DbService interface {
	Set(time float64) error
	Get() (float64, error)
	Remove() error
}

type TimeService struct {
	db DbService
}

func NewTimeService(db DbService) *TimeService {
	return &TimeService{
		db: db,
	}
}

func (ts *TimeService) GetFloat() (float64, error) {
	return strconv.ParseFloat(ts.getTime(), 64)
}

func (ts *TimeService) GetString(timeF float64) (string, error) {
	t := ts.parseFloat(timeF)
	str, err := time.Parse(services.DefaultFormat, t)
	if err != nil {
		return "", err
	}
	return str.String(), nil
}

func (ts *TimeService) Set(timeF float64) error {
	t, err := services.ParseTime(timeF)
	if err != nil {
		return err
	}
	d := time.Now().Sub(*t)
	sec := d.Seconds()
	min := d.Minutes()
	days := int(d.Hours() / 24)
	hours := float64(days*24) - d.Hours()
	str := fmt.Sprintf("%02d.%f%f%f", days, hours, min, sec)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	return ts.db.Set(f)
}

func (ts *TimeService) Reset() error {
	return ts.db.Remove()
}

func (ts *TimeService) Add(timeF float64, delta float64) (float64, error) {
	d := services.ParseDelta(delta)
	t, err := services.ParseTime(timeF)
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

func (ts *TimeService) parseFloat(timeF float64) string {
	return strconv.FormatFloat(timeF, 'f', 6, 64)
}
