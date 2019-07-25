package app

// this is an interface.  It doesn't do anything, but it's a contract
// of behavior.  It says the minimum functionality of a struct that will inherit this.
// https://gobyexample.com/interfaces
type MoveService interface {
	GetMove(id string) (Move, error)
}
