package main

import (
	"SD_2024_Project/Controller"
	"SD_2024_Project/Database"
	"SD_2024_Project/Handler"
	"SD_2024_Project/Model"
	"SD_2024_Project/Node"
	"fmt"
	"net/http"
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

func static_insert() {
	//insert Member
	Controller.MemberInsert("Kurniawan1", "aan1", "0888888")
	Controller.MemberInsert("Kurniawan2", "", "0888888")
	Controller.MemberInsert("Kurniawan3", "aan1", "")
	Controller.MemberInsert("kurniawan4", "aan2", "0000")
	Controller.MemberInsert("kurniawan5", "aan3", "11111")

	//insert Barang

}

func webProgram() {
	http.HandleFunc("/", Handler.ViewHandler)
	http.HandleFunc("/insert", Handler.InsertHandler)
	http.HandleFunc("/delete", Handler.DeleteMemberHandler)
	http.HandleFunc("/update", Handler.UpdateMemberHandler)

	// Menjalankan web server
	http.ListenAndServe(":8080", nil)
}

func modelBarang() {
	//test model barang
	Model.BarangInsert("Kopi", 10000, 10)
	Model.BarangInsert("Teh", 5000, 20)
	Model.BarangInsert("Susu", 15000, 5)
	Model.BarangInsert("Gula", 5000, 10)
	fmt.Println(Model.BarangReadAll())
	fmt.Println(Model.BarangSearch(2))
	fmt.Println(Model.BarangSearch(4))
	fmt.Println(Model.BarangUpdate(1, "Kopi Hitam", 15000, 10))
	fmt.Println(Model.BarangUpdate(3, "Gula Merah", 10000, 10))
	fmt.Println(Model.BarangReadAll())
	fmt.Println(Model.BarangDelete(1))
	fmt.Println(Model.BarangDelete(3))
	fmt.Println(Model.BarangReadAll())
}

func ControllerBarangTest() {
	//test controller barang
	Controller.BarangInsert("Kopi", 10000, 10)
	Controller.BarangInsert("Teh", 5000, 20)
	Controller.BarangInsert("Susu", 15000, 5)
	Controller.BarangInsert("Gula", 5000, 10)
}

func testPenjualan() {
	//test detail Penjualan
	detail := []Node.DetailPenjualanNode{}
	detail = append(detail, Node.DetailPenjualanNode{IdBarang: 1, Jumlah: 2, Harga: 0})
	detail = append(detail, Node.DetailPenjualanNode{IdBarang: 2, Jumlah: 3, Harga: 0})
	Model.PenjualanInsert(1, detail)

	detail2 := []Node.DetailPenjualanNode{}
	detail2 = append(detail2, Node.DetailPenjualanNode{IdBarang: 1, Jumlah: 5, Harga: 0})
	detail2 = append(detail2, Node.DetailPenjualanNode{IdBarang: 2, Jumlah: 6, Harga: 0})
	Model.PenjualanInsert(1, detail2)

	fmt.Println(Database.DBPenjualan.Next.Next.Penjualan.Detail)
	fmt.Println(Database.DBPenjualan.Next.Next.Penjualan.Total)
	fmt.Println(Database.DBMember.Next.Member)
}
func main() {
	//menu_program()
	static_insert()

	ControllerBarangTest()
	testPenjualan()
	webProgram()
}
