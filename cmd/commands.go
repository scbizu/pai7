// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"bytes"
	"fmt"

	"github.com/scbizu/pai7/internal/game"
	"github.com/spf13/cobra"
)

// commandsCmd represents the commands command
var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "show bot avaliable commands",
	Run: func(cmd *cobra.Command, args []string) {
		var allDesc bytes.Buffer
		for c, desc := range game.CommandsDesc {
			allDesc.WriteString(fmt.Sprintf("%s - %s\n", c, desc))
		}
		fmt.Printf("%s", allDesc.String())
	},
}

func init() {
	RootCmd.AddCommand(commandsCmd)
}
