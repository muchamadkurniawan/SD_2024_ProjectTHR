package main

import (
	"SD_2024_Project/Controller"
	"SD_2024_Project/Model"
	"fmt"
	"os"
)

func menu_program() {
	var pilih int
	for {
		fmt.Println("Menu Pilihan yang ada :")
		fmt.Println("1. Menu 1")
		fmt.Println("2. Exit")
		fmt.Println()
		fmt.Print("masukkan menu pilihan : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Println("anda masuk pilihan 1")
		case 2:
			fmt.Println("anda keluar program")
			os.Exit(0)
		}
	}
}

func main() {
	fmt.Println("Project Struktur Data 2024")

	fmt.Println(Controller.MemberInsert("Kurniawan1", "aan1", "0888888"))
	fmt.Println(Controller.MemberInsert("Kurniawan1", "", "0888888"))
	fmt.Println(Controller.MemberInsert("Kurniawan1", "aan1", ""))

	fmt.Println(Controller.MembersView())
	alamat := Model.MemberSearch(1)
	fmt.Println(alamat)
	alamat.Member.Point = 100
	fmt.Println(alamat)
	//View.MemberInsert()
	//View.MemberInsert()
	//View.MemberReadAll()
	//menu_program()
}
