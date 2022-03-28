package handlers

type UserOne struct {
	Id      int
	Account int
}

type UserTwo struct {
	Id      int
	Account int
}

type TwoUsers struct {
	First  UserOne
	Second UserTwo
}
