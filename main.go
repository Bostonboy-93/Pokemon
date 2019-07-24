package main

import (
	"log"

	"github.com/kr/pretty"

	"github.com/Mrcampbell/teaching-pokemon/Pokemon/app"
	"github.com/Mrcampbell/teaching-pokemon/Pokemon/moveservice"
)

func main() {

	// create a moveservice
	var ms app.MoveService
	ms = moveservice.NewMoveService()

	// get a move from mvoeservice
	move, err := ms.GetMove("1")

	// this is how errors are handled.  It seems clunky at first,
	// but golang doesn't have a .then() or .catch() approach.
	if err != nil {
		log.Fatal("Got an Error from MoveService: ", err)
	}

	// pretty is a package you're guarenteed not to have already installed.
	// I included it so you could learn how to install a package.
	// to install it, run
	// go get github.com/kr/pretty

	// and then print it out
	pretty.Println(move)

}
