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

func TestIsMovePseudoLegal(t *testing.T) {
	var cb Chessboard
	cb.Initialize() // From the starting position.

	testCases := []struct {
		from_square string
		to_square   string
		expected    bool
	}{
		// Pawns (3 kinds of move)
		{"d2", "d3", true},  // Pawn 1 square forward
		{"d2", "d4", true},  // Pawn 2 squares forward
		{"d2", "c3", true},  // Pawn diagonally forward
		{"d2", "d2", false}, // Pawn stays in the same square
		{"d2", "d1", false}, // Backwards
		{"d2", "a1", false}, // Backwards and too many files away
		{"d2", "f3", false}, // Diagnoally too many files away
		{"d2", "d8", false}, // Forwards and too many ranks away
		// Knights (8 move directions, 12 pseudo-legal moves from starting position)
		// TODO: Implement more test cases for the remaining pieces below...
		{"b1", "a3", true},
		{"b1", "c3", true},
		{"b1", "d2", true},
		// Bishops (4 move directions, 8 pseudo-legal moves from starting position)
		{"c1", "d2", true},
		{"c1", "a3", true},
		{"c1", "h6", true},
	}

	for _, tc := range testCases {
		result := cb.IsMovePsuedoLegal(tc.from_square, tc.to_square)
		if result != tc.expected {
			t.Errorf("cb.IsMovePsuedoLegal(%s, %s) returned %t, expected %t", tc.from_square, tc.to_square, result, tc.expected)
		}
	}

}
