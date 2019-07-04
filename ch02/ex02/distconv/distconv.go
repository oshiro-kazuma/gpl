package distconv

import "fmt"

type Meter float64
type Feet float64

const (
	coeffience = 3.2808
)

func MToF(m Meter) Feet {
	return Feet(m * coeffience)
}

func FToM(f Feet) Meter {
	return Meter(f / coeffience)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gF", f)
}
