package chain

import (
	"github.com/bauerc/GoBlockchain/blocks"
	"github.com/davecgh/go-spew/spew"
)

// Chain is the global containing the current blockchain
type Chain struct {
	Blocks []blocks.Block
}

var _ Blockchain = (*Chain)(nil)

// Length checks the length of the blocks
func (c *Chain) Length() int {
	return len(c.Blocks)
}

// AddBlock appends a single block to the chain
func (c *Chain) AddBlock(b blocks.Block) {
	c.Blocks = append(c.Blocks, b)
}

// LatestBlock grabs the latest block on the chain
func (c *Chain) LatestBlock() blocks.Block {
	return c.Blocks[c.Length()-1]
}

// GenerateBlock generates a new block for the chain given the underlying blocks implementation
func (c *Chain) GenerateBlock(i interface{}) (blocks.Block, error) {
	newBlock, err := c.LatestBlock().GenerateBlock(i)
	if err != nil {
		return nil, err
	}
	return newBlock, nil
}

// Genesis returns a blockchain generated with the func provided
func Genesis(genesisBlock func() blocks.Block) Blockchain {
	var blockchain Chain
	block := genesisBlock()
	spew.Dump(block)
	blockchain.AddBlock(block)
	return &blockchain
}

// Blockchain contains blockchain implementations
type Blockchain interface {
	Length() int
	AddBlock(b blocks.Block)
	LatestBlock() blocks.Block
	GenerateBlock(i interface{}) (blocks.Block, error)
}
