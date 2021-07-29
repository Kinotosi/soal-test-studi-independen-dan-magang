package main

import "fmt"

// Membuat fungsi reserve dan menggunkan metode rekrusi pada fungsi reserver
func reserve(texts string) string {
	// Jika jumlah huruf di texts sama dengan 0, maka
	if len(texts) == 0 {
		return texts // akan mengembalikkan nilai texts
	} else { // Jika tidak
		// merekrusi fungsi dengan cara mengembalikkan fungsi reserve dan menghilangkan huruf didepan
		return reserve(texts[1:]) + string(texts[0]) // Dan merubah nilai arguments
	}

}

func main() {
	// memanggil fungsi reverse dan disimpan di variabel text
	var text = reserve("Aplikasi")
	fmt.Println(text)
}
