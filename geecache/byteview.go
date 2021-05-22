package geecache

type ByteView struct {
	b []byte // b 将会存储真实的缓存值
}

func (bv ByteView) Len() int {
	return len(bv.b)
}

func (bv ByteView) ByteSlice() []byte {
	return copyByte(bv.b)
}

func (bv ByteView) String() string {
	return string(bv.b)
}

func copyByte(b []byte) []byte {
	c :=  make([]byte, len(b))
	copy(c, b)
	return c
}



