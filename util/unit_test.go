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
package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertQaToZil(t *testing.T) {
	qa := 1000000000000.0
	ret := FromQa(qa, ZIL, false)
	assert.Equal(t, 1.0, ret)
}

func TestConvertQaToLi(t *testing.T) {
	qa := 1000000.0
	ret := FromQa(qa, LI, false)
	assert.Equal(t, 1.0, ret)
}

func TestConvertLiToQa(t *testing.T) {
	val := 1.0
	ret := ToQa(val, LI)
	assert.Equal(t, 1000000.0, ret)
}

func TestConvertZilToQa(t *testing.T) {
	val := 1.0
	ret := ToQa(val, ZIL)
	assert.Equal(t, 1000000000000.0, ret)
}

func TestFromQaNegativeNumber(t *testing.T) {
	val := -1000000000000.0
	ret := FromQa(val, ZIL, false)
	assert.Equal(t, -1.0, ret)
}

func TestFromQaWithPack(t *testing.T) {
	val := 1000000000001.0
	ret := FromQa(val, ZIL, true)
	assert.Equal(t, 1.0, ret)
}

func TestToQaNegativeNumber(t *testing.T) {
	val := -1.0
	ret := ToQa(val, ZIL)
	assert.Equal(t, -1000000000000.0, ret)
}
