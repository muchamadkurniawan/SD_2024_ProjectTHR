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
