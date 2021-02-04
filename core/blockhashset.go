package core

type DSBlockHashSet struct {
	// should be 32 bytes
	shadingHash   []byte
	reservedField [128]byte
}
