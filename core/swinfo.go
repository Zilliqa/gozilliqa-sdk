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
	bns := BIGNumSerialize{}
	// length should be 48
	data := make([]byte, 0)
	data = bns.SetNumber(data, 0, 4, new(big.Int).SetUint64(uint64(sw.ZilliqaMajorVersion)))
	data = bns.SetNumber(data, 4, 4, new(big.Int).SetUint64(uint64(sw.ZilliqaMinorVersion)))
	data = bns.SetNumber(data, 8, 4, new(big.Int).SetUint64(uint64(sw.ZilliqaFixVersion)))
	data = bns.SetNumber(data, 12, 8, new(big.Int).SetUint64(sw.ZilliqaUpgradeDS))
	data = bns.SetNumber(data, 20, 4, new(big.Int).SetUint64(uint64(sw.ZilliqaCommit)))
	data = bns.SetNumber(data, 24, 4, new(big.Int).SetUint64(uint64(sw.ScillaMajorVersion)))
	data = bns.SetNumber(data, 28, 4, new(big.Int).SetUint64(uint64(sw.ScillaMinorVersion)))
	data = bns.SetNumber(data, 32, 4, new(big.Int).SetUint64(uint64(sw.ScillaFixVersion)))
	data = bns.SetNumber(data, 36, 8, new(big.Int).SetUint64(sw.ScillaUpgradeDS))
	data = bns.SetNumber(data, 44, 4, new(big.Int).SetUint64(uint64(sw.ScillaCommit)))
	return data
}

type SWInfoT struct {
	Scilla  []interface{}
	Zilliqa []interface{}
}
