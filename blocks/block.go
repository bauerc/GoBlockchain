package blocks

//Block defines the functionality needed to preform Blockchain operations
type Block interface {
	CaculateHash() string
	GenerateBlock(i interface{}) (Block, error)
	IsBlockValid(newBlock Block) bool
}
