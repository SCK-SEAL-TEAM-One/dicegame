package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name  string
	score int
}
type Dice struct {
	scoreface int
}

type DiceGame struct {
	player1 Player
	player2 Player
	dice1   Dice
	dice2   Dice
}

func (p *Player) setPlayerName(name string) {
	p.name = name
}
func (p *Player) getPlayerName() string {
	return p.name
}

func (p *Player) setPlayerScore(score int) {
	p.score += score
}

func (p *Player) getPlayerScore() int {
	return p.score
}

func DiceGameConstruct() *DiceGame {
	player1 := new(Player)
	player1.setPlayerName("player1")
	player1.setPlayerScore(0)
	player2 := new(Player)
	player2.setPlayerScore(0)
	player2.setPlayerName("player2")

	dice1 := new(Dice)
	dice2 := new(Dice)

	diceGame := &DiceGame{player1: *player1, player2: *player2, dice1: *dice1, dice2: *dice2}
	return diceGame
}

func (d *Dice) roll() {
	d.scoreface = rand.Intn(5) + 1
}
func (d *Dice) getDiceScore() int {
	return d.scoreface
}

func (w *DiceGame) getPlayerToRoll() *Player {
	if w.player1.getPlayerScore() < w.player2.getPlayerScore() {

		return &w.player1
	} else {

		return &w.player2
	}
}

func (w *DiceGame) getWinner() *Player {
	if w.player1.getPlayerScore() < w.player2.getPlayerScore() {

		return &w.player2
	} else {

		return &w.player1
	}
}

func (w *DiceGame) Play() {
	currentPlayer := &w.player1
	for count := 10; count > 0; count-- {
		w.dice1.roll()
		w.dice2.roll()
		roll1 := w.dice1.getDiceScore()
		roll2 := w.dice2.getDiceScore()
		currentPlayer.setPlayerScore(roll1 + roll2)
		fmt.Println(currentPlayer.getPlayerName(), " ", currentPlayer.getPlayerScore())
		currentPlayer = w.getPlayerToRoll()
	}
	fmt.Println("winner ", currentPlayer.getPlayerName())

}

func main() {

	diceGame := DiceGameConstruct()
	diceGame.Play()
}
