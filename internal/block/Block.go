package block

import (
	"fmt"

	content "github.com/tothbence9922/go-blockchain/internal/content"
)

type Block struct {
	Timestamp uint64
	LastHash  string
	Hash      string
	Content   content.Content
}

func (b Block) getTimestamp() uint64 {
	return b.Timestamp
}

func (b Block) getLastHash() string {
	return b.LastHash
}

func (b Block) getHash() string {
	return b.Hash
}

func (b Block) getContent() content.Content {
	return b.Content
}
func (b Block) String() string {
	return fmt.Sprintf("{\nTimestamp:\t%d,\nLastHash:\t%s,\nHash:\t%s,\nContent:\t%s\n}", b.Timestamp, b.LastHash, b.Hash, b.Content.String())
}
