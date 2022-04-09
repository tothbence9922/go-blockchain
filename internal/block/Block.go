package block

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	content "github.com/tothbence9922/go-blockchain/internal/content"
)

type Block struct {
	Timestamp int64
	LastHash  string
	Hash      string
	Content   content.Content
}

func (b Block) getTimestamp() int64 {
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

func Genesis() Block {
	return Block{Timestamp: 0, LastHash: "there-is-no-last-hash", Hash: "first-hash", Content: content.Content{Value: 0}}
}

func MineBlock(prevBlock Block, c content.Content) Block {
	timestamp := time.Now().Unix()
	lastHash := prevBlock.getHash()
	hash := GenerateHash(timestamp, lastHash, c)

	return Block{Timestamp: int64(timestamp), LastHash: lastHash, Hash: hash, Content: c}
}

func GenerateHash(timestamp int64, lastHash string, content content.Content) string {
	hash := sha256.New()
	input := make([]byte, 0)
	input = append(input, []byte(fmt.Sprint(timestamp))...)
	input = append(input, lastHash...)
	input = append(input, []byte(content.String())...)
	hash.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
