package Model

import (
	"SD_2024_Project/Database"
	"SD_2024_Project/Node"
	"time"
)

func MemberInsert(firstName string, lastName string, noTelp string) {
	var tempLL *Node.MemberLL
	tempLL = &Database.DBMember

	member := Node.MemberNode{
		Id:           memberLastId(),
		NamaDepan:    firstName,
		NamaBelakang: lastName,
		NoTelp:       noTelp,
		CraeteAt:     time.Now().Format("2006-1-23"),
	}

	newLL := Node.MemberLL{
		Member: member}
	if tempLL.Next == nil {
		tempLL.Next = &newLL
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newLL
	}
}

func memberLastId() int {
	var tempLL *Node.MemberLL
	tempLL = &Database.DBMember

	if tempLL.Next == nil {
		return 0
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		return tempLL.Member.Id + 1
	}
}

func MemberReadAll() []Node.MemberNode {
	var tempLL *Node.MemberLL
	tempLL = &Database.DBMember
	var memberTable []Node.MemberNode
	for tempLL.Next != nil {
		tempLL = tempLL.Next
		memberTable = append(memberTable, tempLL.Member)
	}
	return memberTable
}

func MemberSearch(id int) *Node.MemberLL {
	var tempLL *Node.MemberLL
	tempLL = &Database.DBMember
	if tempLL.Next != nil {
		for tempLL.Next != nil {
			if tempLL.Next.Member.Id == id {
				return tempLL.Next
			}
			tempLL = tempLL.Next
		}
	} else {
		return nil
	}
	return nil
}
