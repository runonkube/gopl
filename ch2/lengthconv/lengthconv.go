package lengthconv

import "fmt"

type Feet float64
type Meter float64

const lengthMultiplier = 0.3048

func (feet Feet) String() string {
	return fmt.Sprintf("%.4g feet", feet)
}

func (meters Meter) String() string {
	return fmt.Sprintf("%.4g meters", meters)
}

func ToMeters(feet Feet) Meter {
	return Meter(feet) * lengthMultiplier
}

func ToFeet(meter Meter) Feet {
	return Feet(meter) / lengthMultiplier
}
