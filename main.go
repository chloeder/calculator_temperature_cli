package main

import (
	"calculator/model"
	"fmt"
)

type Calculated interface {
	ToCelcius() float64
	ToFahrenheit() float64
	ToKelvin() float64
	ToReamur() float64
}

func main()  {
	fmt.Println("Pilihan Suhu Awal : ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("4. Reamur")


	fmt.Println("Masukkan Suhu Awal : ")
	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 4 {
		fmt.Println("Pilihan suhu awal tidak tersedia dalam kalkulator")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("Pilihan Suhu Akhir : ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("4. Reamur")

	fmt.Println("Masukkan Suhu Akhir : ")
	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)
	for suhuAkhir < 1 || suhuAkhir > 4 {
		fmt.Println("Pilihan suhu akhir tidak tersedia dalam kalkulator")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Println("Masukkan suhu yang diinginkan: ")
	fmt.Scanf("%f", suhu)

	var calculate Calculated
	switch suhuAwal {
		case 1:
			calculate = model.Celcius{Suhu: suhu}
		case 2:
			calculate = model.Fahrenheit{Suhu: suhu}
		case 3:
			calculate = model.Kelvin{Suhu: suhu}
		case 4:
			calculate = model.Reamur{Suhu: suhu}
		default:
			fmt.Println("Pilihan tidak valid")
			return
	}

	var suhu_final float64
	switch suhuAkhir {
		case 1:
			suhu_final = calculate.ToCelcius()
		case 2:
			suhu_final = calculate.ToFahrenheit()
		case 3:
			suhu_final = calculate.ToKelvin()
		case 4:
			suhu_final = calculate.ToReamur()
	}





	fmt.Printf("Suhu akhir dari hasil konversi adalah : %.2f\n", suhu_final)
}
