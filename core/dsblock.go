package core

type DsBlock struct {
	BlockBase
	BlockHeader DsBlockHeader
}

//func (ds *DsBlock) ToProtobuf() []byte {
//	protoDSBlockHeader := ds.BlockHeader.ToProtobuf(false)
//
//}
