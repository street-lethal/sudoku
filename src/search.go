package src

func Search(board *Board) {
	tiles := []*Tile{}
	for _, blankID := range board.BlankIDs {
		tiles = append(tiles, board.Tiles[blankID])
	}

	walker := NewWalker(&tiles)

	for tile := walker.CurrentTile(); tile != nil; tile = walker.CurrentTile() {
		if walker.SteppedBack {
			walker.StoredPossibleDigits().Remove(tile.Digit)
			tile.Clear()
		} else {
			walker.StorePossibleDigits(board.PossibleDigits(tile))
		}

		walker.SetOnePossibleDigit()

		if tile.IsBlank() {
			walker.StepBack()
			continue
		}

		walker.StepForward()
	}
}
