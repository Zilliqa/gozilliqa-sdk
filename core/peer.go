package core

import "math/big"

type Peer struct {
	// Peer IP address (net-encoded)
	// size is 128bits (16 bytes)
	IpAddress *big.Int
	// Peer listen port (host-encoded)
	ListenPortHost uint32
	HostName       string
}

func (p *Peer) Serialize() []byte {
	bns := BIGNumSerialize{}
	data := make([]byte, 0)
	port := new(big.Int).SetUint64(uint64(p.ListenPortHost))
	data = bns.SetNumber(data, 0, 16, p.IpAddress)
	data = bns.SetNumber(data, 16, 4, port)
	return data
}
