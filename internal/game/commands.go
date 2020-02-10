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
)

var (
	LabelToCommand = map[string]Command{
		"new":    CMDNewGame,
		"join":   CMDJoinGame,
		"start":  CMDStartGame,
		"end":    CMDEndGame,
		"status": CMDGameStatus,
	}

	CommandsDesc = map[string]string{
		"new":    "创建一局排7游戏",
		"join":   "加入一局排7游戏",
		"start":  "开始一局排7游戏",
		"end":    "结束一局排7游戏",
		"status": "牌局状态",
	}

	CommandHandler = map[Command]CommandHandleFunc{
		CMDNewGame:    New(),
		CMDJoinGame:   Join(),
		CMDStartGame:  Start(),
		CMDEndGame:    Close(),
		CMDGameStatus: Status(),
	}
)
