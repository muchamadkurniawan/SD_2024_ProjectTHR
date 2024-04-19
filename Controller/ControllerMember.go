package Controller

import (
	"SD_2024_Project/Model"
	"SD_2024_Project/Node"
)

func MemberInsert(firstName string, lastName string, noTelp string) bool {
	if firstName != "" && noTelp != "" {
		Model.MemberInsert(firstName, lastName, noTelp)
		return true
	}
	return false
}

func MembersView() []Node.MemberNode {
	members := Model.MemberReadAll()
	if members == nil {
		return nil
	}
	return members
}

func MemberDelete(id int) (bool, Node.MemberNode) {
	member := Node.MemberNode{}
	if id < 1 {
		return false, member
	}
	return true, Model.MemberDelete(id)
}

func MemberUpdate(id int, namaDepan string, namaBlk string, noTelp string) Node.MemberNode {
	return Model.MemberUpdate(id, namaDepan, namaBlk, noTelp)
}

func MemberSearch(id int) Node.MemberNode {
	address := Model.MemberSearch(id)
	if address != nil {
		return address.Next.Member
	}
	return Node.MemberNode{}
}
