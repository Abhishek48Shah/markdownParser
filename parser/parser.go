package parser

import (
	"bytes"
	"fmt"
	"strings"
)

type Token struct {
	Type  TokenType
	Mark  string
	Value string
}
type TokenType int

const (
	Heading TokenType = iota
	Bold
	Italic
)

func createToken(markdown []byte) ([]Token, error) {
	var (
		currentIndex int
	)
	var tokens []Token
	length := len(markdown)
	for currentIndex < length {
		var newToken Token
		if markdown[currentIndex] == '#' {
			if currentIndex+1 < length && (markdown[currentIndex+1] == '#' || markdown[currentIndex+1] == ' ') {
				newToken.Type = Heading
				count := 1
				for currentIndex < len(markdown) {
					currentIndex += 1
					if markdown[currentIndex] == '#' {
						count += 1
					} else {
						break
					}
				}
				currentIndex += 1
				newToken.Mark = strings.Repeat("#", count)
				var temp string
				for currentIndex < length {
					if markdown[currentIndex] != '\n' {
						temp += string(markdown[currentIndex])
					} else {
						break
					}
					currentIndex += 1
				}
				newToken.Value = temp
				tokens = append(tokens, newToken)
				newToken = Token{}
			}
		}
		currentIndex += 1
	}
	return tokens, nil
}
func Convert(src []byte, dest *bytes.Buffer) error {
	tokens, err := createToken(src)
	if err != nil {
		return err
	}
	for _, value := range tokens {
		fmt.Println(value.Mark, value.Type, string(value.Value))
	}
	return nil
}
