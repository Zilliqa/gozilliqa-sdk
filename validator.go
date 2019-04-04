package LaksaGo

import (
	"regexp"
)

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
	return IsAddress(address)
}

// self.address?(address) && Laksa::Account::Wallet::to_checksum_address(address) == address
// end
