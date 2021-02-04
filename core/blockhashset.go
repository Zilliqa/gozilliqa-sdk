package core

// Hashes for DSBlockHashSet
type SharingHash [32]byte
type TxSharingHash [32]byte

type DSBlockHashSet struct {
	shadingHash   SharingHash
	reservedField []byte
}
