package main

import (
	"fmt"
	"regexp"
)

func main() {
	var email = "argadwi90@.gmail.co.id"

	// Mencari text yang memiliki ciri2 '@.'teks'
	var emailSerch = regexp.MustCompile(`@.[a-z]+`)
	var emailResult = emailSerch.Split(email, -1) // akan memisahkan nama email dan domain, sehingga menghasilkan [namaEmail domain]

	// Apabila tidak ada @ yang diikuti tanda titi '.' maka hanya menghasilkan 1 nilai saja
	if len(emailResult) == 1 {
		fmt.Println("Email tidak valid.")
	} else {
		// Jika nama email kurang dari 20 karakter
		if len(emailResult[0]) <= 20 {
			// Dan memiliki domain .id atau .co.id, maka menampilkan email valid
			if emailResult[1] == ".id" || emailResult[1] == ".co.id" {
				fmt.Println("Email Valid")
				// Jika tidak maka akan menampilkan domain tidak valid
			} else {
				fmt.Println("Domain tidak sesuai yang ditentukan")
			}
			// Jika tidak maka akan menampilkan Email terlalu panjang
		} else {
			fmt.Println("Email terlalu panjang.")
		}
	}
}
