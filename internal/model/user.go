package model

type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
}

type UserSignInInput struct {
	UserName string
	Password string
}
