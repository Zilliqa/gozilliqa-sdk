/*
 * Copyright (C) 2021 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package core

const (
	signatureChallengeSize = 32
)

type Pair struct {
	FirstList  []uint32
	First      map[uint32]uint32
	SecondList []uint32
	Second     map[uint32]uint32
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
	SWInfo         *SWInfoT
	ShardingHash   string
	Governance     []GovernanceElementT
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

// ds block transfer struct (via rpc)
type DsBlockT struct {
	B1         []bool         `json:"B1"`
	B2         []bool         `json:"B2"`
	CS1        string         `json:"CS1"`
	PrevDSHash string         `json:"PrevDSHash"`
	Header     DsBlockHeaderT `json:"header"`
	Serialized SerializedT    `json:"serialized"`
	Signatures string         `json:"signature"`
}

type VoteT struct {
	VoteCount uint32
	VoteValue uint32
}

type GovernanceElementT struct {
	DSVotes    []VoteT
	ShardVotes []VoteT
	ProposalId uint32
}

type PairOfNode struct {
	PubKey string
	Peer
}
