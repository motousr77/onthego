package common

import (
	"log"
)

type SomeNode struct {
	Node  *SomeNode
	Value interface{}
}

func (sn *SomeNode) Print() {
	log.Printf("%+v", sn)
}
