package counter

import (
	"github.com/justedd/hwgl-hw-1-1/internal/entity"

	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Counter struct {
	wordRegexp *regexp.Regexp
	logger     *slog.Logger
}

var ErrFileOpen = errors.New("unable to open file")

func New(logger *slog.Logger) (*Counter, error) {
	wordRegexp, err := regexp.Compile("[^a-z]+")

	if err != nil {
		return nil, fmt.Errorf("NewCounter: regexp compile error: %w", err)
	}

	return &Counter{
		wordRegexp: wordRegexp,
		logger:     logger,
	}, nil
}

func (c *Counter) countWords(reader io.Reader) ([]*entity.CountedWord, error) {
	wordMap := make(map[string]*entity.CountedWord)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		raw := scanner.Text()
		word := c.wordRegexp.ReplaceAllString(strings.ToLower(raw), "")

		if word == "" {
			continue
		}

		cw, ok := wordMap[word]

		if !ok {
			cw = &entity.CountedWord{Word: word, Count: 0}
			wordMap[word] = cw
		}

		cw.Count++
	}

	result := make([]*entity.CountedWord, 0, len(wordMap))

	for _, cw := range wordMap {
		result = append(result, cw)
	}

	return result, nil
}

func (c *Counter) countWordsFromFile(filePath string) ([]*entity.CountedWord, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("countWordsFromFile: %w: %v", ErrFileOpen, err)
	}

	defer func() {
		err := file.Close()

		if err != nil {
			c.logger.Error("countWordsFromFile: error while closing file", slog.Any("err", err))
		}
	}()

	return c.countWords(file)
}

func (c *Counter) getTop(n uint, words []*entity.CountedWord) []*entity.CountedWord {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	topN := n
	length := uint(len(words))
	if topN > length {
		c.logger.Warn("number of words is less than topN")
		topN = length
	}

	result := make([]*entity.CountedWord, topN)
	copy(result, words)

	return result
}

func (c *Counter) FileTop(topN uint, filename string) ([]*entity.CountedWord, error) {
	words, err := c.countWordsFromFile(filename)

	if err != nil {
		return nil, fmt.Errorf("FileTop: %w", err)
	}

	return c.getTop(topN, words), nil
}
