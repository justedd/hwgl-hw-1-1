package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

var errFileOpen = errors.New("unable to open file")

type CountedWord struct {
	Word  string
	Count uint
}

func countWords(reader io.Reader) ([]*CountedWord, error) {
	reg, err := regexp.Compile("[^a-z]+")

	if err != nil {
		return nil, fmt.Errorf("[countWords] regexp compile error: %v", err)
	}

	wordMap := make(map[string]*CountedWord)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		raw := scanner.Text()
		word := reg.ReplaceAllString(strings.ToLower(raw), "")

		if word == "" {
			continue
		}

		cw, ok := wordMap[word]

		if !ok { 
			cw = &CountedWord{Word: word, Count: 0}
			wordMap[word] = cw
		}

		cw.Count += 1
	}

	result := make([]*CountedWord, 0, len(wordMap))

	for _, cw := range wordMap {
		result = append(result, cw)
	}

	return result, nil
}

func countWordsFromFile(filePath string) ([]*CountedWord, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("[countWords] %w: %v", errFileOpen, err)
	}

	defer func() {
		err := file.Close()

		if err != nil {
			// TODO: use slog
			fmt.Printf("Error while closing file: %v", err)
		}
	}()

	return countWords(file)
}

func getTop(n uint, words []*CountedWord) []*CountedWord {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	topN := n
	length := uint(len(words))
	if topN > length {
		topN = length
	}

	result := make([]*CountedWord, topN)
	copy(result, words)

	return result
}

func FileTop(topN uint, filename string) []*CountedWord {
	words, err := countWordsFromFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	return getTop(topN, words)
}
