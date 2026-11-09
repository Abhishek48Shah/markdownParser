package parser

import (
	"bytes"
	"fmt"
)

const (
	Paragraph int = iota
)

type Node struct {
	Type      int
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func Convert(src []byte,dest *bytes.Buffer) error {
	for val := range  bytes.FieldsSeq(src){
		fmt.Println(string(val))
	}
	return nil
}
