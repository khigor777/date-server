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

func (ts *TimeService) GetString() string {
	return ts.getTime()
}

func (ts *TimeService) getTime() string {
	return time.Now().UTC().Format(defaultFormat)
}
