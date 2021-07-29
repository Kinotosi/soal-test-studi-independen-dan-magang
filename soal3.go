package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var result string
	var pukul = "07:30:00 PM"

	// Mencari dan memisahkan number dan huruf. Dengan cara mencari angka 0-9 dan huruf P, A, dan M menggunakan regex
	var regex, err = regexp.Compile(`[0-9PAM]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(pukul, -1)  // Maka akan menghasilkan [jam menit detik AM/PM]
	var res1Convert, _ = strconv.Atoi(res1[0]) // Merubah string menjadi int

	// Jika PM, maka
	if res1[3] == "PM" {
		// Dan jika jam 12, maka
		if res1[0] == "12" {
			result = res1[0] + ":" + res1[1]
			fmt.Println(result) // Akan menampilkan 12:00
			// Jika, tidak
		} else {
			// maka setiap jam akan ditambah dengan nilai 12
			var jam = strconv.Itoa(res1Convert + 12)
			result = jam + ":" + res1[1]
			fmt.Println(result)
		}
		// Jika, tidak PM. Maka
	} else {
		// Dan jika jam 12,
		if res1[0] == "12" {
			// Akan dirubah menjadi 00:00
			result = "00:00"
			fmt.Println(result)
			// Jika tidak
		} else {
			// Maka akan menampilkan jam dan menit
			result = res1[0] + ":" + res1[1]
			fmt.Println(result)
		}
	}
}
