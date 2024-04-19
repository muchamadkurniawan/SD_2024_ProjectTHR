package View

import (
	"SD_2024_Project/Controller"
	"SD_2024_Project/Node"
	"fmt"
)

func viewAttribute(member Node.MemberNode) {
	fmt.Println("ID Member : ", member.Id)
	fmt.Println("Nama Depan : ", member.NamaDepan)
	fmt.Println("Nama Belakang : ", member.NamaBelakang)
	fmt.Println("No Telp ", member.NoTelp)
	fmt.Println("Point Member ", member.Point)
	fmt.Println("---------------------------")
}

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
			viewAttribute(member)
		}
	} else {
		fmt.Println("Data Member Masih Kosong")
	}
}

func MemberDelete() {
	var id int
	fmt.Print("Masukkan ID Member : ")
	fmt.Scan(&id)
	cek, member := Controller.MemberDelete(id)
	if cek == false {
		fmt.Println("ID yang anda masukkan tidak valid (ID harus lebih dari 0)")
	} else if member.Id == 0 {
		fmt.Println("ID yang anda cari tidak ada")
	} else {
		viewAttribute(member)
		fmt.Println("Data tersebut berhasil id HAPUS")
	}
}

func MemberUpdate() {
	var namaDepan, namaBelakang, noTelp string
	var id int
	fmt.Print("masukkan ID yang akan di UPDATE : ")
	fmt.Scan(&id)
	if Controller.MemberSearch(id).Id != 0 {
		fmt.Print("masukkan Nama Depan : ")
		fmt.Scan(&namaDepan)
		fmt.Print("masukkan Nama Belakang : ")
		fmt.Scan(&namaBelakang)
		fmt.Print("masukkan No TELP : ")
		fmt.Scan(&noTelp)
		member := Controller.MemberUpdate(id, namaDepan, namaBelakang, noTelp)
		fmt.Println("Data berhasil Di update")
		viewAttribute(member)
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}

func MemberSearch() {
	var id int
	fmt.Print("masukkan ID yang akan di Cari : ")
	fmt.Scan(&id)
	member := Controller.MemberSearch(id)
	if member.Id == 0 {
		fmt.Println("ID tidak valid ")
	} else {
		viewAttribute(member)
	}
}
