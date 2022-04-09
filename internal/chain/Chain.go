package chain

import (
	block "github.com/tothbence9922/go-blockchain/internal/block"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

type Chain struct {
	Blocks []block.Block
}

func (c Chain) getBlocks() []block.Block {
	return c.Blocks
}

func (c Chain) addBlock(content content.Content) {
	lastBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := block.MineBlock(lastBlock, content)
	c.Blocks = append(c.Blocks, newBlock)
}
