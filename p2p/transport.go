package p2p

// Peer represents the remote node
type Peer interface {
	// ??
}

// Transport is anythig that handles communication
// between nodes in network
type Transport interface {
	ListenAndAccept() error
}
