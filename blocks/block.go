package blocks

//Block defines the functionality needed to perform Blockchain operations
type Block interface {
	CaculateHash() string
	GenerateBlock(i interface{}) (Block, error)
	IsBlockValid(newBlock Block) bool
}
