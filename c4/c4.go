package c4

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Starting Game Logistics

type Colour struct {
	slug string
}
type Algo struct {
	slug string
}

var (
	Empty  = Colour{""}
	Red    = Colour{"red"}
	Yellow = Colour{"yellow"}
)
var (
	UnknownAlgo    = Algo{""}
	SimpleAlgo     = Algo{"simple"}
	RandomAlgo     = Algo{"random"}
	MinimaxAlgo    = Algo{"minimax"}
	NormalDistAlgo = Algo{"nd"}
	MonteCarloAlgo = Algo{"monte"}
)

// play starts the game
func Play() {
	Tutorial()
	algo := ChooseAlgo()
	AlgoConfirmation(algo)
	colour := ChooseColour()
	ColourConfirmation(colour)
	// Draw(rowWinBoard)
	// Draw(colWinBoard)
	// Draw(diagwinBLTRBoard)
	// Draw(diagwinTLBRBoard)

	StartGame(colour, algo)
	getPlayerMove(diagwinBLTRBoard, colour)
}

// %%%%%%%%%%%%%%%%%%%%%%% Gameplay Logistics %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

// getPlayerMove(Board, Colour, Algo) is true when Player chooses the move to play on Board.
func getPlayerMove(board [][]Colour, colour Colour) {
	caser := cases.Title(language.English)
	fmt.Print("Your turn as ", caser.String(colour.slug)+". ")
	availableMoves := getAvailableMoves(board)
	Draw(board)

	fmt.Println("moved..!")
	for {
		fmt.Println("Choose from:", availableMoves)
		move := ReadInt()
		if isValidMove(move, availableMoves) {
			break
		}
		fmt.Println("Invalid move. Try again.")
	}
}

// move(Board, computer, Colour, Algo) :-
// getAvailableMoves(Board, AvailableMoves),
// computerMove(Board, Algo, AvailableMoves, Colour, CPUMove),
// newBoard(Board, CPUMove, Colour, NewBoard),
// transitionMove(NewBoard, computer, Colour, Algo).

// isValidMove returns true if the move is in availableMoves and then goes onto transitionMove. If not, it asks the player to try again.
func isValidMove(move int, availableMoves []int) bool {
	return IsMember(move, availableMoves)
}

// getAvailableMoves returns an array of available moves on board
func getAvailableMoves(board [][]Colour) []int {
	availableMoves := make([]int, 0)
	topRow := board[0]
	for i, slot := range topRow {
		if slot.slug == "" {
			availableMoves = append(availableMoves, i+1)
		}
	}
	return availableMoves
}
