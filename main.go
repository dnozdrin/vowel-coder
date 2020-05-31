package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var encoderMap = []string{
	"a", "1",
	"e", "2",
	"i", "3",
	"o", "4",
	"u", "5",
}

var decoderMap = []string{
	"1", "a",
	"2", "e",
	"3", "i",
	"4", "o",
	"5", "u",
}

func main() {
	var result string

	decode := flag.Bool("d", false, "Decode.")
	flag.Parse()

	fmt.Println("Please enter the text to process:")
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	if s.Scan() {
		if *decode {
			result = processText(decoderMap, s.Text())
		} else {
			result = processText(encoderMap, s.Text())
		}

		fmt.Println(result)
	}
	if err := s.Err(); err != nil {
		log.Printf("reading input: %s", err)
	}
}

func processText(mapping []string, text string) string {
	r := strings.NewReplacer(mapping...)
	lowerCase := strings.ToLower(text)
	return r.Replace(lowerCase)
}
