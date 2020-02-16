package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/mytg/plugin"
	"github.com/sirupsen/logrus"
)

func NewP7Plugin() *P7Plugin {
	return &P7Plugin{}
}

// P7 impls mytg.Plugin interface
type P7Plugin struct{}

// HandleMessage handles telegram game messages
func (p *P7Plugin) HandleMessage(incommingMsg *api.Message) (api.Chattable, error) {

	logrus.Debugf("game: incoming message: %s", incommingMsg.Command())

	if !incommingMsg.IsCommand() {
		return api.MessageConfig{}, plugin.ErrMessageNotMatched
	}

	cmd, ok := LabelToCommand[incommingMsg.Command()]
	if !ok {
		return api.MessageConfig{}, plugin.ErrMessageNotMatched
	}

	fn, ok := CommandHandler[cmd]
	if !ok {
		return api.MessageConfig{}, plugin.ErrMessageNotMatched
	}

	return fn(incommingMsg)
}
