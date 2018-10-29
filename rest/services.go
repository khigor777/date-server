package web

import (
	"github.com/sirupsen/logrus"
)

type Services struct {
	Logger      *logrus.Logger
	TimeService Timer
}

type Timer interface {
	GetFloat64() (float64, error)
	GetString() string
}
