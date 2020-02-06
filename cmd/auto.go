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
	"github.com/scbizu/pai7/internal/game"
	"github.com/scbizu/pai7/tests/auto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Auto play pai7 for testing purpose",
	Run: func(cmd *cobra.Command, args []string) {
		game.InitGame()
		for _, p := range auto.Players {
			gotCards := game.AssignCards(len(auto.Players))
			logrus.Infof("cmd: auto: player: %s,cards: %d", p, len(gotCards))
			auto.PlayersCards[p] = gotCards
		}

		auto.ShowPlayersCards()
	},
}

func init() {
	RootCmd.AddCommand(autoCmd)
}
