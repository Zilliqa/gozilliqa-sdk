package contract

import (
	"crypto/sha256"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"strconv"
)

func GetAddressFromContract(tx *transaction.Transaction) string {
	senderAddress := keytools.GetAddressFromPublic(LaksaGo.DecodeHex(tx.SenderPubKey))
	sha256Ctx := sha256.New()
	fmt.Printf("send address = %s",senderAddress)
	sha256Ctx.Write(LaksaGo.DecodeHex(senderAddress))

	var nonce int64

	if tx.Nonce != "" {
		n, _ := strconv.ParseInt(tx.Nonce, 10, 64)
		nonce = n
		nonce = nonce - 1
	}

	hexNonce := LaksaGo.IntToHex(int(nonce),16)

	sha256Ctx.Write(LaksaGo.DecodeHex(hexNonce))

	bytes := sha256Ctx.Sum(nil)

	hexString := LaksaGo.EncodeHex(bytes)

	return hexString[24:]
}
