package user

type Role int

const (
	RoleAdmin Role = iota
	RoleMember
)

type QueryBy int

const (
	QueryById QueryBy = iota
	QueryByName
)