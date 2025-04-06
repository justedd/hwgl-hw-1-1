package controller

import (
	"testing"
	"log/slog"
	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"

	"github.com/stretchr/testify/require"
)

func mustCreateController(t *testing.T) *Controller {
	logger := slog.Default()
	counter, err := counter.New(logger)
	require.NoError(t, err)

	controller := New(logger, counter)
	require.NoError(t, err)

	return controller
}

func TestParseArgs(t *testing.T) {
	ctrl := mustCreateController(t)

	t.Run("valid", func(t *testing.T) {
		out, err := ctrl.parseArgs([]string{"topcounter", "3", "foo.txt"})

		require.NoError(t, err)
		require.Equal(t, &Args{FileName: "foo.txt", Top: 3}, out)
	})

	t.Run("wrong filename", func(t *testing.T) {
		_, err := ctrl.parseArgs([]string{"topcounter", "3", ""})

		require.ErrorIs(t, err, ErrWrongFilename)
	})

	t.Run("invalid number of args", func(t *testing.T) {
		_, err := ctrl.parseArgs([]string{"topcounter", "3"})

		require.ErrorIs(t, err, ErrWrongNumberOfArgs)
	})

	t.Run("invalid top", func(t *testing.T) {
		_, err := ctrl.parseArgs([]string{"topcounter", "foo", "file.txt"})

		require.ErrorIs(t, err, ErrWrongTop)
	})
}

