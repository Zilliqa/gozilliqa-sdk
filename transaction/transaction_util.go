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
package transaction

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/protobuf"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/golang/protobuf/proto"
	"math/big"
	"strconv"
)

func EncodeTransactionProto(txParams TxParams) ([]byte, error) {
	amount, ok := new(big.Int).SetString(txParams.Amount, 10)
	if !ok {
		return nil, errors.New("amount error")
	}

	gasPrice, ok2 := new(big.Int).SetString(txParams.GasPrice, 10)
	if !ok2 {
		return nil, errors.New("gas price error")
	}

	v, err := strconv.ParseUint(txParams.Version, 10, 32)
	if err != nil {
		return nil, err
	}
	version := uint32(v)

	nonce, err2 := strconv.ParseUint(txParams.Nonce, 10, 64)
	if err2 != nil {
		return nil, err2
	}

	senderpubkey := protobuf.ByteArray{
		Data: util.DecodeHex(txParams.SenderPubKey),
	}

	amountArray := protobuf.ByteArray{
		Data: bigIntToPaddedBytes(amount, 32),
	}

	gasPriceArray := protobuf.ByteArray{
		Data: bigIntToPaddedBytes(gasPrice, 32),
	}

	gasLimit, err3 := strconv.ParseUint(txParams.GasLimit, 10, 64)
	if err3 != nil {
		return nil, err3
	}

	protoTransactionCoreInfo := protobuf.ProtoTransactionCoreInfo{
		Version:      &version,
		Nonce:        &nonce,
		Toaddr:       util.DecodeHex(txParams.ToAddr),
		Senderpubkey: &senderpubkey,
		Amount:       &amountArray,
		Gasprice:     &gasPriceArray,
		Gaslimit:     &gasLimit,
	}

	if txParams.Data == "\"\"" {
		txParams.Data = ""
	}

	if txParams.Data != "" {
		protoTransactionCoreInfo.Data = []byte(txParams.Data)
	}

	if txParams.Code != "" {
		protoTransactionCoreInfo.Code = []byte(txParams.Code)
	}

	bytes, err4 := proto.Marshal(&protoTransactionCoreInfo)
	if err4 != nil {
		return nil, err4
	}
	return bytes, nil

}

func bigIntToPaddedBytes(i *big.Int, paddedSize int32) []byte {
	bytes := i.Bytes()
	padded, _ := hex.DecodeString(fmt.Sprintf("%0*x", paddedSize, bytes))
	return padded
}
