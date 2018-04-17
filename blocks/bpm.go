package blocks

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"
)

// BPM is a blockchain impl centered on BPM recording
type BPM struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var _ Block = (*BPM)(nil)

// CalculateHash creates the BPM hash via fields and SHA256
func (b *BPM) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + string(b.BPM) + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock creates a new block from the reciever and input
func (b *BPM) GenerateBlock(i interface{}) (Block, error) {
	bpm, ok := i.(int)
	if !ok {
		return nil, errors.New("The input could not be properly cast to create a new Block")
	}
	var newBlock = BPM{Index: b.Index + 1,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		PrevHash:  b.Hash,
	}
	newBlock.Hash = newBlock.CalculateHash()
	return &newBlock, nil
}

// IsBlockValid runs validation checks against a block
func (b *BPM) IsBlockValid(newBlock Block) bool {
	newBPM, ok := newBlock.(*BPM)
	if !ok {
		log.Print("Not ok!")
		return false
	}

	if b.Index+1 != newBPM.Index {

		log.Print("Index bad!")
		return false
	}

	if b.Hash != newBPM.PrevHash {
		log.Print("Prev Hash bad!")
		return false
	}

	log.Print("A-ok!")
	return true
}

// BPMGenesisBlock returns a BPM Block to start a chain
func BPMGenesisBlock() Block {
	return &BPM{
		Index:     0,
		Timestamp: time.Now().String(),
		BPM:       0,
		PrevHash:  "",
		Hash:      ""}
}
