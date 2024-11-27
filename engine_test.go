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
