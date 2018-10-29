package response

import "fmt"

type FloatNumber float64

func (n FloatNumber) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.6f", n)), nil
}
