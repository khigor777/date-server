package timeservice

import (
	"testing"
)

func TestNewTimeService(t *testing.T) {
	//181029.230059
	ts := NewTimeService(nil)
	ts.Add(181029.100202, 00.331112)
}
