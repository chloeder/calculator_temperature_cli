package model

type Kelvin struct {
	Suhu float64
}

func (k Kelvin) ToKelvin() float64 {
	return k.Suhu
}

func (k Kelvin) ToCelcius() float64 {
	return k.Suhu - 273.15

}

func (k Kelvin) ToFahrenheit() float64 {
	return (k.Suhu * 9/5) - 459.67
}

func (k Kelvin) ToReamur() float64 {
	return (k.Suhu - 273.15) * 4/5
}
