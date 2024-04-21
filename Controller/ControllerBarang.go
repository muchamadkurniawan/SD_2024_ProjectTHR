package Controller

import (
	"SD_2024_Project/Model"
	"SD_2024_Project/Node"
)

func BarangInsert(namaBarang string, harga float32, stok int) bool {
	if namaBarang != "" && harga > 0 && stok > 0 {
		Model.BarangInsert(namaBarang, harga, stok)
		return true
	}
	return false
}

func BarangView() []Node.BarangNode {
	barangs := Model.BarangReadAll()
	if barangs == nil {
		return nil
	}
	return barangs
}

func BarangDelete(id int) (bool, Node.BarangNode) {
	barang := Node.BarangNode{}
	if id < 1 {
		return false, barang
	}
	return true, Model.BarangDelete(id)
}

func BarangUpdate(id int, namaBarang string, harga float32, stok int) (bool, Node.BarangNode) {
	barang := Model.BarangUpdate(id, namaBarang, harga, stok)
	if barang.IdBarang == 0 {
		return false, barang
	}
	return true, barang
}

func BarangSearch(id int) (bool, Node.BarangNode) {
	barang := Model.BarangSearch(id)
	if barang.IdBarang == 0 {
		return false, barang
	}
	return true, barang
}

func BarangUpdateStok(id int, stok int) (bool, Node.BarangNode) {
	barang := Model.BarangUpdateStok(id, stok)
	if barang.IdBarang == 0 {
		return false, barang
	}
	return true, barang
}
