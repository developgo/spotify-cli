package internal

import (
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	err := play(api)
	require.NoError(t, err)
}
