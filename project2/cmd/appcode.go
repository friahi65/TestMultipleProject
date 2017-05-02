package cmd

import (
	"fmt"
)

// Game allocates rolls an
type Game struct {
	rolls   []int
	current int
}

// NewGame allocates and starts a new game of bowling.
func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, maxThrowsPerGame)
	return game
}

// Roll rolls the ball and knocks down the number of pins specified by pins.
func (selftori *Game) Roll(pins int) {
	selftori.rolls[selftori.current] = pins
	selftori.current++
}

// Score calculates and returns the player's current score.
func (selftori *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if selftori.isStrike(throw) {
			sum += selftori.strikeBonusFor(throw)
			throw++
		} else if selftori.isSpare(throw) {
			sum += selftori.spareBonusFor(throw)
			throw += 2
		} else {
			sum += selftori.framePointsAt(throw)
			throw += 2
		}
	}
	return sum
}

// isStrike determines if a given throw is a strike or not. A strike is knocking
// down all pins in one throw.
func (selftori *Game) isStrike(throw int) bool {
	return selftori.rolls[throw] == allPins
}

// strikeBonusFor calculates and returns the strike bonus for a throw.
func (selftori *Game) strikeBonusFor(throw int) int {
	return allPins + selftori.framePointsAt(throw+1)
}

// isSpare determines if a given frame is a spare or not. A spare is knocking
// down all pins in one frame with two throws.
func (selftori *Game) isSpare(throw int) bool {
	return selftori.framePointsAt(throw) == allPins
}

// spareBonusFor calculates and returns the spare bonus for a throw.
func (selftori *Game) spareBonusFor(throw int) int {
	return allPins + selftori.rolls[throw+2]
}

// framePointsAt computes and returns the score in a frame specified by throw.
func (selftori *Game) framePointsAt(throw int) int {
	return selftori.rolls[throw] + selftori.rolls[throw+1]
}

// testing utilities:

func (selftori *Game) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		selftori.Roll(pins)
	}
}
func (selftori *Game) rollSpare() {
	selftori.Roll(5)
	selftori.Roll(5)
}
func (selftori *Game) rollStrike() {
	selftori.Roll(10)
}

const (
	// allPins is the number of pins allocated per fresh throw.
	allPins = 10

	// framesPerGame is the number of frames per bowling game.
	framesPerGame = 10

	// maxThrowsPerGame is the maximum number of throws possible in a single game.
	maxThrowsPerGame = 21
)

// Main s the entry point to the application code.
// Implements whatever you need there.
func Main() {
	// Write the application code here.
	fmt.Println("Yo, I'm da app!👍 ")
	return
}
