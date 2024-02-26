package gotrackr

type Store interface {
	// Users
	CreateUser() error
}
