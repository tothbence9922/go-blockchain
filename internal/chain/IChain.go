package chain

import (
	block "github.com/tothbence9922/go-blockchain/internal/block"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

type IChain interface {
	getBlocks() []block.Block
	addBlock(content content.Content)
}
