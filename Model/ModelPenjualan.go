package Model

import (
	"SD_2024_Project/Database"
	"SD_2024_Project/Node"
	"time"
)

func PenjualanInsert(idMember int, details []Node.DetailPenjualanNode) {
	var tempLL *Node.PenjualanLL
	tempLL = &Database.DBPenjualan
	newDataPenjulan := Node.PenjualanNode{
		IdPenjualan: PenjualanLastID(),
		IdMember:    idMember,
		Total:       0,
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		Detail:      details,
	}
	if tempLL.Next == nil {
		tempLL.Next = &Node.PenjualanLL{Penjualan: newDataPenjulan, Next: nil}
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &Node.PenjualanLL{Penjualan: newDataPenjulan, Next: nil}
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
