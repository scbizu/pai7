package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/scbizu/pai7/internal/core"
)

var allCards []*Card
var leftCards []*Card

type Card struct {
	kind   core.Kind
	number core.CardNumber
}

func InitCards() {
	// empty allCards
	allCards = []*Card{}
	leftCards = []*Card{}
	for i := core.MinCardNumber; i < core.MaxCardNumber+1; i++ {
		allCards = append(
			allCards,
			&Card{kind: core.KindRedHeart, number: i},
			&Card{kind: core.KindBlackHeart, number: i},
			&Card{kind: core.KindGrassFlower, number: i},
			&Card{kind: core.KindCube, number: i},
		)
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(allCards), func(i, j int) {
		allCards[i], allCards[j] = allCards[j], allCards[i]
	})
	leftCards = allCards
}

func IfAll7Assigned() bool {
	for _, card := range leftCards {
		if card.number == core.InitCardNumber {
			return false
		}
	}
	return true
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

// GetAvaliableCards returns avaliable cards and if current turn is able to skip
func GetAvaliableCards(cards []*Card) ([]*Card, bool, error) {
	kindCards := make(map[core.Kind][]core.CardNumber)
	for _, card := range cards {
		kindCards[card.kind] = append(kindCards[card.kind], card.number)
	}

	var resCards []*Card
	var isSkipTurn bool
	for kind, numbers := range kindCards {
		ava, err := getCardPool().GetValidNums(kind, numbers)
		if err != nil {
			if errors.Is(err, core.ErrSkipNoFirst7) {
				// skip turn
				isSkipTurn = true
				continue
			}
			return nil, false, err
		}
		numbers := ava[kind]
		for _, num := range numbers {
			resCards = append(resCards, &Card{
				kind:   kind,
				number: num,
			})
		}
	}
	if len(resCards) == 0 && isSkipTurn {
		return resCards, true, nil
	}
	return resCards, false, nil
}

func (c *Card) Drop(user string) {
	getCardPool().AddDroppedNum(
		user,
		c.number,
	)
}

func (c *Card) Play() error {
	return getCardPool().Insert(
		c.kind,
		c.number,
	)
}

func (c Card) Label() string {
	return fmt.Sprintf("%s%s", core.CardKindLabel[c.kind], core.CardNumberLabel[c.number])
}
