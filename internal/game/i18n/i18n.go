package i18n

type GameMessage uint8

const (
	GameMessageCreate GameMessage = iota
	GameMessageJoin
	GameMessageStart
	GameMessageClose
	GameMessagePlay
	GameMessageDrop
	GameMessageSkip
	GameMessageNextPlayer
	GameMessageAlreadyStart
)
