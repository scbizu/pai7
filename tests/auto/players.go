package auto

import (
	"github.com/scbizu/pai7/internal/game"
	"github.com/sirupsen/logrus"
)

var Players = []string{
	"Ji",
	"Ou",
	"HuaHua",
	"YAYA",
	"Song",
}

var PlayersCards map[string][]*game.Card = make(map[string][]*game.Card)

func ShowPlayersCards() {
	for p, pcs := range PlayersCards {
		var labels []string
		for _, pc := range pcs {
			labels = append(labels, pc.Label())
		}
		logrus.Infof("auto: players: %s, cards: %v", p, labels)
	}
}
