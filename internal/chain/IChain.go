package chain

import (
	block "github.com/tothbence9922/go-blockchain/internal/block"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

type IChain interface {
	GetBlocks() []block.Block
	AddBlock(content content.Content)
}
