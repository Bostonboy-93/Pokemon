package app

// https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
type ElementType string

const (
	UNKNOWN  ElementType = "UNKNOWN"
	BUG      ElementType = "BUG"
	DARK     ElementType = "DARK"
	DRAGON   ElementType = "DRAGON"
	ELECTRIC ElementType = "ELECTRIC"
	FAIRY    ElementType = "FAIRY"
	FIGHTING ElementType = "FIGHTING"
	FIRE     ElementType = "FIRE"
	FLYING   ElementType = "FLYING"
	GHOST    ElementType = "GHOST"
	GRASS    ElementType = "GRASS"
	GROUND   ElementType = "GROUND"
	ICE      ElementType = "ICE"
	NORMAL   ElementType = "NORMAL"
	POISON   ElementType = "POISON"
	PSYCHIC  ElementType = "PSYCHIC"
	ROCK     ElementType = "ROCK"
	STEEL    ElementType = "STEEL"
	WATER    ElementType = "WATER"
)
