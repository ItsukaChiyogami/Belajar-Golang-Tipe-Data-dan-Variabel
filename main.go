package main

import (
	"fmt"
)

func main() {

	// about integer and data floating
	// complex64 float32 int8 unint8

	fmt.Println("satu = ", 1)
	fmt.Println(`dua = `, 2)
	fmt.Println(`Tiga koma lima = `, 3.5)

	fmt.Println()
	// boolean
	fmt.Println("true = ", true)
	fmt.Println("false = ", false)

	// string
	fmt.Println("")
	fmt.Println("String")

	fmt.Println("")

	// function for String
	fmt.Println(len("wkwkwkwk")) //menghitung jumlah string
	fmt.Println("nani"[1])       // menghitung ASCII string

	// Variable
	var name string
	name = "Chiyogami"
	fmt.Println(name)

	name = "Tobiichi Origami"
	fmt.Println(name)

	Country := "Indonesia +62" //menambahkan : di := akan membuat variable gak bisa diulang contoh nama := dan nama := akan error
	fmt.Println(Country)

	Country = "japan +81"
	fmt.Println(Country)

	// Deklarasi Multiple Variable
	var (
		Nama = "Tobiichi Origammi"
		Umur = 18
		// wajib deklarasikan dan gunakan variablenya atau akan error
	)

	fmt.Println(Nama)
	fmt.Println(Umur)

}
