package chain

import (
	block "github.com/tothbence9922/go-blockchain/internal/block"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

type Chain struct {
	Blocks []block.Block
}

func New() *Chain {
	blocks := make([]block.Block, 0)
	blocks = append(blocks, block.Genesis())
	return &Chain{Blocks: blocks}
}

func (c Chain) GetBlocks() []block.Block {
	return c.Blocks
}

func (c *Chain) AddBlock(content content.Content) {
	lastBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := block.MineBlock(lastBlock, content)
	c.Blocks = append(c.Blocks, newBlock)
}

func (c Chain) IsValid() bool {
	if c.Blocks[0].String() != block.Genesis().String() {
		return false
	}

	for i := 1; i < len(c.Blocks); i++ {
		lastBlock := c.Blocks[i-1]
		curBlock := c.Blocks[i]
		if (curBlock.LastHash != lastBlock.Hash) || (curBlock.Hash != block.GenerateHashForBlock(curBlock)) {
			return false
		}
	}
	return true
}
