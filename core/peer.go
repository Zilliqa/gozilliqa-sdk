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
	data := make([]byte, 0)
	port := new(big.Int).SetUint64(uint64(p.ListenPortHost))
	data = UintToByteArray(data,0,p.IpAddress,16)
	data = UintToByteArray(data,16,port,4)
	return data
}
