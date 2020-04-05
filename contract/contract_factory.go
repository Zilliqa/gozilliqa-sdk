/*
 * Copyright (C) 2019 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
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
	fmt.Printf("send address = %s", senderAddress)
	sha256Ctx.Write(util.DecodeHex(senderAddress))

	var nonce int64

	if tx.Nonce != "" {
		n, _ := strconv.ParseInt(tx.Nonce, 10, 64)
		nonce = n
		nonce = nonce - 1
	}

	hexNonce := util.IntToHex(int(nonce), 16)

	sha256Ctx.Write(util.DecodeHex(hexNonce))

	bytes := sha256Ctx.Sum(nil)

	hexString := util.EncodeHex(bytes)

	return hexString[24:]
}
