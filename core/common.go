package core

const (
	signatureChallengeSize = 32
)

type Pair struct {
	First  map[uint32]uint32
	Second map[uint32]uint32
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
