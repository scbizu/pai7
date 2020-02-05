package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func New() CommandHandleFunc {
	return CommandHandleFunc(func(msg *api.Message) (*api.MessageConfig, error) {
		return nil, nil
	})
}
