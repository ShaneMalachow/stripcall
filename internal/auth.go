package stripcall

type Role int

const (
	SuperAdmin Role = iota
	Admin
	Armorer
)
