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
package bech32

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToBech32Address(t *testing.T) {
	addr, _ := ToBech32Address("1d19918a737306218b5cbb3241fcdcbd998c3a72")
	assert.Equal(t, addr, "zil1r5verznnwvrzrz6uhveyrlxuhkvccwnju4aehf")

	addr, _ = ToBech32Address("cc8ee24773e1b4b28b3cc5596bb9cfc430b48453")
	assert.Equal(t, addr, "zil1ej8wy3mnux6t9zeuc4vkhww0csctfpznzt4s76")

	addr, _ = ToBech32Address("e14576944443e9aeca6f12b454941884aa122938")
	assert.Equal(t, addr, "zil1u9zhd9zyg056ajn0z269f9qcsj4py2fc89ru3d")

	addr, _ = ToBech32Address("179361114cbfd53be4d3451edf8148cde4cfe774")
	assert.Equal(t, addr, "zil1z7fkzy2vhl2nhexng50dlq2gehjvlem5w7kx8z")

	addr, _ = ToBech32Address("5a2b667fdeb6356597681d08f6cd6636aed94784")
	assert.Equal(t, addr, "zil1tg4kvl77kc6kt9mgr5y0dntxx6hdj3uy95ash8")

	addr, _ = ToBech32Address("537342e5e0a6b402f281e2b4301b89123ae31117")
	assert.Equal(t, addr, "zil12de59e0q566q9u5pu26rqxufzgawxyghq0vdk9")

	addr, _ = ToBech32Address("5e61d42a952d2df1f4e5cbed7f7d1294e9744a52")
	assert.Equal(t, addr, "zil1tesag25495klra89e0kh7lgjjn5hgjjj0qmu8l")

	addr, _ = ToBech32Address("5f5db1c18ccde67e513b7f7ae820e569154976ba")
	assert.Equal(t, addr, "zil1tawmrsvvehn8u5fm0aawsg89dy25ja46ndsrhq")
}

func TestFromBech32Addr(t *testing.T) {
	addr, _ := FromBech32Addr("zil1fwh4ltdguhde9s7nysnp33d5wye6uqpugufkz7")
	assert.Equal(t, strings.ToUpper(addr), "4BAF5FADA8E5DB92C3D3242618C5B47133AE003C")

	addr, _ = FromBech32Addr("zil1gjpxry26srx7n008c7nez6zjqrf6p06wur4x3m")
	assert.Equal(t, strings.ToUpper(addr), "448261915A80CDE9BDE7C7A791685200D3A0BF4E")

	addr, _ = FromBech32Addr("zil1mmgzlktelsh9tspy80f02t0sytzq4ks79zdnkk")
	assert.Equal(t, strings.ToUpper(addr), "DED02FD979FC2E55C0243BD2F52DF022C40ADA1E")

	addr, _ = FromBech32Addr("zil1z0cxucpf004x50zq9ahkf3qk56e3ukrwaty4g8")
	assert.Equal(t, strings.ToUpper(addr), "13F06E60297BEA6A3C402F6F64C416A6B31E586E")

	addr, _ = FromBech32Addr("zil1r2gvy5c8c0x8r9v2s0azzw3rvtv9nnenynd33g")
	assert.Equal(t, strings.ToUpper(addr), "1A90C25307C3CC71958A83FA213A2362D859CF33")

	addr, _ = FromBech32Addr("zil1vfdt467c0khf4vfg7we6axtg3qfan3wlf9yc6y")
	assert.Equal(t, strings.ToUpper(addr), "625ABAEBD87DAE9AB128F3B3AE99688813D9C5DF")

	addr, _ = FromBech32Addr("zil1x6argztlscger3yvswwfkx5ttyf0tq703v7fre")
	assert.Equal(t, strings.ToUpper(addr), "36BA34097F861191C48C839C9B1A8B5912F583CF")

	addr, _ = FromBech32Addr("zil16fzn4emvn2r24e2yljnfnk7ut3tk4me6qx08ed")
	assert.Equal(t, strings.ToUpper(addr), "D2453AE76C9A86AAE544FCA699DBDC5C576AEF3A")

	addr, _ = FromBech32Addr("zil1wg3qapy50smprrxmckqy2n065wu33nvh35dn0v")
	assert.Equal(t, strings.ToUpper(addr), "72220E84947C36118CDBC580454DFAA3B918CD97")

	addr, _ = FromBech32Addr("zil12rujxpxgjtv55wzu5m8xe454pn56x6pedpl554")
	assert.Equal(t, strings.ToUpper(addr), "50F92304C892D94A385CA6CE6CD6950CE9A36839")

	addr, _ = FromBech32Addr("zil1r5verznnwvrzrz6uhveyrlxuhkvccwnju4aehf")
	assert.Equal(t, strings.ToLower(addr), "1d19918a737306218b5cbb3241fcdcbd998c3a72")

	addr, _ = FromBech32Addr("zil1tawmrsvvehn8u5fm0aawsg89dy25ja46ndsrhq")
	assert.Equal(t, strings.ToLower(addr), "5f5db1c18ccde67e513b7f7ae820e569154976ba")
}

func TestDecode(t *testing.T) {
	bech32, err := ToBech32Address("0x2ce491a0fd9e318b39172258101b7c836da7449b")
	assert.Nil(t, err, err)
	assert.Equal(t, "zil19njfrg8anccckwghyfvpqxmusdk6w3ymwmdg6g", bech32)
}
