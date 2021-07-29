package main

import (
	"fmt"
	"strings"
)

// Membuat fungsi cekTeks dengan 2 parameter dan kembalian boolean.
func cekTeks(teksAwal, teksAkhir string) bool {
	var jumlahTeks = len(teksAkhir)

	// Jika jumlah teksAwal sama dengan 0 maka akan mengembalikan true. Sekaligus menghentikan rekrusi
	if len(teksAwal) == 0 {
		return true
	}

	// Jika huruf awal teksAwal tidak sesuai dengan huruf akhir teksAkhir,
	if string(teksAwal[0]) != string(teksAkhir[jumlahTeks-1]) {
		return false // maka akan mengembalikkan false. Sekaligus menghentikkan rekrusi
	} else {
		// membuat rekrusi dengan mengembalikkan fungsi. Dan menghapus huruf awal dari teksAwal dan huruf akhir dari teksAkhir
		return cekTeks(string(teksAwal[1:]), string(teksAkhir[:jumlahTeks-1]))
	}
}

func main() {
	// Memisahkan nilai dengan cara Split menggunakan strings.Split
	var teks = strings.Split("Katak di balik kataK", " ")
	// Memasukkan teks pertama kedalam variabel teksAwal. Dan teks akhir ke variable teksAkhir
	var teksAwal = teks[0]
	var teksAkhir = teks[3]

	// Memanggil fungsi cekTeks kemudian disimpan di variabel result
	var result = cekTeks(teksAwal, teksAkhir)
	fmt.Println(result)
}
