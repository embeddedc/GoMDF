package mf4

import (
	"github.com/LincolnG4/GoMDF/blocks/DG"
	"github.com/LincolnG4/GoMDF/readeratwrapper"
)

type DataGroup struct {
	block           *DG.Block
	ChannelGroup    []*ChannelGroup
	CachedDataGroup []byte
}

func NewDataGroup(file *readeratwrapper.ReaderAtWrapper, address int64) DataGroup {
	dataGroupBlock := DG.New(file, address)
	return DataGroup{
		block:        dataGroupBlock,
		ChannelGroup: []*ChannelGroup{},
	}
}

func (d *DataGroup) DataAddress() int64 {
	return d.block.Link.Data
}
