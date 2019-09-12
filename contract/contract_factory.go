package contract

import (
	"crypto/sha256"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
)

func GetAddressFromContract(tx *transaction.Transaction) string {
	senderAddress := keytools.GetAddressFromPublic(util.DecodeHex(tx.SenderPubKey))
	sha256Ctx := sha256.New()
	fmt.Printf("send address = %s",senderAddress)
	sha256Ctx.Write(util.DecodeHex(senderAddress))

	var nonce int64

	if tx.Nonce != "" {
		n, _ := strconv.ParseInt(tx.Nonce, 10, 64)
		nonce = n
		nonce = nonce - 1
	}

	hexNonce := util.IntToHex(int(nonce),16)

	sha256Ctx.Write(util.DecodeHex(hexNonce))

	bytes := sha256Ctx.Sum(nil)

	hexString := util.EncodeHex(bytes)

	return hexString[24:]
}
