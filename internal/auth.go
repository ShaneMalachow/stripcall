package stripcall

type Role int

const (
	Viewer Role = iota
	Standard
	Admin
	SuperAdmin
)
