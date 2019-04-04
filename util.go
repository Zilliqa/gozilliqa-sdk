package LaksaGo

import (
	"encoding/hex"
)

func Pack(a int, b int) int {
	return a<<16 + b
}

func EncodeHex(src []byte) string {
	return hex.EncodeToString(src)
}

func DecodeHex(src string) []byte {
	ret, _ := hex.DecodeString(src)
	return ret
}
