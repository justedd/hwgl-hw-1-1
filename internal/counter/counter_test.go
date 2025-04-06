package counter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTop(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		in := []*CountedWord{
			{Word: "a", Count: 5},
			{Word: "b", Count: 50},
			{Word: "c", Count: 1},
			{Word: "d", Count: 17},
		}

		expected := []*CountedWord{
			{Word: "b", Count: 50},
			{Word: "d", Count: 17},
		}

		require.Equal(t, expected, getTop(2, in))
	})

	t.Run("big N", func(t *testing.T) {
		in := []*CountedWord{
			{Word: "b", Count: 50},
		}

		expected := []*CountedWord{
			{Word: "b", Count: 50},
		}

		require.Equal(t, expected, getTop(20, in))
	})

	t.Run("empty list", func(t *testing.T) {
		in := []*CountedWord{}

		expected := []*CountedWord{}

		require.Equal(t, expected, getTop(20, in))
	})
}
