package game

import (
	"fmt"
	"strconv"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scbizu/pai7/internal/core"
	"github.com/scbizu/pai7/internal/game/i18n"
	"github.com/sirupsen/logrus"
)

type ActionType uint8

const (
	ActionTypePlay ActionType = iota
	ActionTypeDrop
	ActionTypeSkip
)

func InlineHandler(msg api.Update) ([]interface{}, error) {
	g, err := GetGame()
	if err != nil {
		item := api.NewInlineQueryResultArticle(
			"game_not_found", "没有加入游戏哦", "没有游戏进行中哦",
		)
		return []interface{}{item}, nil
	}

	user := msg.InlineQuery.From.UserName

	if msg.InlineQuery.Query == "my" {
		var items []interface{}
		for _, card := range g.GetPlayerCards(user) {
			items = append(items, api.NewInlineQueryResultArticle(
				"my_cards",
				fmt.Sprintf("View: %s", card.Label()),
				"viewing my cards...",
			))
		}
		return items, nil
	}

	if g.GetCurrentPlayer().Name != user {
		item := api.NewInlineQueryResultArticle(
			"game_not_your_turn", "还没轮到你出牌哦", "还没轮到你出牌哦",
		)
		return []interface{}{item}, nil
	}

	var items []api.InlineQueryResultArticle
	logrus.Debugf("Inline Handler: user: %s", user)
	cards, isSkip, err := g.GetPlayerAvaliableCards(user)
	if err != nil {
		return nil, err
	}

	// Skip

	if isSkip {
		items = append(items, api.NewInlineQueryResultArticle(
			encodeResultID(
				ActionTypeSkip, user, nil, 0,
			), " 跳过回合", "Skip Turn"))
	}

	// Play

	for idx, card := range cards {
		items = append(items, api.NewInlineQueryResultArticle(
			encodeResultID(ActionTypePlay, user, card, idx),
			fmt.Sprintf("Play: %s", card.Label()),
			card.Label(),
		))
	}

	// Drop
	if len(items) == 0 {
		for idx, card := range g.GetPlayerCards(user) {
			items = append(items, api.NewInlineQueryResultArticle(
				encodeResultID(ActionTypeDrop, user, card, idx),
				fmt.Sprintf("Drop: %s", card.Label()),
				card.Label(),
			))
		}
	}

	var resItems []interface{}

	for _, item := range items {
		resItems = append(resItems, item)
	}

	return resItems, nil
}

func encodeResultID(at ActionType, playerName string, card *Card, index int) string {
	if card == nil {
		return fmt.Sprintf("%d-%s-%d/%d-%d", at, playerName, 0, 0, index)
	}
	return fmt.Sprintf("%d-%s-%d/%d-%d", at, playerName, card.kind, card.number, index)
}

func decodeResultID(id string) (ActionType, string, *Card, error) {
	ps := strings.Split(id, "-")
	if len(ps) != 4 {
		return ActionTypeDrop, "", nil, fmt.Errorf("game: inlineHandler: decode ResultID: expected 4 parts, got %d parts", len(ps))
	}
	t, err := strconv.ParseInt(ps[0], 10, 64)
	if err != nil {
		return ActionTypeDrop, "", nil, fmt.Errorf("game: inlineHandler: decode ResultID: %q", err)
	}
	las := strings.Split(ps[2], "/")
	kind, err := strconv.ParseInt(las[0], 10, 64)
	if err != nil {
		return ActionTypeDrop, "", nil, fmt.Errorf("game: inlineHandler: decode ResultID: %q", err)
	}
	number, err := strconv.ParseInt(las[1], 10, 64)
	if err != nil {
		return ActionTypeDrop, "", nil, fmt.Errorf("game: inlineHandler: decode ResultID: %q", err)
	}
	return ActionType(t), ps[1], &Card{kind: core.Kind(kind), number: core.CardNumber(number)}, nil
}

func OnChosenInlineMsgHander(res *api.ChosenInlineResult, bot *api.BotAPI) error {

	g, err := GetGame()
	if err != nil {
		return err
	}

	logrus.Debugf("onChosenHandler: ResultID:  %s", res.ResultID)

	action, user, card, err := decodeResultID(res.ResultID)
	if err != nil {
		return err
	}

	next := g.GetNextPlayer(user)

	switch action {
	case ActionTypeDrop:
		if err := g.PlayerDropsCard(user, card); err != nil {
			return err
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessageDropCNZH(user))); err != nil {
			return err
		}
		if next == nil {
			return nil
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessageNextPlayerCNZH(next.Name))); err != nil {
			return err
		}
	case ActionTypePlay:
		if err := g.PlayerPlaysCard(user, card); err != nil {
			return err
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessagePlayCNZH(user, card.Label()))); err != nil {
			return err
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessageNextPlayerCNZH(next.Name))); err != nil {
			return err
		}
	case ActionTypeSkip:
		if err := g.PlayerSkipTurn(); err != nil {
			return err
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessageSkipCNZH(user))); err != nil {
			return err
		}
		if _, err := bot.Send(api.NewMessage(g.GetChatID(),
			i18n.NewGameMessageNextPlayerCNZH(next.Name))); err != nil {
			return err
		}
	}

	if g.IsAllPlayerHasNoCard() {
		g.Close()
		if _, err := bot.Send(api.NewMessage(
			g.GetChatID(),
			i18n.NewGameMessageCloseCNZH(
				g.GetEndReport(),
			))); err != nil {
			return err
		}
	}

	return nil
}
