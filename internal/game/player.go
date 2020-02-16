package game

import (
	"sync"
)

const (
	NoPlayerIndex = 0
)

type Player struct {
	cards map[string]*Card
	Name  string
	index int

	sync.Mutex
}

func NewPlayer(index int, nickname string) *Player {
	return &Player{
		cards: make(map[string]*Card),
		Name:  nickname,
		index: index,
	}
}

func (p *Player) AssignCards(cards []*Card) {
	p.Lock()
	defer p.Unlock()
	for _, card := range cards {
		p.cards[card.Label()] = card
	}
}

func (p *Player) Cards() []*Card {
	var cards []*Card

	for _, card := range p.cards {
		cards = append(cards, card)
	}
	return cards
}

func (p *Player) GetRemains() int {
	return len(p.cards)
}

func (p *Player) PlayCard(card *Card) error {
	p.Lock()
	defer p.Unlock()
	if err := card.Play(); err != nil {
		return err
	}
	delete(p.cards, card.Label())
	return nil
}

func (p *Player) DropCard(card *Card) {
	p.Lock()
	defer p.Unlock()
	card.Drop(p.Name)
	delete(p.cards, card.Label())
}

func (p *Player) DropFirstCard() *Card {
	p.Lock()
	defer p.Unlock()
	if len(p.Cards()) == 0 {
		return nil
	}
	card := p.Cards()[0]
	card.Drop(p.Name)
	delete(p.cards, card.Label())
	return card
}
