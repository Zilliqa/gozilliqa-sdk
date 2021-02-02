package core

// https://github.com/Zilliqa/Zilliqa/blob/04162ef0c3c1787ebbd940b7abd6b7ff1d4714ed/src/libData/BlockData/BlockHeader/DSBlockHeader.h
type DsBlock struct {
	DsDifficulty uint8
	Difficulty   uint8
	// The one who proposed this DS block
	// base16 string
	LeaderPubKey string
	// Block index, starting from 0 in the genesis block
	BlockNum uint64
	// Tx Epoch Num then the DS block was generated
	EpochNum     uint64
	GasPrice     string
	PoWDSWinners []string
	SwInfo       SWInfo
	// key is (base16) public key
	RemoveDSNodePubKeys map[string]Peer
	// todo concrete data type
	DSBlockHashSet     interface{}
	GovDSShardVotesMap interface{}
}

type SWInfo struct {
	ZilliqaMajorVersion uint32
	ZilliqaMinorVersion uint32
	ZilliqaFixVersion   uint32
	ZilliqaUpgradeDS    uint32
	ZilliqaCommit       uint32
	ScillaMajorVersion  uint32
	ScillaMinorVersion  uint32
	ScillaFixVersion    uint32
	ScillaUpgradeDS     uint32
	ScillaCommit        uint32
}
