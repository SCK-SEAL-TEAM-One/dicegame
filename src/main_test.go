package main

import "testing"

func Test_getPlayerName_Input_Player1_Should_be_Player1(t *testing.T) {
	name := "player1"
	expected := "player1"
	player1 := new(Player)
	player1.setPlayerName(name)
	actual := player1.getPlayerName()

	if expected != actual {

		t.Errorf("expected is %s but got %s", expected, actual)
	}

}

func Test_getPlayerScore_Input_4_Should_be_4(t *testing.T) {
	score := 4
	expected := 4

	score1 := new(Player)
	score1.setPlayerScore(score)
	actual := score1.getPlayerScore()
	if expected != actual {

		t.Errorf("expected is %d but got %d", expected, actual)
	}

}
func Test_getDiceScore_Should_Be_Between_1_to_6(t *testing.T) {

	expected1 := 1
	expected2 := 6
	dice := new(Dice)
	dice.roll()

	actual := dice.getDiceScore()
	if actual < expected1 && actual > expected2 {

		t.Errorf("actual is %d but got between %d and %d", actual, expected1, expected2)
	}

}

func Test_getWinner_Should_Be_Player(t *testing.T) {
	player1 := new(Player)
	player1.setPlayerName("player1")
	player1.setPlayerScore(20)
	player2 := new(Player)
	player2.setPlayerName("player2")
	player2.setPlayerScore(15)
	dice1 := new(Dice)
	dice2 := new(Dice)
	expected := player1.getPlayerName()

	diceGame := &DiceGame{player1: *player1, player2: *player2, dice1: *dice1, dice2: *dice2}

	actual := diceGame.getWinner().getPlayerName()

	if expected != actual {

		t.Errorf("expected is %s but got %s", expected, actual)
	}
}
