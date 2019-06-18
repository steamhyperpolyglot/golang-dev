package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	// Locks and unlocks the map whe we iterate at the end. Strictly speaking,
	// this isn't necessary because we know tha this section won't happen until
	// all files are processed.
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}

// We track words in a struct. We can use a map type, but using struct
// here makes the next refactor easier.
type words struct {
	sync.Mutex // The words struct now inherits the mutex lock.
	found      map[string]int
}

// Create a new words instance
func newWords() *words {
	return &words{found: map[string]int{}}
}

// Track the number of times we have seen the word.
func (w *words) add(word string, n int) {
	// Locks the object, modifies the map, and then unlocks the object
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
