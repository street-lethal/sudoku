package src

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		board  *Board
		answer *Board
	}

	board1, err := LoadFromFile("../data/board1_test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	answer1, err := LoadFromFile("../data/answer1_test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	board2, err := LoadFromFile("../data/board2_test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	answer2, err := LoadFromFile("../data/answer2_test.txt")
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				board:  board1,
				answer: answer1,
			},
		},
		{
			name: "2",
			args: args{
				board:  board2,
				answer: answer2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.board.Init()
			tt.args.answer.Init()
			Search(tt.args.board)
			if tt.args.board.String() != tt.args.answer.String() {
				t.Errorf(
					"want:\n%s\ngot:\n%s\n",
					tt.args.answer.String(), tt.args.board.String(),
				)
			}
		})
	}
}
