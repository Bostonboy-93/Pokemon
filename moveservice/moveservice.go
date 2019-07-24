package moveservice

import (
	"errors"

	"github.com/Mrcampbell/teaching-pokemon/Pokemon/app"
)

// MoveService uses a map[string]app.Move to store and retrieve moves by id.
// this will be moved into a database, but for now, I'm just doing it "in-memory"
// as they say.
type MoveService struct {
	// lowercase properties are "private", as they can't be read or accessed outside this struct
	moves map[string]app.Move
}

// this is how golang does constructors.  You have a function NewYadaYada() that has
// everything needed to return a new instance of the struct.  If you need a database connection,
// or an http client or something, you would pass that in, like NewYadaYada(db *sql.DB).
func NewMoveService() app.MoveService {
	movelist := make(map[string]app.Move)

	movelist["1"] = app.Move{Name: "Pound", Type: app.NORMAL, Power: 40}
	movelist["2"] = app.Move{Name: "Karate Chop", Type: app.FIGHTING, Power: 50}
	movelist["3"] = app.Move{Name: "Double Slap", Type: app.NORMAL, Power: 15}

	return &MoveService{moves: movelist}
}

// this is how golang does Class Functions
// https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
// https://gobyexample.com/methods
// https://tour.golang.org/methods/3
func (ms *MoveService) GetMove(id string) (app.Move, error) {

	// accessing a map by key is dangerous. If the map doesn't have a value
	// at the key provided (lets say you have 0-4, but you ask for for 5) the
	// program will fail.  the "ok" is a variable passed back where we can
	// check to make sure there was a value.  If not, it's false.
	move, ok := ms.moves[id]
	if ok {
		return move, nil
	}
	// if we didn't find a move (ok == false), return an error
	return app.Move{}, errors.New("Could not find Move with ID of " + id)
}
