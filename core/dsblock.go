package core

import (
	"github.com/Zilliqa/gozilliqa-sdk/protobuf"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"math/big"
	"net"
	"strconv"
)

type DsBlock struct {
	BlockHeader DsBlockHeader
}

//func (ds *DsBlock) ToProtobuf() []byte {
//	protoDSBlockHeader := ds.BlockHeader.ToProtobuf(false)
//
//}

// https://github.com/Zilliqa/Zilliqa/blob/04162ef0c3c1787ebbd940b7abd6b7ff1d4714ed/src/libData/BlockData/BlockHeader/DSBlockHeader.h
type DsBlockHeader struct {
	blockHeaderBase BlockHeaderBase
	DsDifficulty    uint32
	Difficulty      uint32
	// The one who proposed this DS block
	// base16 string
	LeaderPubKey string
	// Block index, starting from 0 in the genesis block
	BlockNum uint64
	// Tx Epoch Num then the DS block was generated
	EpochNum uint64
	GasPrice string
	// key is (base16) public key
	PoWDSWinners map[string]Peer
	SwInfo       SWInfo
	// (base16) public key
	RemoveDSNodePubKeys []string
	// todo concrete data type
	DSBlockHashSet     interface{}
	GovDSShardVotesMap map[uint32]Pair
}

type Pair struct {
	first  map[uint32]uint32
	second map[uint32]uint32
}

func NewFromDsBlockT(dst *DsBlockT) *DsBlockHeader {
	dsBlockHeader := &DsBlockHeader{}
	dsBlockHeader.DsDifficulty = dst.Header.DifficultyDS
	dsBlockHeader.DsDifficulty = dst.Header.Difficulty
	dsBlockHeader.LeaderPubKey = dst.Header.LeaderPubKey

	blockNum, _ := strconv.ParseUint(dst.Header.BlockNum, 10, 64)
	dsBlockHeader.BlockNum = blockNum

	epochNum, _ := strconv.ParseUint(dst.Header.EpochNum, 10, 64)
	dsBlockHeader.EpochNum = epochNum

	dsBlockHeader.GasPrice = dst.Header.GasPrice

	// todo skip swinfo

	winnermap := make(map[string]Peer, len(dst.Header.PoWWinners))
	for i := 0; i < len(dst.Header.PoWWinners); i++ {
		ip := dst.Header.PoWWinnersIP[i].IP
		port := dst.Header.PoWWinnersIP[i].Port

		IPAddress := net.ParseIP(ip)

		peer := Peer{
			IpAddress:      new(big.Int).SetBytes(IPAddress),
			ListenPortHost: port,
		}
		winnermap[dst.Header.PoWWinners[i]] = peer
	}

	dsBlockHeader.PoWDSWinners = winnermap

	var removeDSNodePubKeys []string
	for _, key := range dst.Header.MembersEjected {
		removeDSNodePubKeys = append(removeDSNodePubKeys, key)
	}
	dsBlockHeader.RemoveDSNodePubKeys = removeDSNodePubKeys

	// skip hashset
	// todo skip Governance

	dsBlockHeader.blockHeaderBase.Version = dst.Header.Version
	ch := util.DecodeHex(dst.Header.CommitteeHash)
	var commitHash [32]byte
	copy(commitHash[:], ch)
	dsBlockHeader.blockHeaderBase.CommitteeHash = commitHash

	ph := util.DecodeHex(dst.Header.PrevHash)
	var prevHash [32]byte
	copy(prevHash[:], ph)
	dsBlockHeader.blockHeaderBase.PrevHash = prevHash

	return dsBlockHeader
}

// the default value of concreteVarsOnly should be false
func (d *DsBlockHeader) ToProtobuf(concreteVarsOnly bool) *protobuf.ProtoDSBlock_DSBlockHeader {
	protoDSBlockHeader := &protobuf.ProtoDSBlock_DSBlockHeader{}
	protoBlockHeaderBase := d.blockHeaderBase.ToProtobuf()
	protoDSBlockHeader.Blockheaderbase = protoBlockHeaderBase

	if !concreteVarsOnly {
		protoDSBlockHeader.Dsdifficulty = d.DsDifficulty
		protoDSBlockHeader.Difficulty = d.Difficulty
		data := make([]byte, 16)
		gasPriceInt, _ := new(big.Int).SetString(d.GasPrice, 16)
		data = UintToByteArray(data, 0, gasPriceInt, 16)
		protoDSBlockHeader.Gasprice = &protobuf.ByteArray{
			Data: data,
		}

		var protobufWinners []*protobuf.ProtoDSBlock_DSBlockHeader_PowDSWinners
		for key, winner := range d.PoWDSWinners {
			protobufWinner := &protobuf.ProtoDSBlock_DSBlockHeader_PowDSWinners{
				Key: &protobuf.ByteArray{Data: util.DecodeHex(key)},
				Val: &protobuf.ByteArray{Data: winner.Serialize()},
			}
			protobufWinners = append(protobufWinners, protobufWinner)
		}
		protoDSBlockHeader.Dswinners = protobufWinners

		var proposals []*protobuf.ProtoDSBlock_DSBlockHeader_Proposal
		for proposal, pair := range d.GovDSShardVotesMap {
			protoproposal := &protobuf.ProtoDSBlock_DSBlockHeader_Proposal{}
			protoproposal.Proposalid = proposal

			var dsvotes []*protobuf.ProtoDSBlock_DSBlockHeader_Vote
			for value, count := range pair.first {
				dsvote := &protobuf.ProtoDSBlock_DSBlockHeader_Vote{
					Value: value,
					Count: count,
				}
				dsvotes = append(dsvotes, dsvote)
			}

			var minerVotes []*protobuf.ProtoDSBlock_DSBlockHeader_Vote
			for value, count := range pair.second {
				minerVote := &protobuf.ProtoDSBlock_DSBlockHeader_Vote{
					Value: value,
					Count: count,
				}
				minerVotes = append(minerVotes, minerVote)
			}

			proposals = append(proposals, protoproposal)
		}
		protoDSBlockHeader.Proposals = proposals

		var dsremoved []*protobuf.ByteArray
		for _, key := range d.RemoveDSNodePubKeys {
			dr := &protobuf.ByteArray{
				Data: util.DecodeHex(key),
			}
			dsremoved = append(dsremoved, dr)
		}
		protoDSBlockHeader.Dsremoved = dsremoved
	}

	protoDSBlockHeader.Leaderpubkey = &protobuf.ByteArray{Data: util.DecodeHex(d.LeaderPubKey)}
	protoDSBlockHeader.Oneof6 = &protobuf.ProtoDSBlock_DSBlockHeader_Blocknum{
		Blocknum: d.BlockNum,
	}
	protoDSBlockHeader.Oneof7 = &protobuf.ProtoDSBlock_DSBlockHeader_Epochnum{
		Epochnum: d.EpochNum,
	}

	return protoDSBlockHeader
}

// ds block transfer struct (via rpc)
type DsBlockT struct {
	B1         []bool         `json:"B1"`
	B2         []bool         `json:"B2"`
	CS1        string         `json:"CS1"`
	Header     DsBlockHeaderT `json:"header"`
	Serialized SerializedT    `json:"serialized"`
	Signatures string         `json:"signatures"`
}

type DsBlockHeaderT struct {
	BlockNum       string
	CommitteeHash  string
	Difficulty     uint32
	DifficultyDS   uint32
	EpochNum       string
	GasPrice       string
	LeaderPubKey   string
	MembersEjected []string
	PoWWinners     []string
	PoWWinnersIP   []IPAndPort
	PrevHash       string
	ReservedField  string
	SWInfo         SWInfoT
	ShardingHash   string
	Timestamp      string
	Version        uint32
}

type IPAndPort struct {
	IP   string `json:"IP"`
	Port uint32 `json:"port"`
}

type SerializedT struct {
	Data   string `json:"data"`
	Header string `json:"header"`
}
