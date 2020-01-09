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
package provider

type TransactionPayload struct {
	Version   int    `json:"version"`
	Nonce     int    `json:"nonce"`
	ToAddr    string `json:"toAddr"`
	Amount    string `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  string `json:"gasPrice"`
	GasLimit  string `json:"gasLimit"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
	Priority  bool   `json:"priority"`
}
