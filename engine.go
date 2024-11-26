package main

import "fmt"

// Piece Representation
const (
	Empty int = iota // https://go.dev/wiki/Iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

// Board Representation
type Chessboard [8][8]int

func (board *Chessboard) Initialize() {
	*board = Chessboard{
		{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
		{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
		{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
	}
}

func pieceToRune(piece int) rune {
	switch piece {
	case Empty:
		return '.'
	case WhitePawn:
		return 'P'
	case WhiteKnight:
		return 'N'
	case WhiteBishop:
		return 'B'
	case WhiteRook:
		return 'R'
	case WhiteQueen:
		return 'Q'
	case WhiteKing:
		return 'K'
	case BlackPawn:
		return 'p'
	case BlackKnight:
		return 'n'
	case BlackBishop:
		return 'b'
	case BlackRook:
		return 'r'
	case BlackQueen:
		return 'q'
	case BlackKing:
		return 'k'
	default:
		return '?'
	}
}

func (board Chessboard) Print() {
	for rank := 0; rank <= 7; rank++ {
		for file := 0; file < 8; file++ { // A->H
			fmt.Printf("%s ", string(pieceToRune(board[rank][file])))
		}
		fmt.Println() // Newline after each rank
	}
}

// Evaluate() evaluates a chessboard and returns an advantage score.
// Pawns are worth 1, Knights and Bishops are worth 3, Rooks are worth 5, Queens are worth 9.
// Whoever has a higher number has the advantage, expressed as a difference from 0.
// Black is negative, White is positive.
func (board Chessboard) Evaluate() int {
	var score int = 0

	for rank := 0; rank <= 7; rank++ {
		for file := 0; file < 8; file++ {
			switch board[rank][file] {
			case WhitePawn:
				score += 1
			case WhiteKnight, WhiteBishop:
				score += 3
			case WhiteRook:
				score += 5
			case WhiteQueen:
				score += 9
			case BlackPawn:
				score -= 1
			case BlackKnight, BlackBishop:
				score -= 3
			case BlackRook:
				score -= 5
			case BlackQueen:
				score -= 9
			}
		}
	}

	return score
}

func main() {
	var cb Chessboard
	cb.Initialize()
	cb.Print()
	fmt.Println("Position Evaluation: ", cb.Evaluate())
}
