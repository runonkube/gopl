package weightconv

import "fmt"

const weightMultiplier = 2.2046

type Pound float64
type Kilo float64

func ToKilos(pounds Pound) Kilo {
	return Kilo(pounds) / weightMultiplier
}

func ToPounds(kilos Kilo) Pound {
	return Pound(kilos) * weightMultiplier
}

func (k Kilo) String() string {
	return fmt.Sprintf("%.4g Kilos", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.4g Pounds", p)
}
