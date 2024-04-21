package Model

import (
	"SD_2024_Project/Database"
	"SD_2024_Project/Node"
)

func barangLastId() int {
	var tempLL *Node.BarangLL
	tempLL = &Database.DBBarang

	if tempLL.Next == nil {
		return 1
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		return tempLL.Barang.IdBarang + 1
	}
}

func BarangInsert(namaBarang string, harga float32, stok int) {
	var tempLL *Node.BarangLL
	tempLL = &Database.DBBarang

	var newBarang Node.BarangNode
	newBarang.IdBarang = barangLastId()
	newBarang.NamaBarang = namaBarang
	newBarang.Harga = harga
	newBarang.Stok = stok

	for tempLL.Next != nil {
		tempLL = tempLL.Next
	}

	tempLL.Next = &Node.BarangLL{Barang: newBarang, Next: nil}
}

func BarangReadAll() []Node.BarangNode {
	var tempLL *Node.BarangLL
	tempLL = &Database.DBBarang
	var barangTable []Node.BarangNode
	for tempLL.Next != nil {
		tempLL = tempLL.Next
		barangTable = append(barangTable, tempLL.Barang)
	}
	return barangTable
}

func barangSearchAddress(id int) *Node.BarangLL {
	var tempLL *Node.BarangLL
	tempLL = &Database.DBBarang
	for tempLL.Next != nil {
		if tempLL.Next.Barang.IdBarang == id {
			return tempLL
		}
		tempLL = tempLL.Next
	}
	return nil
}

func BarangDelete(id int) Node.BarangNode {
	var tempLL *Node.BarangLL
	tempLL = barangSearchAddress(id)
	BarangDelete := tempLL.Next.Barang
	if tempLL != nil {
		tempLL.Next = tempLL.Next.Next
	}
	return BarangDelete
}

func BarangUpdate(id int, namaBarang string, harga float32, stok int) Node.BarangNode {
	var tempLL *Node.BarangLL
	tempLL = barangSearchAddress(id)
	barangupdate := tempLL.Next.Barang
	if tempLL != nil {
		tempLL.Next.Barang.NamaBarang = namaBarang
		tempLL.Next.Barang.Harga = harga
		tempLL.Next.Barang.Stok = stok
	}
	return barangupdate
}

func BarangSearch(id int) Node.BarangNode {
	tempLL := barangSearchAddress(id)
	if tempLL != nil {
		return tempLL.Next.Barang
	}
	return Node.BarangNode{}
}

func BarangUpdateStok(id int, amount int) Node.BarangNode {
	tempLL := barangSearchAddress(id)
	if tempLL != nil {
		tempLL.Next.Barang.Stok = tempLL.Next.Barang.Stok + amount
	}
	return tempLL.Next.Barang
}
