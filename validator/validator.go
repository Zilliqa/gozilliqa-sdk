package validator

import (
	"github.com/FireStack-Lab/LaksaGo"
	"regexp"
	"strconv"
)

func IsBech32(addr string) bool {
	match, _ := regexp.MatchString("^zil1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{38}", addr)
	return match
}
func IsPublicKey(public_key string) bool {
	match, _ := regexp.MatchString("^(0x)?[[:xdigit:]]{66}$", public_key)
	return match
}

func IsPrivateKey(private_key string) bool {
	match, _ := regexp.MatchString("^(0x)?[[:xdigit:]]{64}$", private_key)
	return match
}

func IsAddress(address string) bool {
	match, _ := regexp.MatchString("^(0x)?[[:xdigit:]]{40}$", address)
	return match
}

func IsSignature(signature string) bool {
	match, _ := regexp.MatchString("^(0x)?[[:xdigit:]]{128}$", signature)
	return match
}

// checksum_address?
//
// takes hex-encoded string and returns boolean if address is checksumed
//
// @param {string} address
// @returns {boolean}
func IsChecksumAddress(address string) bool {
	return IsAddress(address) && address == LaksaGo.ToCheckSumAddress(address)
}

func IsByteString(str string, len int) bool {
	pattern := "^(0x)?[0-9a-fA-F]{" + strconv.FormatInt(int64(len), 10) + "}"
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// self.address?(address) && Laksa::Account::Wallet::to_checksum_address(address) == address
// end
