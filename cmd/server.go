// Copyright © 2020 NAME HERE scbizu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/scbizu/mytg"
	"github.com/scbizu/pai7/internal/game"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "pai7 server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := registerTelegramServer(); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

var (
	isDebug = os.Getenv("IS_DEBUG_MODE")
)

func registerTelegramServer() error {
	var debugMode bool
	if isDebug == "true" {
		debugMode = true
		logrus.SetLevel(logrus.DebugLevel)
	}
	bot, err := mytg.NewBot(debugMode)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}
	bot.RegisterWebhook()
	bot.RegisterMsgChannel(mytg.MSGTypeText, mytg.MSGTypeInline)
	go func() {
		if err := bot.ServeBotUpdateMessage(&game.P7Plugin{}); err != nil {
			logrus.Errorf("register: onUpdate Message: %q", err)
			return
		}
	}()

	if err := bot.ServeInlineMode(game.InlineHandler, game.OnChosenInlineMsgHander); err != nil {
		return fmt.Errorf("register: onUpdate Inline Message: %q", err)
	}
	return nil
}
