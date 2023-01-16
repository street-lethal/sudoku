package src

func Search(board *Board) {
	current := 0

	possibleDigitsMap := map[int]*PossibleDigits{}

	for current < len(board.BlankIDs) {
		tile := board.Tiles[board.BlankIDs[current]]

		if possibleDigitsMap[current] == nil {
			possibleDigitsMap[current] = board.PossibleDigits(tile)
		} else {
			// 戻ってきた
			possibleDigitsMap[current].Remove(tile.Digit)
			tile.Clear()
		}

		for digit := 1; digit <= 9; digit++ {
			if !(*possibleDigitsMap[current])[digit] {
				continue
			}

			tile.Set(digit)
			break
		}

		if !tile.IsBlank() {
			current++
			continue
		}

		possibleDigitsMap[current] = nil
		current--
		if current < 0 {
			return
		}
	}
}
