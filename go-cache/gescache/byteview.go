package gescache

// ByteView a byte view holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice return a copy of data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneByte(v.b)
}

// String return data as a string
func (v ByteView) String() string {
	return string(v.b)
}

func cloneByte(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
