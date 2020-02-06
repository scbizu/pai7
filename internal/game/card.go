package game

import (
	"fmt"
	"math/rand"

	"github.com/scbizu/pai7/internal/core"
)

var allCards []*Card
var leftCards []*Card

type Card struct {
	kind   core.Kind
	number core.CardNumber
}

func InitCards() {
	for i := core.MinCardNumber; i < core.MaxCardNumber+1; i++ {
		allCards = append(
			allCards,
			&Card{kind: core.KindRedHeart, number: i},
			&Card{kind: core.KindBlackHeart, number: i},
			&Card{kind: core.KindGrassFlower, number: i},
			&Card{kind: core.KindCube, number: i},
		)
	}
	rand.Shuffle(len(allCards), func(i, j int) {
		allCards[i], allCards[j] = allCards[j], allCards[i]
	})
	leftCards = allCards
}

func NewCard() *Card {
	return &Card{}
}

func GetRandomCards(total int) []*Card {
	if len(allCards) == 0 {
		InitCards()
	}
	if total > len(leftCards) {
		return leftCards
	}
	assigned := leftCards[:total]
	leftCards = leftCards[total:]
	return assigned
}

func (c *Card) Drop(kind core.Kind, num core.CardNumber) {
	getCardPool().AddDroppedNum(kind, num)
}

func (c *Card) Play(kind core.Kind, num core.CardNumber) error {
	return getCardPool().Insert(
		kind,
		num,
	)
}

var CardNumberLabel = map[core.CardNumber]string{
	1:  "A",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "J",
	12: "Q",
	13: "K",
}

var CardKindLabel = map[core.Kind]string{
	core.KindRedHeart:    "❤️",
	core.KindBlackHeart:  "🖤",
	core.KindGrassFlower: "♣️",
	core.KindCube:        "♦️",
}

func (c Card) Label() string {
	return fmt.Sprintf("%s%s", CardKindLabel[c.kind], CardNumberLabel[c.number])
}
