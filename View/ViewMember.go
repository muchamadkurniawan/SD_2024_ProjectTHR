package View

import (
	"SD_2024_Project/Controller"
	"fmt"
)

func MemberInsert() {
	var namaDepan, namaBelakang, noTelp string
	fmt.Print("masukkan Nama Depan : ")
	fmt.Scan(&namaDepan)
	fmt.Print("masukkan Nama Belakang : ")
	fmt.Scan(&namaBelakang)
	fmt.Print("masukkan No TELP : ")
	fmt.Scan(&noTelp)
	cek := Controller.MemberInsert(namaDepan, namaBelakang, noTelp)
	if cek == true {
		fmt.Println("insert berhasil")
	} else {
		fmt.Println("insert gagal")
	}
}

func MemberReadAll() {
	members := Controller.MembersView()
	if members != nil {
		fmt.Println("--------------------------")
		for _, member := range members {
			fmt.Println("ID Member : ", member.Id)
			fmt.Println("Nama Depan : ", member.NamaDepan)
			fmt.Println("Nama Belakang : ", member.NamaBelakang)
			fmt.Println("No Telp ", member.NoTelp)
			fmt.Println("Point Member ", member.Point)
			fmt.Println("---------------------------")
		}
	} else {
		fmt.Println("Data Member Masih Kosong")
	}
}
