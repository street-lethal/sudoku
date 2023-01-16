package src

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Tiles    []*Tile
	Rows     map[int]TileSet
	Cols     map[int]TileSet
	Squares  map[int]TileSet
	BlankIDs []int
}

func (b *Board) Init() {
	b.Rows = map[int]TileSet{}
	b.Cols = map[int]TileSet{}
	b.Squares = map[int]TileSet{}
	b.BlankIDs = []int{}
	for i := 0; i < 81; i++ {
		row := i / 9
		col := i % 9
		square := (row/3)*3 + col/3
		tile := b.Tiles[i]
		b.Rows[row] = append(b.Rows[row], tile)
		b.Cols[col] = append(b.Cols[col], tile)
		b.Squares[square] = append(b.Squares[square], tile)
		if tile.IsBlank() {
			b.BlankIDs = append(b.BlankIDs, i)
		} else {
			tile.Default = true
		}
		tile.Row = row
		tile.Col = col
		tile.Square = square
	}
}

func LoadFromFile(fileName string) (*Board, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tiles := []*Tile{}
	for _, record := range records {
		for _, value := range record {
			if value == "" {
				continue
			}
			strDigits := strings.Split(value, "")
			for _, strDigit := range strDigits {
				digit, _ := strconv.Atoi(strDigit)
				tile := Tile{Digit: digit}
				tiles = append(tiles, &tile)
			}
		}
	}

	board := &Board{Tiles: tiles}
	return board, nil
}

func (b *Board) String() string {
	bar := "+-------+-------+-------+\n"
	display := bar
	for row := 0; row < 9; row++ {
		display += fmt.Sprintf(
			"| %s %s %s | %s %s %s | %s %s %s |\n",
			b.Rows[row][0].String(), b.Rows[row][1].String(), b.Rows[row][2].String(),
			b.Rows[row][3].String(), b.Rows[row][4].String(), b.Rows[row][5].String(),
			b.Rows[row][6].String(), b.Rows[row][7].String(), b.Rows[row][8].String(),
		)
		if row%3 == 2 {
			display += bar
		}
	}

	return display
}

func (b *Board) Validate() error {
	for row := 0; row < 9; row++ {
		if !b.Rows[row].IsFilled() {
			return fmt.Errorf(fmt.Sprintf("Row %d is not filled", row))
		}
		if b.Rows[row].Duplicates() {
			return fmt.Errorf(fmt.Sprintf("Row %d duplicates", row))
		}
	}

	for col := 0; col < 9; col++ {
		if !b.Cols[col].IsFilled() {
			return fmt.Errorf(fmt.Sprintf("Col %d is not filled", col))
		}
		if b.Cols[col].Duplicates() {
			return fmt.Errorf(fmt.Sprintf("Col %d duplicates", col))
		}
	}

	for square := 0; square < 9; square++ {
		if !b.Squares[square].IsFilled() {
			return fmt.Errorf(fmt.Sprintf("Square %d is not filled", square))
		}
		if b.Squares[square].Duplicates() {
			return fmt.Errorf(fmt.Sprintf("Square %d duplicates", square))
		}
	}

	return nil
}

func (b *Board) PossibleDigits(tile *Tile) *PossibleDigits {
	digitsRow := b.Rows[tile.Row].PossibleDigits()
	digitsCol := b.Cols[tile.Col].PossibleDigits()
	digitsSquare := b.Squares[tile.Square].PossibleDigits()

	return digitsRow.Multiple(digitsCol).Multiple(digitsSquare)
}
