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
		for file := 0; file <= 7; file++ { // A->H
			fmt.Printf("%s ", string(pieceToRune(board[rank][file])))
		}
		fmt.Printf("%d", rank)
		fmt.Println() // Newline after each rank
	}
	fmt.Println("---------------")
	fmt.Println("0 1 2 3 4 5 6 7")
	fmt.Println("a b c d e f g h")
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

func runeToFile(r rune) int {
	// A->H (from left to right)
	file := int(r - 'a')
	if file > 7 {
		file -= 8
	}
	return file
}

// a8 = 0,0
// a1 = 7,0
// h1 = 7,7
// h8 = 0,7
func squareStringToIndex(square string) (int, int) {
	fileRune := rune(square[0])
	rankRune := rune(square[1])

	// Subtract charpoint '0' from charpoint '0'-'9', to get the int value.
	// https://stackoverflow.com/a/21322694
	rankIdx := 7 - (int(rankRune-'0') - 1)
	// Subtract ranks from 7 to invert values 0-7
	// For Indexing Purposes, 7-7=0, 0-7=7
	fileIdx := runeToFile(fileRune)
	return rankIdx, fileIdx
}

func (board *Chessboard) MovePiece(from, to string) bool {

	// Subtract ranks from 7 to invert 0-7
	from_rank, from_file := squareStringToIndex(from)
	to_rank, to_file := squareStringToIndex(to)

	fmt.Println("Moving Piece From: ", int(from_rank), int(from_file))
	fmt.Println("Moving Piece To: ", to_rank, to_file)

	// How can I invert the ranks?
	source_piece := board[from_rank][from_file]
	target_piece := board[to_rank][to_file]

	if target_piece != Empty {
		fmt.Println("We captured a piece: ", target_piece)
	}

	board[from_rank][from_file] = Empty
	board[to_rank][to_file] = source_piece

	newBoard := board

	*board = *newBoard

	return true
}

func main() {
	var cb Chessboard
	cb.Initialize()
	cb.Print()
	fmt.Println("Position Evaluation: ", cb.Evaluate())

	cb.MovePiece("d2", "d4")
	cb.MovePiece("g8", "f6")
	cb.Print()
	fmt.Println("Position Evaluation: ", cb.Evaluate())
}
