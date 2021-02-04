package core

import "math/big"

type SWInfo struct {
	ZilliqaMajorVersion uint32
	ZilliqaMinorVersion uint32
	ZilliqaFixVersion   uint32
	ZilliqaUpgradeDS    uint64
	ZilliqaCommit       uint32
	ScillaMajorVersion  uint32
	ScillaMinorVersion  uint32
	ScillaFixVersion    uint32
	ScillaUpgradeDS     uint64
	ScillaCommit        uint32
}

func (sw *SWInfo) Serialize() []byte {
	// length should be 48
	data := make([]byte, 48)
	UintToByteArray(data,0,new(big.Int).SetUint64(uint64(sw.ZilliqaMajorVersion)),4)
	UintToByteArray(data,4,new(big.Int).SetUint64(uint64(sw.ZilliqaMinorVersion)),4)
	UintToByteArray(data,8,new(big.Int).SetUint64(uint64(sw.ZilliqaFixVersion)),4)
	UintToByteArray(data,12,new(big.Int).SetUint64(uint64(sw.ZilliqaUpgradeDS)),8)
	UintToByteArray(data,20,new(big.Int).SetUint64(uint64(sw.ZilliqaCommit)),4)
	UintToByteArray(data,24,new(big.Int).SetUint64(uint64(sw.ScillaMajorVersion)),4)
	UintToByteArray(data,28,new(big.Int).SetUint64(uint64(sw.ScillaMinorVersion)),4)
	UintToByteArray(data,32,new(big.Int).SetUint64(uint64(sw.ScillaFixVersion)),4)
	UintToByteArray(data,36,new(big.Int).SetUint64(uint64(sw.ScillaUpgradeDS)),8)
	UintToByteArray(data,44,new(big.Int).SetUint64(uint64(sw.ScillaCommit)),4)
	return data
}

type SWInfoT struct {
	Scilla  []interface{}
	Zilliqa []interface{}
}
