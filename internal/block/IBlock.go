package block

import (
	content "github.com/tothbence9922/go-blockchain/internal/content"
)

type IBlock interface {
	getTimestamp() uint64
	getLastHash() string
	getHash() string
	getContent() content.Content
	String() string
}
