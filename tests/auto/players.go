package auto

import (
	"github.com/scbizu/pai7/internal/game"
	"github.com/sirupsen/logrus"
)

var Players = []string{
	"Ji",
	"Ou",
	"Hua",
	"Ya",
	"Song",
}

var AllPlayers []*game.Player

func InitPlayers() {
	for _, p := range Players {
		player := game.NewPlayer(game.NoPlayerIndex, p)
		AllPlayers = append(AllPlayers, player)
	}
}

func ShowPlayersCards() {
	for _, p := range AllPlayers {
		var labels []string
		for _, pc := range p.Cards() {
			labels = append(labels, pc.Label())
		}
		logrus.Infof("auto: players: %s, cards: %v", p.Name, labels)
	}
}

func ShowAvalibaleCards() {
	for _, p := range AllPlayers {
		ava, isSkip, err := game.GetAvaliableCards(p.Cards())
		if err != nil {
			panic(err)
		}
		var labels []string
		for _, pc := range ava {
			labels = append(labels, pc.Label())
		}
		logrus.Infof("auto: ava: players: %s, isSkip turn %v, cards:%v", p.Name, isSkip, labels)
	}
}

func StartPlay() {
	turn := 1
	for {
		logrus.Infof("auto: turn: %d", turn)
		if IfAllPlayersHasNoCards() {
			return
		}
		for _, p := range AllPlayers {
			if p.GetRemains() == 0 {
				continue
			}
			ava, isSkip, err := game.GetAvaliableCards(p.Cards())
			if err != nil {
				panic(err)
			}
			if !isSkip {
				if len(ava) == 0 {
					// drop
					card := p.DropFirstCard()
					logrus.Infof("auto: drop: player %s drops card %s", p.Name, card.Label())
					continue
				}
				if err := p.PlayCard(ava[0]); err != nil {
					panic(err)
				}
				logrus.Infof("auto: play: player %s play card %s", p.Name, ava[0].Label())
				continue
			}
			continue
		}
		logrus.Infof("auto: pool:\n%s", game.PrintPoolStatus())
		turn++
	}
}

func IfAllPlayersHasNoCards() bool {
	for _, p := range AllPlayers {
		if p.GetRemains() > 0 {
			return false
		}
	}
	return true
}
