package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Command int32
type CommandHandleFunc func(*api.Message) (api.Chattable, error)

const (
	// CMDNewGame inits a game (per group)
	CMDNewGame Command = iota // CMDJoinGame represents member joins the game (after /new)
	// cannot join neither game is not exist nor game is started.
	CMDJoinGame
	// CMDStartGame starts the game , member cannot join during the game
	CMDStartGame
	// CMDEndGame closed the game
	CMDEndGame
	// CMDGameStatus show current game status
	CMDGameStatus
	// CMDCoverCard covers the card
	CMDDropCard
	// CMDDoubt call someone's buff
	CMDDoubt
	// CMDMyCards shows my cards
	CMDMyCards
	// CMDPlayCards plays one's card to the card  pool
	CMDPlayCard
	// CMDCommands make py with BotFather
	CMDCommands
)

var (
	CommandToLabel = map[Command]string{
		CMDNewGame:    "/new",
		CMDJoinGame:   "/join",
		CMDStartGame:  "/start",
		CMDEndGame:    "/end",
		CMDGameStatus: "/status",
		CMDDropCard:   "/drop_card",
		CMDDoubt:      "/doubt",
		CMDMyCards:    "/my_cards",
		CMDPlayCard:   "/play_card",
	}

	LabelToCommand = map[string]Command{
		"/new":       CMDNewGame,
		"/join":      CMDJoinGame,
		"/start":     CMDStartGame,
		"/end":       CMDEndGame,
		"/status":    CMDGameStatus,
		"/drop_card": CMDDropCard,
		"/doubt":     CMDDoubt,
		"/my_cards":  CMDMyCards,
		"/play_card": CMDPlayCard,
	}

	CommandHandler = map[Command]CommandHandleFunc{
		CMDNewGame:    New(),
		CMDJoinGame:   Join(),
		CMDStartGame:  Start(),
		CMDEndGame:    Close(),
		CMDGameStatus: Status(),
		CMDDropCard:   nil,
		CMDDoubt:      nil,
		CMDMyCards:    nil,
		CMDPlayCard:   nil,
	}
)
