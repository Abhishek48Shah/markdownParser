package parser

import (
	"bytes"
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}
type TokenType string

const (
	Heading     TokenType = "Heading"
	Bold        TokenType = "Blod"
	Text        TokenType = "Text"
	NewLine     TokenType = "NewLine"
	UnOrderList TokenType = "UnOrderList"
)

func createToken(markdown []byte) ([]Token, error) {
	fmt.Printf("%q", string(markdown))
	var tokens []Token
	length := len(markdown)
	var currIndex int
	for currIndex < length {
		switch markdown[currIndex] {
		case '#':
			newToken := Token{Type: Heading}
			for currIndex+1 < length && (markdown[currIndex+1] == ' ' || markdown[currIndex+1] == '#') {
				newToken.Value += string(markdown[currIndex])
				currIndex += 1
			}
			tokens = append(tokens, newToken)
		case '\n':
			if markdown[currIndex-1] == '\n' {
				newToken := Token{Type: NewLine}
				tokens = append(tokens, newToken)
			}
		case '-':
			if currIndex+1 < length && markdown[currIndex+1] == ' ' {
				newToken := Token{Type: UnOrderList, Value: string(markdown[currIndex])}
				tokens = append(tokens, newToken)
			}
		case '*':
			prevIndex := currIndex
			currIndex += 1
			isClosed := false
			for currIndex < length {
				if markdown[currIndex] == '*'{
					isClosed = true
				}
				currIndex += 1
			}
		default:
			var temp string
			newNode := Token{Type: Text}
			for currIndex < length {
				if markdown[currIndex] == '\n' {
					break
				}

				temp += string(markdown[currIndex])
				currIndex++
			}
			newNode.Value = temp
			tokens = append(tokens, newNode)
		}
		currIndex += 1
	}
	return tokens, nil
}
func Convert(src []byte, dest *bytes.Buffer) error {
	tokens, err := createToken(src)
	if err != nil {
		return err
	}
	for _, value := range tokens {
		fmt.Println(value.Type, string(value.Value))
	}
	return nil
}
