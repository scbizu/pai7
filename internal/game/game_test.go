package game

import (
	"testing"

	"github.com/scbizu/pai7/internal/core"
	"github.com/stretchr/testify/require"
)

func init() {
	InitGame()
}

func TestPlayer(t *testing.T) {
	g, err := NewGame("scnace", 123456)
	require.NoError(t, err)
	defer g.Close()
	g.Join(NewPlayer(0, "scnace"))
	members := g.GetMembers()
	require.Equal(t, 1, len(members))
	next := g.GetNextPlayer("scnace")
	require.Equal(t, "scnace", next.Name)
	g.Join(NewPlayer(1, "scbizu"))
	members = g.GetMembers()
	require.Equal(t, 2, len(members))
	next = g.GetNextPlayer("scnace")
	require.Equal(t, "scbizu", next.Name)
	g.Close()
}

func TestPlayCard(t *testing.T) {
	g, err := NewGame("scnace", 123456)
	require.NoError(t, err)
	defer g.Close()
	g.Join(NewPlayer(0, "scnace"))
	err = g.PlayerPlaysCard("scnace", &Card{kind: core.KindBlackHeart, number: core.CardNumber(7)})
	require.NoError(t, err)
	status := g.Status()
	require.Contains(t, status, "Kind: 🖤, List: 7")
	require.Contains(t, status, "Kind: ♣️, List:")
	require.Contains(t, status, "Kind: ♦️, List:")
	require.Contains(t, status, "Kind: ❤️, List:")
}

func TestDropCard(t *testing.T) {
	g, err := NewGame("scnace", 123456)
	require.NoError(t, err)
	defer g.Close()
	g.Join(NewPlayer(0, "scnace"))
	// /start
	g.Start()
	rpt := g.GetEndReport()
	require.Contains(t, rpt, "scnace: 0")
	err = g.PlayerDropsCard("scnace", &Card{kind: core.KindBlackHeart, number: core.CardNumber(7)})
	require.NoError(t, err)
	rpt = g.GetEndReport()
	require.Contains(t, rpt, "scnace: 7")
}
