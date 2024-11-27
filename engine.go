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
	from_rank, from_file := squareStringToIndex(from)
	to_rank, to_file := squareStringToIndex(to)

	fmt.Println("Moving Piece From: ", int(from_rank), int(from_file))
	fmt.Println("Moving Piece To: ", to_rank, to_file)

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

// True if the move is psuedo legal.
// Does not exhaustively check for checkmate, check, castle rights, etc.
//
// Pawns are going forward, by 1 or 2 squares, or diagonally for en passant
// Knights are moving in L-Shapes
// Bishops are moving Diagonally
// Rooks are moving Orthogonally
// Queen moves either Diagonally or Orthogonally, but not both
// King moves only one square (or two orthogonally)
func (board Chessboard) IsMovePsuedoLegal(from, to string) bool {
	from_rank, from_file := squareStringToIndex(from)
	to_rank, to_file := squareStringToIndex(to)

	source_piece := board[from_rank][from_file]
	//target_piece := board[to_rank][to_file]

	// Precalculations useful for multiple piece types.
	rank_delta := max(to_rank, from_rank) - min(to_rank, from_rank)
	file_delta := max(to_file, from_file) - min(to_file, from_file)

	switch source_piece {
	case Empty:
		return false
	case WhitePawn, BlackPawn:
		// White Pawn Rank MUST to decrease by 1 or 2
		// Black Pawn Rank MUST to increase by 1 or 2
		correct_forward_for_piece_type := (source_piece == WhitePawn && (to_rank < from_rank)) || (source_piece == BlackPawn && (to_rank > from_rank))
		rank_change_in_bounds := rank_delta == 1 || rank_delta == 2
		// Pawn file MAY increase or decrease by 1 (for en passant)
		file_change_in_bounds := file_delta <= 1
		return rank_change_in_bounds && file_change_in_bounds && correct_forward_for_piece_type
	case WhiteKnight, BlackKnight:
		rank_change_by_two := to_rank == from_rank+2 || to_rank == from_rank-2
		file_change_by_two := to_file == from_file+2 || to_file == from_file-2
		rank_change_by_one := to_rank == from_rank+1 || to_rank == from_rank-1
		file_change_by_one := to_file == from_file+1 || to_file == from_file-1
		// Rank OR file MUST change by 2 squares, and the other by 1 square.
		return (rank_change_by_two && file_change_by_one) || (rank_change_by_one && file_change_by_two)
	case WhiteBishop, BlackBishop:
		// Rank AND file MUST change by an equivalent amount, and 7 at most.
		return (rank_delta <= 7 && file_delta <= 7) && rank_delta == file_delta
	case WhiteRook, BlackRook:
		// TODO: Implement the cases here and below.
		// Rank OR file MUST change by 7 at most, and the other must not change.
		if to_rank == from_rank+1 || to_rank == from_rank-1 {
			return true
		}
		fallthrough
	case WhiteQueen, BlackQueen:
		if to_rank == from_rank+1 || to_rank == from_rank-1 {
			return true
		}
		fallthrough
	case WhiteKing, BlackKing:
		if to_rank == from_rank+1 || to_rank == from_rank-1 {
			return true
		}
	default:
		return false
	}
	return false // Default to false
}

// Returns true if the move is legal according to the rules and game state.
// This is an exhaustive check for checkmate, check, castle rights, etc.
// It essentially requires a review of the entire game history.
func IsMoveStrictlyLegal(piece int, from, to string) bool {
	return false
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
