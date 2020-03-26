package gescache

import (
	pb "github.com/simple-framework-golang/go-cache/gescache/gescachepb"
)

// PeerPicker is the interface that must be implemented to locate
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented to peer
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
