package c4

import (
	"fmt"

	"github.com/fatih/color"
)

func Draw(board [][]Colour) {
	drawHeadings()
	for _, row := range board {
		for _, colour := range row {
			colour.draw()
		}
		fmt.Println()
	}

}

func (c Colour) draw() {
	switch c.slug {
	case "":
		fmt.Print("   ")
	case "red":
		color.New(color.FgRed).Print("R  ")
	case "yellow":
		color.New(color.FgYellow).Print("Y  ")
	}
}

func drawHeadings() {
	fmt.Println("1  2  3  4  5  6  7")
	fmt.Println()
}
