package game

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/mytg/plugin"
)

func NewP7Plugin() *P7Plugin {
	return &P7Plugin{}
}

// P7 impls mytg.Plugin interface
type P7Plugin struct {
}

// HandleMessage handles telegram game messages
func (p *P7Plugin) HandleMessage(incommingMsg *api.Message) (*api.MessageConfig, error) {
	if !incommingMsg.IsCommand() {
		return nil, plugin.ErrMessageNotMatched
	}

	cmd, ok := LabelToCommand[incommingMsg.Command()]
	if !ok {
		return nil, plugin.ErrMessageNotMatched
	}

	fn, ok := CommandHandler[cmd]
	if !ok {
		return nil, plugin.ErrMessageNotMatched
	}

	return fn(incommingMsg)
}
