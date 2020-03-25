package gescache

// PeerPicker is the interface that must be implemented to locate
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented to peer
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
