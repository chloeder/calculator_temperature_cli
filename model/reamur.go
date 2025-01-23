package model

type Reamur struct {
	Suhu float64
}

func (k Reamur) ToReamur() float64 {
	return k.Suhu
}

func (r Reamur) ToCelcius() float64 {
	return r.Suhu * 5/4

}

func (r Reamur) ToFahrenheit() float64 {
	return (r.Suhu * 9/4) + 32
}

func (r Reamur) ToKelvin() float64 {
	return (r.Suhu * 5/4) + 273.15
}
