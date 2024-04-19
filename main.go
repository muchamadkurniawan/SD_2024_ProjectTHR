package main

import (
	"SD_2024_Project/Controller"
	"SD_2024_Project/View"
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
	fmt.Println(Controller.MemberInsert("Kurniawan2", "", "0888888"))
	fmt.Println(Controller.MemberInsert("Kurniawan3", "aan1", ""))
	Controller.MemberInsert("kurniawan4", "aan2", "0000")
	Controller.MemberInsert("kurniawan5", "aan3", "11111")
	View.MemberSearch()
}
