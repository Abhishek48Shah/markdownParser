package parser

import (
	"bytes"
	"fmt"
)

type Token struct {
	Type  string
	Name  string
	Value string
}

func createToken(markdown []byte) ([]Token, error) {
	for val := range bytes.FieldsSeq(markdown) {
		fmt.Println(string(val))
	}
	return nil, nil
}
func Convert(src []byte, dest *bytes.Buffer) error {
	tokens, err := createToken(src)
	if err != nil {
		return err
	}
	fmt.Println(tokens)
	return nil
}
