package game

import (
	"github.com/scbizu/pai7/internal/core"
	"github.com/scbizu/pai7/internal/game/i18n"
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

var currentGame *Game

type Game struct {
	players            map[string]*Player
	creator            string
	isStart            bool
	currentPlayerIndex int
	chatID             int64
}

func GetGame() (*Game, error) {
	if currentGame == nil {
		return nil, i18n.ErrGameNotExisted
	}
	return currentGame, nil
}

func NewGame(creator string, chatID int64) (*Game, error) {
	if currentGame != nil {
		return nil, i18n.ErrGameExisted
	}
	currentGame = newGame(creator, chatID)
	return currentGame, nil
}

func newGame(creator string, chatID int64) *Game {
	return &Game{
		creator: creator,
		chatID:  chatID,
		players: make(map[string]*Player),
	}
}

func (g Game) GetChatID() int64 { return g.chatID }

func (g Game) GetMembers() []string {
	var members []string
	for _, p := range g.players {
		members = append(members, p.Name)
	}
	return members
}

func (g Game) GetFirstPlayer() *Player {
	var stIndex int
	for {
		if len(g.players) == 0 {
			return nil
		}
		for _, p := range g.players {
			if p.index == stIndex {
				return p
			}
		}
		stIndex++
	}
}

func (g Game) IsAllPlayerHasNoCard() bool {
	for _, p := range g.players {
		if len(p.cards) > 0 {
			return false
		}
	}
	return true
}

func (g Game) GetEndReport() string {
	return getCardPool().PrintDropped()
}

func (g *Game) Join(ps ...*Player) {
	for _, p := range ps {
		g.players[p.Name] = p
	}
}

func (g *Game) IsGameStart() bool {
	return g.isStart
}

func (g *Game) Start() {
	g.isStart = true
assign:
	InitGame()
	var pnames []string
	for _, p := range g.players {
		pnames = append(pnames, p.Name)
		gotCards := AssignCards(len(g.GetMembers()))
		p.AssignCards(gotCards)
	}
	if !IfAll7Assigned() {
		goto assign
	}
	getCardPool().SetPlayers(pnames)
}

func (g *Game) GetPlayerCards(name string) []*Card {

	if _, ok := g.players[name]; !ok {
		return []*Card{}
	}

	return g.players[name].Cards()
}

func (g *Game) GetPlayerAvaliableCards(name string) ([]*Card, bool, error) {

	if _, ok := g.players[name]; !ok {
		return []*Card{}, false, nil
	}

	return GetAvaliableCards(g.GetPlayerCards(name))
}

func (g *Game) GetCurrentPlayer() *Player {
	for _, p := range g.players {
		if p.index == g.currentPlayerIndex {
			return p
		}
	}
	return nil
}

func (g *Game) GetNextPlayer(name string) *Player {
	var index int
	for _, p := range g.players {
		if p.Name == name {
			index = p.index
			break
		}
	}

	i := (index + 1) % len(g.GetMembers())
	for _, p := range g.players {
		if p.index == i {
			return p
		}
	}

	return nil
}

func (g *Game) Status() string {
	return getCardPool().PrintStatus()
}

func (g *Game) Close() {
	currentGame = nil
}

func (g *Game) PlayerPlaysCard(name string, card *Card) error {
	defer g.resetGameIndex()
	p, ok := g.players[name]
	if !ok {
		return nil
	}
	return p.PlayCard(card)
}

func (g *Game) PlayerDropsCard(name string, card *Card) error {
	defer g.resetGameIndex()
	p, ok := g.players[name]
	if !ok {
		return nil
	}
	p.DropCard(card)
	return nil
}

func (g *Game) PlayerSkipTurn() error {
	defer g.resetGameIndex()
	return nil
}

func (g *Game) resetGameIndex() {
	g.currentPlayerIndex++
	g.currentPlayerIndex = g.currentPlayerIndex % len(g.GetMembers())
}
