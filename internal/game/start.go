package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/pai7/internal/game/i18n"
)

func Start() CommandHandleFunc {
	return CommandHandleFunc(func(msg *api.Message) (api.Chattable, error) {
		g, err := GetGame()
		if err != nil {
			return api.MessageConfig{}, i18n.Err(err)
		}
		g.Start()
		return api.NewMessage(msg.Chat.ID, i18n.NewGameMessageStartCNZH()), nil
	})
}
