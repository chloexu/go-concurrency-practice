package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name        string
	latestDraw  int
	latestScore int
	totalScore  int
}

func (p *Player) Play() {
	rand.Seed(time.Now().UnixNano())
	draw := rand.Intn(13) + 1
	var rs int
	switch draw {
	case 1, 3, 5:
		rs = 5
	case 2, 4:
		rs = 2
	case 6, 8:
		rs = -8
	case 7:
		rs = -7
	case 9, 10:
		rs = 2
	case 11, 12:
		rs = 5
	case 13:
		rs = -9
	}
	p.latestDraw = draw
	p.latestScore = rs
	p.totalScore += rs
}

func EndRound(players []*Player, c chan *Player) {
	for _, player := range players {
		c <- player
	}
}

func Monitor(c chan *Player, q chan *Player) {
	for {
		select {
		case player := <-c:
			fmt.Printf("-------- Player %v: | draw [%v] | score [%v] | ====> total score %v\n",
				player.name, player.latestDraw, player.latestScore, player.totalScore)
		case winner := <-q:
			fmt.Printf("\n************* Game over! Winner is %v with a total score of %v, Congrats!\n\n",
				winner.name, winner.totalScore)
			return
		}
	}
}

func AnounceWinner(players []*Player, q chan *Player) {
	var winner *Player
	winnerScore := -99999
	for _, player := range players {
		if player.totalScore > winnerScore {
			winner = player
			winnerScore = player.totalScore
		}
	}
	q <- winner
}

func main() {
	c := make(chan *Player)
	q := make(chan *Player)
	players := []*Player{
		{name: "Chloe", totalScore: 0},
		{name: "Peixuan", totalScore: 0},
		{name: "Chen", totalScore: 0},
		{name: "Waz", totalScore: 0},
		{name: "Laurence", totalScore: 0},
	}
	rounds := make([]int, 10)

	go func() {
		for r := range rounds {
			fmt.Printf("\n************* About to start Round #%v *************\n\n", r+1)
			// for each round, have each player start the round to get the score
			for _, player := range players {
				go player.Play()
			}
			// then end the round after all players finish playing
			EndRound(players, c)
			// sleep for 5 seconds before proceeding to the next round
			time.Sleep(1 * time.Second)
		}
		AnounceWinner(players, q)
	}()
	Monitor(c, q)
}
