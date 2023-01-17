package src

type Walker struct {
	tiles          []*Tile
	step           int
	len            int
	possibleDigits map[int]*PossibleDigits
	SteppedBack    bool
}

func NewWalker(tiles *[]*Tile) *Walker {
	return &Walker{
		tiles:          *tiles,
		step:           0,
		len:            len(*tiles),
		possibleDigits: map[int]*PossibleDigits{},
	}
}

func (w *Walker) CurrentTile() *Tile {
	if w.step >= w.len || w.step < 0 {
		return nil
	}

	return w.tiles[w.step]
}

func (w *Walker) StorePossibleDigits(possibleDigits *PossibleDigits) {
	w.possibleDigits[w.step] = possibleDigits
}

func (w *Walker) StoredPossibleDigits() *PossibleDigits {
	return w.possibleDigits[w.step]
}

func (w *Walker) SetOnePossibleDigit() {
	for digit := 1; digit <= 9; digit++ {
		if (*w.StoredPossibleDigits())[digit] {
			w.CurrentTile().Set(digit)
			return
		}
	}
}

func (w *Walker) StepForward() {
	w.step++
	w.SteppedBack = false
}

func (w *Walker) StepBack() {
	w.step--
	w.SteppedBack = true
}
