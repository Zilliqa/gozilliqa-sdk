package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIsPublicKey(t *testing.T) {
	public_key := "039E43C9810E6CC09F46AAD38E716DAE3191629534967DC457D3A687D2E2CDDC6A"

	assert.Equal(t, true, IsPublicKey(public_key))
	assert.Equal(t, false, IsPublicKey(public_key[0:64]))
}

func TestIsPrivateKey(t *testing.T) {
	private_key := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"

	assert.Equal(t, true, IsPrivateKey(private_key))
	assert.Equal(t, false, IsPrivateKey(private_key[0:63]))
}

func TestIsAddress(t *testing.T) {
	address := "B5C2CDD79C37209C3CB59E04B7C4062A8F5D5271"

	assert.Equal(t, true, IsAddress(address))
	assert.Equal(t, false, IsAddress(address[0:38]))
}

func TestIsSignature(t *testing.T) {
	signature := "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7"

	assert.Equal(t, true, IsSignature(signature))
	assert.Equal(t, false, IsSignature(signature[0:38]))
}

func TestIsChecksumAddress(t *testing.T) {
	checksum_address := "0x4BAF5faDA8e5Db92C3d3242618c5B47133AE003C"

	assert.Equal(t, true, IsChecksumAddress(checksum_address))
	assert.Equal(t, false, IsChecksumAddress(checksum_address[0:38]))

	fmt.Println(IsChecksumAddress("0xC10ded4923e1F1E0AdfEe276578DEb105A60cA55"))

	fmt.Println(strings.ToLower("0xc10dEd4923e1f1E0ADFEE276578Deb105a60Ca55"))
}

func TestIsBech32(t *testing.T) {
	isBech32 := IsBech32("zil16jrfrs8vfdtc74yzhyy83je4s4c5sqrcasjlc4")
	fmt.Println(isBech32)
}
