package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer(t *testing.T) {
	g, err := NewGame("scnace", 123456)
	require.NoError(t, err)
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
}
