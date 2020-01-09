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
package subscription

import "encoding/json"

type Topic interface {
	Stringify() ([]byte, error)
}

type NewBlockQuery struct {
	Query string `json:"query"`
}

func (query *NewBlockQuery) Stringify() ([]byte, error) {
	return json.Marshal(query)
}

type NewEventLogQuery struct {
	Query     string   `json:"query"`
	Addresses []string `json:"addresses"`
}

func (query *NewEventLogQuery) Stringify() ([]byte, error) {
	return json.Marshal(query)
}

type Unsubscribe struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

func (query *Unsubscribe) Stringify() ([]byte, error) {
	return json.Marshal(query)
}
