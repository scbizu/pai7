package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/pai7/internal/game/i18n"
)

func Join() CommandHandleFunc {
	return CommandHandleFunc(func(msg *api.Message) (api.Chattable, error) {
		g, err := GetGame()
		if err != nil {
			return api.MessageConfig{}, i18n.Err(err)
		}
		if g.IsGameStart() {
			return api.NewMessage(msg.Chat.ID, i18n.NewGameMessageAlreadyStartCNZH()), nil
		}
		index := len(g.GetMembers())
		p := NewPlayer(index, msg.From.UserName)
		g.Join(p)
		config := api.NewMessage(msg.Chat.ID, i18n.NewGameMessageJoinCNZH(g.GetMembers()))
		return config, nil
	})
}
