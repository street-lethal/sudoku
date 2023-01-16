package src

import "strconv"

type Tile struct {
	Digit   int
	Default bool
	Row     int
	Col     int
	Square  int
}

func (t *Tile) IsBlank() bool {
	return t.Digit == 0
}

func (t *Tile) Set(digit int) {
	t.Digit = digit
}

func (t *Tile) Clear() {
	t.Digit = 0
}

func (t *Tile) String() string {
	if t.IsBlank() {
		return " "
	}

	return strconv.Itoa(t.Digit)
}
