package SD

import (
	"github.com/LincolnG4/GoMDF/blocks"
	"github.com/LincolnG4/GoMDF/readeratwrapper"
)

type Block struct {
	Header blocks.Header
	Data   []byte
}

func New(file *readeratwrapper.ReaderAtWrapper, startAdress int64) *Block {
	var b Block
	var err error

	b.Header = blocks.Header{}

	b.Header, err = blocks.GetHeader(file, startAdress, blocks.SdID)
	if err != nil {
		return b.BlankBlock()
	}

	return &b
}

func (b *Block) BlankBlock() *Block {
	return &Block{}
}
