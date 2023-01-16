package src

type TileSet []*Tile

func (ts TileSet) IsFilled() bool {
	for _, tile := range ts {
		if tile.IsBlank() {
			return false
		}
	}

	return true
}

func (ts TileSet) Duplicates() bool {
	duplicate := map[int]bool{}
	for _, tile := range ts {
		if tile.IsBlank() {
			continue
		}

		if duplicate[tile.Digit] {
			return true
		}

		duplicate[tile.Digit] = true
	}

	return false
}

func (ts TileSet) PossibleDigits() *PossibleDigits {
	possible := PossibleDigits{
		1: true, 2: true, 3: true, 4: true, 5: true,
		6: true, 7: true, 8: true, 9: true,
	}

	for _, tile := range ts {
		if !tile.IsBlank() {
			possible[tile.Digit] = false
		}
	}

	return &possible
}
