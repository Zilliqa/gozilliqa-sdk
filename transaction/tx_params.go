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

import "github.com/Zilliqa/gozilliqa-sdk/core"

type TxParams struct {
	ID           string                  `json:"ID"`
	Version      string                  `json:"version"`
	Nonce        string                  `json:"nonce"`
	Amount       string                  `json:"amount"`
	GasPrice     string                  `json:"gasPrice"`
	GasLimit     string                  `json:"gasLimit"`
	Signature    string                  `json:"signature"`
	Receipt      core.TransactionReceipt `json:"receipt"`
	SenderPubKey string                  `json:"senderPubKey"`
	ToAddr       string                  `json:"toAddr"`
	Code         string                  `json:"code"`
	Data         string                  `json:"data"`
}
