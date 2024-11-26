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
	// TODO: Use names rank and file.
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

func main() {
	var cb Chessboard
	cb.Initialize()
	cb.Print()
}
