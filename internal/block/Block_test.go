package block_test

import (
	"testing"

	block "github.com/tothbence9922/go-blockchain/internal/block"
	"github.com/tothbence9922/go-blockchain/internal/content"
)

var (
	firstBlock block.Block
	newBlock   block.Block
)

func init() {
	firstBlock = block.Genesis()
	newBlock = block.MineBlock(firstBlock, content.Content{Value: 22})
}

func TestGenesis(t *testing.T) {
	blockString := firstBlock.String()
	if blockString != "{\nTimestamp:\t0,\nLastHash:\tthere-is-no-last-hash,\nHash:\tfirst-hash,\nContent:\t0\n}" {
		t.Errorf("Genesis() String() does not match the expected output.")
	}
}

func TestMineBlock(t *testing.T) {
	if newBlock.Content.Value != 22 {
		t.Errorf("Genesis() sets Content.Value incorrectly.")
	}
	if newBlock.LastHash != firstBlock.Hash {
		t.Errorf("Genesis() sets LastHash incorrectly.")
	}
}
