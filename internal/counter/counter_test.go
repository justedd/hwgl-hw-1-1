package counter

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/require"
)

func TestGetTop(t *testing.T) {
	var cases = []struct {
		desc string
		n uint
		in  []*CountedWord
		out []*CountedWord
	}{
		{
			desc: "simple",
			n: 2,
			in: []*CountedWord{
				{Word: "a", Count: 5},
				{Word: "b", Count: 50},
				{Word: "c", Count: 1},
				{Word: "d", Count: 17},
			},
			out: []*CountedWord{{Word: "b", Count: 50},	{Word: "d", Count: 17}},
		},
		{
			desc: "N overflow",
			n: 50,
			in: []*CountedWord{{Word: "b", Count: 50}, {Word: "d", Count: 17}},
			out: []*CountedWord{{Word: "b", Count: 50},	{Word: "d", Count: 17}},
		},
		{
			desc: "empty list",
			n: 50,
			in: []*CountedWord{},
			out: []*CountedWord{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T){
			require.Equal(t, tc.out, getTop(tc.n, tc.in))
		})
	}
}

func TestCountWords(t *testing.T) {
	var cases = []struct {
		desc string
		in string
		out []*CountedWord
	}{
		{
			desc: "simple",
			in: ".a    a; b\na,",
			out: []*CountedWord{
				{Word: "a", Count: 3},
				{Word: "b", Count: 1},
			},
		},
		{
			desc: "empty",
			in: "",
			out: []*CountedWord{},
		},
		{
			desc: "spaces",
			in: "   \n   ",
			out: []*CountedWord{},
		},
		{
			desc: "punctuation spaces",
			in: "   ,, ..  a ",
			out: []*CountedWord{{Word: "a", Count: 1}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T){
			words, err :=countWords(strings.NewReader(tc.in))
			require.NoError(t, err)
			require.Equal(t, tc.out, words)
		})
	}
}
