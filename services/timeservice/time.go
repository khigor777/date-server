package timeservice

import (
	"strconv"
	"time"
)

const defaultFormat = "060102.030405" //YYMMDD.hhmmss
type DbService interface {
	SetDelta(time float64) error
	GetDelta() (float64, error)
}

type TimeService struct {
	time time.Time
	db   DbService
}

func NewTimeService(db DbService) *TimeService {
	return &TimeService{
		time: time.Now().UTC(),
		db:   db,
	}
}

func (ts *TimeService) GetFloat64() (float64, error) {
	f, err := strconv.ParseFloat(ts.getTime(), 64)

	if err != nil {
		return 0, err
	}
	return f, nil
}

func (ts *TimeService) GetString() string {
	return ts.getTime()
}

func (ts *TimeService) getTime() string {
	return ts.time.Format(defaultFormat)
}
