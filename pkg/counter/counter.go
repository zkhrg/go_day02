package counter

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"unicode"
)

type Opt struct {
	Word bool
	Line bool
	Symb bool
}

func innerProcess(opt Opt, filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	var charCount, wordCount, lineCount int
	inWord := false

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := scanner.Text()
		charCount++

		if char == "\n" {
			lineCount++
		}

		if unicode.IsSpace(rune(char[0])) {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return
	}
	res_str := ""
	if opt.Line {
		res_str += fmt.Sprintf("%d", lineCount)
	} else if opt.Symb {
		res_str += fmt.Sprintf("%d", charCount)
	} else if opt.Word {
		res_str += fmt.Sprintf("%d", wordCount)
	}
	fmt.Println(res_str + "\t" + filename)
}

func Process(opt Opt, files []string) {
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go innerProcess(opt, file, &wg)
	}

	wg.Wait()
}
