package main

import "fmt"

func main()  {
	fmt.Println("Pilihan kalkulator suhu yang ingin dipakai : ")
	fmt.Println("1. Konversi Celcius ke Fahrenheit ")
	fmt.Println("2. Konversi Celcius ke Kelvin ")
	fmt.Println("3. Konversi Celcius ke Reamur ")
	fmt.Println("4. Konversi Fahrenheit ke Celcius ")
	fmt.Println("5. Konversi Fahrenheit ke Kelvin ")
	fmt.Println("6. Konversi Fahrenheit ke Reamur ")
	fmt.Println("7. Konversi Kelvin ke Celcius ")
	fmt.Println("8. Konversi Kelvin ke Fahrenheit ")
	fmt.Println("9. Konversi Kelvin ke Reamur ")
	fmt.Println("10. Konversi Reamur ke Celcius ")
	fmt.Println("11. Konversi Reamur ke Fahrenheit ")
	fmt.Println("12. Konversi Reamur ke Kelvin ")

	var pilihan int
	fmt.Println("Masukkan pilihan anda: ")
	fmt.Scanf("%d", &pilihan)
	for pilihan < 1 || pilihan > 12 {
		fmt.Println("Pilihan tidak tersedia dalam kalkulator")
		fmt.Scanf("%d", &pilihan)
	}

	var suhu float64
	fmt.Println("Masukkan suhu yang diinginkan: ")
	fmt.Scanf("%f", suhu)

	var suhu_akhir float64
	switch pilihan {
		case 1:
			suhu_akhir = ConvertCelciusToFahrenheit(suhu)
		case 2:
			suhu_akhir = ConvertCelciusToKelvin(suhu)
		case 3:
			suhu_akhir = ConvertCelciusToReamur(suhu)
		case 4:
			suhu_akhir = ConvertFahrenheitToCelcius(suhu)
		case 5:
			suhu_akhir = ConvertFahrenheitToKelvin(suhu)
		case 6:
			suhu_akhir = ConvertFahrenheitToReamur(suhu)
		case 7:
			suhu_akhir = ConvertKelvinToCelcius(suhu)
		case 8:
			suhu_akhir = ConvertKelvinToFahrenheit(suhu)
		case 9:
			suhu_akhir = ConvertKelvinToReamur(suhu)
		case 10:
			suhu_akhir = ConvertReamurToCelcius(suhu)
		case 11:
			suhu_akhir = ConvertReamurToFahrenheit(suhu)
		case 12:
			suhu_akhir = ConvertReamurToKelvin(suhu)
		default:
			 fmt.Println("KOSONG")
	}

	fmt.Printf("Suhu akhir dari hasil konversi adalah : %.2f\n", suhu_akhir)
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
