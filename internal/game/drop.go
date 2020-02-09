package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/pai7/internal/game/i18n"
)

func Drop() CommandHandleFunc {
	return CommandHandleFunc(func(msg *api.Message) (api.Chattable, error) {
		g, err := GetGame()
		if err != nil {
			return api.MessageConfig{}, i18n.Err(err)
		}
		g.Close()
		return api.NewMessage(msg.Chat.ID, g.Status()), nil
	})
}
