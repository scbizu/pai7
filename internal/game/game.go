package game

import (
	"github.com/scbizu/pai7/internal/core"
)

var cardPool *core.MSets

func InitGame() {
	// init cards pool
	cardPool = core.NewMSets()
	// init the whole cards
	InitCards()
}

func getCardPool() *core.MSets {
	return cardPool
}

func AssignCards(players int) []*Card {
	total := len(allCards) / players
	return GetRandomCards(total)
}

func PrintPoolStatus() string {
	return getCardPool().PrintStatus()
}
