package Node

type MemberNode struct {
	Id           int
	NoTelp       string
	NamaDepan    string
	NamaBelakang string
	Point        float32
	CraeteAt     string
}

type MemberLL struct {
	Member MemberNode
	Next   *MemberLL
}
