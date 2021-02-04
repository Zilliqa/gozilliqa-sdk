package core

import (
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"testing"
)

func TestSWInfo_Serialize(t *testing.T) {
	swinfo := &SWInfo{
		ZilliqaMajorVersion: 7,
		ZilliqaMinorVersion: 1,
		ZilliqaFixVersion:   1,
		ZilliqaUpgradeDS:    0,
		ZilliqaCommit:       0,
		ScillaMajorVersion:  3,
		ScillaMinorVersion:  2,
		ScillaFixVersion:    0,
		ScillaUpgradeDS:     0,
		ScillaCommit:        1,
	}

	data := swinfo.Serialize()
	t.Log(util.EncodeHex(data))
	// 000000070000000100000001000000000000000000000000000000030000000200000000000000000000000000000001
	// todo assertion
}
