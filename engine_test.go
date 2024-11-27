package main

import "testing"

func TestChessboardEvaluate(t *testing.T) {
	testCases := []struct {
		boardPosition Chessboard
		expected      int
	}{
		{Chessboard{{WhitePawn, WhiteQueen, BlackRook, BlackRook}}, 0},
		{Chessboard{{BlackPawn, BlackPawn, BlackPawn}}, -3},
		{Chessboard{{BlackQueen, WhiteKing, WhiteQueen, WhiteQueen}}, 9},
	}

	for _, tc := range testCases {
		result := tc.boardPosition.Evaluate()
		if result != tc.expected {
			t.Errorf("cb.Evaluate() returned %d, expected %d", result, tc.expected)
		}
	}
}

/*
a8 = 0,0
a1 = 7,0
h1 = 7,7
h8 = 0,7

r n b q k b n r 0
p p p p p p p p 1
. . . . . . . . 2
. . . . . . . . 3
. . . . . . . . 4
. . . . . . . . 5
P P P P P P P P 6
R N B Q K B N R 7
0 1 2 3 4 5 6 7
*/
func TestSquareStringToIndex(t *testing.T) {
	testCases := []struct {
		square   string
		rank_idx int
		file_idx int
	}{
		{"a1", 7, 0},
		{"a8", 0, 0},
		{"h1", 7, 7},
		{"h8", 0, 7},
	}

	for _, tc := range testCases {
		result_rank_idx, result_file_idx := squareStringToIndex(tc.square)
		if result_rank_idx != tc.rank_idx {
			t.Errorf("squareStringToIndex(%s) returned rank_idx %d, expected rank_idx %d", tc.square, result_rank_idx, tc.rank_idx)
		}
		if result_file_idx != tc.file_idx {
			t.Errorf("squareStringToIndex(%s) returned file_idx %d, expected file_idx %d", tc.square, result_file_idx, tc.file_idx)
		}
	}
}
