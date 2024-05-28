package c4

import (
	"fmt"
)

// Tutorial shows you how to play
func Tutorial() {
	fmt.Println("Type the column followed by a period to place a marker. Eg. \"4.\" to place a marker in column 4.")
}

// PlayAgain is text that shows up after a game ends
func PlayAgain() {
	fmt.Println("Type play. to play again!")
}

// ChooseAlgo(FinalAlgo) is true when FinalAlgo is the algorithm chosen by the player
func ChooseAlgo() Algo {
	fmt.Println("Type \"simple.\" to play against a simple AI,")
	fmt.Println("\"random.\" to play against a random AI,")
	fmt.Println("\"minimax.\" to play against a minimax AI,")
	fmt.Println("\"nd.\" to play against an AI that chooses moves along a normal distribution,")
	fmt.Println("and anything else followed by a period to play against Monte Carlo AI.")
	return GetAlgo(Read())
}

func GetAlgo(str string) Algo {
	algoMap := map[string]Algo{
		"simple":  SimpleAlgo,
		"random":  RandomAlgo,
		"minimax": MinimaxAlgo,
		"mm":      MinimaxAlgo,
		"nd":      NormalDistAlgo,
		"monte":   MonteCarloAlgo,
	}
	algo, ok := algoMap[str]
	if !ok {
		return MonteCarloAlgo
	}
	return algo
}

func AlgoConfirmation(algo Algo) {
	messageMap := map[string]string{
		SimpleAlgo.slug:     "You chose simple AI.",
		RandomAlgo.slug:     "You chose random AI.",
		MinimaxAlgo.slug:    "You chose minimax AI. Note: Positive score is better for red. Negative score is better for yellow.",
		NormalDistAlgo.slug: "You chose moves-normally-distributed AI.",
		MonteCarloAlgo.slug: "You chose Monte Carlo AI.",
	}
	fmt.Println(messageMap[algo.slug])
}

// ChooseColour(FinalColour) is true when FinalColour is the colour chosen by the player
func ChooseColour() Colour {
	fmt.Println("Type \"red.\" to play as red or anything else followed by a period to play as yellow.")
	return GetColour(Read())
}

// GetColour(Colour, TypedColour) parses the player's text to get the desired colour to play as
func GetColour(str string) Colour {
	colours := map[string]Colour{
		"red":    Red,
		"yellow": Yellow,
	}
	Colour, ok := colours[str]
	if !ok {
		return Yellow
	}
	return Colour
}

func ColourConfirmation(Colour Colour) {
	messages := map[string]string{
		Red.slug:    "You chose red.",
		Yellow.slug: "You chose yellow.",
	}
	fmt.Println(messages[Colour.slug])
}

// StartGame(Colour, Algo) starts the game for the player as Colour and against Algo
func StartGame(colour Colour, algo Algo) {
	if colour.slug == "red" {
		// move(initialBoard, player, red, algo)
	} else {
		// move(initialBoard, computer, red, algo)
	}
}
