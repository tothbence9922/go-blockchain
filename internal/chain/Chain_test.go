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

// It validates a valid chain
func TestChainIsValid(t *testing.T) {
	blockChain.AddBlock(content.Content{Value: 1})
	valid := blockChain.IsValid()
	if !valid {
		t.Errorf("IsValid() invalidates a valid chain.")
	}
}

// It invalidates a chain with an invalid genesis block
func TestChainIsValidGenesis(t *testing.T) {
	blockChain.Blocks[0].Content = content.Content{Value: -1}
	valid := blockChain.IsValid()
	if valid {
		t.Errorf("IsValid() validates a genesis block with invalid content value.")
	}
}

// It invalidates a chain with an invalid non-genesis block
func TestChainIsValidNonGenesis(t *testing.T) {
	blockChain.AddBlock(content.Content{Value: 22})
	blockChain.Blocks[1].LastHash = "Invalid last hash"
	valid := blockChain.IsValid()
	if valid {
		t.Errorf("IsValid() validates a non-genesis block with invalid last hash.")
	}
	blockChain.AddBlock(content.Content{Value: 22})
	blockChain.Blocks[2].Hash = "Invalid hash"
	valid = blockChain.IsValid()
	if valid {
		t.Errorf("IsValid() validates a non-genesis block with invalid hash.")
	}
}
