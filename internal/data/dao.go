package data

type DAO interface {
	FindUser(username string) (*User, error)
}
