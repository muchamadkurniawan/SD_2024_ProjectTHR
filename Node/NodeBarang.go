package Node

type BarangNode struct {
	IdBarang   int
	NamaBarang string
	Harga      float32
	Stok       int
}

type BarangLL struct {
	Barang BarangNode
	Next   *BarangLL
}
