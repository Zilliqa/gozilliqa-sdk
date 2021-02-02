package core

type Peer struct {
	// Peer IP address (net-encoded)
	IpAddress uint64
	// Peer listen port (host-encoded)
	ListenPortHost uint32
	HostName       string
}
