package Model

import (
	"SD_2024_Project/Database"
	"SD_2024_Project/Node"
	"time"
)

func PenjualanInsert(idMember int, details []Node.DetailPenjualanNode) {

	for i, detail := range details {
		barangTemp := BarangSearch(detail.IdBarang)
		details[i].Harga = barangTemp.Harga
	}

	var tempLL *Node.PenjualanLL
	tempLL = &Database.DBPenjualan
	newDataPenjulan := Node.PenjualanNode{
		IdPenjualan: PenjualanLastID(),
		IdMember:    idMember,
		Total:       getTotalDettail(details),
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		Detail:      details,
	}
	if tempLL.Next == nil {
		tempLL.Next = &Node.PenjualanLL{Penjualan: newDataPenjulan, Next: nil}
		MemberUpdatePoint(idMember, float32(getTotalDettail(details)))
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &Node.PenjualanLL{Penjualan: newDataPenjulan, Next: nil}
		MemberUpdatePoint(idMember, float32(getTotalDettail(details)))
	}
}

func PenjualanLastID() int {
	var tempLL *Node.PenjualanLL
	tempLL = &Database.DBPenjualan

	if tempLL.Next == nil {
		return 1
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		return tempLL.Penjualan.IdPenjualan + 1
	}
}

func getTotalDettail(details []Node.DetailPenjualanNode) float32 {
	var total float32
	for _, detail := range details {
		total += float32(detail.Jumlah) * detail.Harga
	}
	return total
}
