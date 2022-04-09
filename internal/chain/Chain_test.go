package chain_test

import (
	"testing"

	block "github.com/tothbence9922/go-blockchain/internal/block"
	chain "github.com/tothbence9922/go-blockchain/internal/chain"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

var (
	blockChain *chain.Chain
	firstBlock block.Block
	newBlock   block.Block
)

func init() {
	firstBlock = block.Genesis()
	newBlock = block.MineBlock(firstBlock, content.Content{Value: 22})
	blockChain = chain.New()
}

func TestNewChain(t *testing.T) {
	if blockChain.Blocks[0] != firstBlock {
		t.Errorf("NewChain does not add the genesis block.")
	}
}
func TestAddBlock(t *testing.T) {
	blockChain.AddBlock(content.Content{Value: 22})
	if blockChain.Blocks[1].Content.Value != 22 {
		t.Errorf("AddBlock does not add the Content Value correctly.")
	}
}
