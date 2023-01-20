/*
 * Copyright (C) 2021 Zilliqa
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
package mpt

import (
	"errors"
	"github.com/Zilliqa/gozilliqa-sdk/v3/util"
	"github.com/ethereum/go-ethereum/crypto"
	"sync"
)

type Database struct {
	db   map[string][]byte
	lock sync.RWMutex
}

func NewFromProof(proof [][]byte) *Database {
	db := &Database{}
	db.db = make(map[string][]byte, len(proof))
	for _, v := range proof {
		data := v
		key := util.EncodeHex(crypto.Keccak256(data))
		db.db[key] = data
	}
	return db
}

func (db *Database) Get(key []byte) ([]byte, error) {
	keystr := util.EncodeHex(key)
	if v, ok := db.db[keystr]; ok {
		return v, nil
	}
	return nil, errors.New("cannot find")
}
