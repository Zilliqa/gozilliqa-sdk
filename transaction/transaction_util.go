package transaction

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/protobuf"
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
		Data: LaksaGo.DecodeHex(txParams.SenderPubKey),
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
		Toaddr:       LaksaGo.DecodeHex(txParams.ToAddr),
		Senderpubkey: &senderpubkey,
		Amount:       &amountArray,
		Gasprice:     &gasPriceArray,
		Gaslimit:     &gasLimit,
	}

	if txParams.Data != "" {
		protoTransactionCoreInfo.Data = LaksaGo.DecodeHex(txParams.Data)
	}

	if txParams.Code != "" {
		protoTransactionCoreInfo.Data = LaksaGo.DecodeHex(txParams.Code)
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
