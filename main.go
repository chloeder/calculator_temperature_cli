package main

func main()  {

}

func ConvertCelciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

func ConvertCelciusToKelvin(celcius float64) float64 {
	return celcius + 273.15
}

func ConvertCelciusToReamur(celcius float64) float64 {
	return celcius * 4/5
}

func ConvertFahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9
}
func ConvertFahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit + 459.67) * 5/9
}
func ConvertFahrenheitToReamur(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 4/9
}
func ConvertKelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}
func ConvertKelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin * 9/5) - 459.67
}
func ConvertKelvinToReamur(kelvin float64) float64 {
	return (kelvin - 273.15) * 4/5
}
func ConvertReamurToCelcius(reamur float64) float64 {
	return reamur * 5/4
}
func ConvertReamurToFahrenheit(reamur float64) float64 {
	return (reamur * 9/4) + 32
}
func ConvertReamurToKelvin(reamur float64) float64 {
	return (reamur * 5/4) + 273.15
}
