package timeservice

import (
	"fmt"
	"testing"
)

func TestNewTimeService(t *testing.T) {
	ts := NewTimeService(nil)
	fmt.Println(ts.GetFloat64())
	fmt.Println(ts.GetString())
}
