package timeservice

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNewTimeService(t *testing.T) {
	f, _ := new(big.Float).SetPrec(500).SetString("181029.0449270")
	fmt.Println(f.Float64())
	ts := NewTimeService(nil)
	fmt.Println(ts.GetFloat64())
	fmt.Println(ts.GetString())
}
