package entity

type User struct {
	ID       int64
	UUID     string
	Name     string
	Email    string
	Password string
	Active   bool
}
