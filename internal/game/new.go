package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/pai7/internal/game/i18n"
)

func New() CommandHandleFunc {
	return CommandHandleFunc(func(msg *api.Message) (api.Chattable, error) {
		creator := msg.From.UserName
		_, err := NewGame(creator, msg.Chat.ID)
		if err != nil {
			return api.MessageConfig{}, i18n.Err(err)
		}
		newGameStr := i18n.NewGameMessageCreateCNZH(creator)
		config := api.NewMessage(msg.Chat.ID, newGameStr)
		return config, nil
	})
}
