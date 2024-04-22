package Node

type PenjualanNode struct {
	IdPenjualan int
	IdMember    int
	Total       float32
	CreateAt    string
	Detail      []DetailPenjualanNode
}

type DetailPenjualanNode struct {
	IdBarang int
	Jumlah   int
	Harga    float32
}

type PenjualanLL struct {
	Penjualan PenjualanNode
	Next      *PenjualanLL
}
