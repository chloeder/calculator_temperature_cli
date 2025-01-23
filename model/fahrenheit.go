package model


type Fahrenheit struct {
	Suhu float64
}

func (f Fahrenheit) ToFahrenheit() float64 {
	return f.Suhu
}

func (f Fahrenheit) ToCelcius() float64 {
	return (f.Suhu - 32) * 5/9

}

func (f Fahrenheit) ToKelvin() float64 {
	return (f.Suhu + 459.67) * 5/9
}

func (f Fahrenheit) ToReamur() float64 {
	return (f.Suhu - 32) * 4/9
}
