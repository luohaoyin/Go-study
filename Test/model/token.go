package model

type FreshToken struct {
	UserName string `db:username`
	Password string `db:password`
}
type Data struct {
	Status int
	Info string
	RefreshToken string
	Token        string
}