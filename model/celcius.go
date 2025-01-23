package model

type Celcius struct {
	Suhu float64
}

func (c Celcius) ToCelcius() float64 {
	return c.Suhu
}

func (c Celcius) ToFahrenheit() float64 {
	return (c.Suhu * 9.0 / 5.0) + 32
}

func (c Celcius) ToKelvin() float64 {
	return c.Suhu + 273.15
}

func (c Celcius) ToReamur() float64 {
	return c.Suhu * 4/5
}
