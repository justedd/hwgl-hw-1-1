package main

import (
	"bufio"
	"errors"
	"fmt"
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

func countWords(filePath string) ([]*CountedWord, error) {
	file, err := os.Open(filePath)

	defer func() {
		err := file.Close()

		if err != nil {
			// TODO: use slog
			fmt.Printf("Error while closing file: %v", err)
		}
	}()

	if err != nil {
		return nil, fmt.Errorf("[countWords] %w: %v", errFileOpen, err)
	}

	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		return nil, fmt.Errorf("[countWords] regexp compile error: %v", err)
	}

	wordMap := make(map[string]*CountedWord)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		raw := scanner.Text()
		word := reg.ReplaceAllString(strings.ToLower(raw), "")

		// TODO: ignore empty words
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

func calculateTop(n uint, words []*CountedWord) []*CountedWord {
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

func main() {
	words, err := countWords("text.txt")

	if err != nil {
		fmt.Println(err)
	}

	top := calculateTop(5, words)

	for i := range top {
		line := fmt.Sprintf("%s: %d", words[i].Word, words[i].Count)
		fmt.Println(line)
	}
}
