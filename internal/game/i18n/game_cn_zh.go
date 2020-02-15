package i18n

import (
	"fmt"
)

var (
	gameMessageCNZH = map[GameMessage]string{
		GameMessageCreate:     "新的一局排7游戏被@%s创建，输入`/join`加入(包括%s哦)",
		GameMessageJoin:       "欢迎加入排7游戏，当前玩家列表: %v",
		GameMessageStart:      "排7游戏开始,请%s出牌",
		GameMessageClose:      "排7游戏结束",
		GameMessagePlay:       "@%s 出了 %s",
		GameMessageDrop:       "@%s 盖了一张牌",
		GameMessageSkip:       "@%s 跳过了该回合",
		GameMessageNextPlayer: "👉 轮到下一个玩家@%s出牌(请圈pai7 bot以获得出牌详情)",
	}
)

func NewGameMessageCreateCNZH(creator string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageCreate], creator, creator)
}

func NewGameMessageJoinCNZH(members []string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageJoin], members)
}

func NewGameMessageStartCNZH(player string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageStart], player)
}

func NewGameMessageCloseCNZH() string {
	return gameMessageCNZH[GameMessageClose]
}

func NewGameMessagePlayCNZH(user string, label string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessagePlay], user, label)
}

func NewGameMessageDropCNZH(user string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageDrop], user)
}

func NewGameMessageSkipCNZH(user string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageSkip], user)
}

func NewGameMessageNextPlayerCNZH(user string) string {
	return fmt.Sprintf(gameMessageCNZH[GameMessageNextPlayer], user)
}
