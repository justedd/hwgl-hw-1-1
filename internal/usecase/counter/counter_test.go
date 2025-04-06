package counter

import (
	"testing"
	"strings"

	"github.com/justedd/hwgl-hw-1-1/internal/entity"

	"github.com/stretchr/testify/require"
)

func mustCreateCounter(t *testing.T) *Counter {
	counter, err := NewCounter()
	require.NoError(t, err)

	return counter
}

func TestGetTop(t *testing.T) {
	var cases = []struct {
		desc string
		n uint
		in  []*entity.CountedWord
		out []*entity.CountedWord
	}{
		{
			desc: "simple",
			n: 2,
			in: []*entity.CountedWord{
				{Word: "a", Count: 5},
				{Word: "b", Count: 50},
				{Word: "c", Count: 1},
				{Word: "d", Count: 17},
			},
			out: []*entity.CountedWord{{Word: "b", Count: 50}, {Word: "d", Count: 17}},
		},
		{
			desc: "N overflow",
			n: 50,
			in: []*entity.CountedWord{{Word: "b", Count: 50}, {Word: "d", Count: 17}},
			out: []*entity.CountedWord{{Word: "b", Count: 50}, {Word: "d", Count: 17}},
		},
		{
			desc: "empty list",
			n: 50,
			in: []*entity.CountedWord{},
			out: []*entity.CountedWord{},
		},
	}

	counter := mustCreateCounter(t)

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T){
			require.Equal(t, tc.out, counter.getTop(tc.n, tc.in))
		})
	}
}

func TestCountWords(t *testing.T) {
	var cases = []struct {
		desc string
		in string
		out []*entity.CountedWord
	}{
		{
			desc: "simple",
			in: ".a    a; b\na,",
			out: []*entity.CountedWord{
				{Word: "a", Count: 3},
				{Word: "b", Count: 1},
			},
		},
		{
			desc: "empty",
			in: "",
			out: []*entity.CountedWord{},
		},
		{
			desc: "spaces",
			in: "   \n   ",
			out: []*entity.CountedWord{},
		},
		{
			desc: "punctuation spaces",
			in: "   ,, ..  a ",
			out: []*entity.CountedWord{{Word: "a", Count: 1}},
		},
	}

	for _, tc := range cases {
		counter := mustCreateCounter(t)

		t.Run(tc.desc, func(t *testing.T){
			words, err := counter.countWords(strings.NewReader(tc.in))
			require.NoError(t, err)
			require.Equal(t, tc.out, words)
		})
	}
}
