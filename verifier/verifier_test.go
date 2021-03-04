package verifier

import (
	"container/list"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

// todo change this test for community testnet or mainnet in the future
func TestVerify(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	p := provider.NewProvider("https://kaus-poly-merged2-api.dev.z7a.xyz")
	verifier := &Verifier{NumOfDsGuard: 9}
	dsComm := list.New()

	// 9 guard nodes
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02105342331FCD7CA95648DF8C5373C596982544F35E90849B1E619DFC59F03D48",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "021D439D1CCCAE17C3D6E855BC78E96438C808D16D1CBF8D7ABD391E41CEE9B1BF",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "021EDDE95598F5F59708D2E728E00EDB2ECF278C16BD389384320B1AF998DCC2FD",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02445FE498E7FBB240BDF9185EB5E7642AF1AF36852D1E132E198A222FBAC617A0",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0256EC4BC62FB56C83A3F6160E67499A9E381CF7A613EBF34B9ECDB9E64171DDF4",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0264D991762D81DD6557BCB33EC8AA3F621B4CB790852F2231C864921387B76862",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "027A00916BDD3CF954ED13A0494BFB73FF95BF28C54004F2749F1A8E8CC1AB5B3D",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0297C693FBEBAF397CBDE616F605920EF70D7F6E5EC8DD82E71AE1E812E5E0B303",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02AE5ADF63E9161000713987B5EBB490B5E6B57CF5B7F9799B4AB907BA19D468F6",
	})
	// one normal node
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02D3CB3FFC8DDE2A55AC29D013CEB5636806C6FC61C5AF077B6313DC636027A602",
	})


	dst, _ := p.GetDsBlockVerbose("1")
	dsBlock := core.NewDsBlockFromDsBlockT(dst)

	dsComm1, err := verifier.VerifyDsBlock(dsBlock, dsComm)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	printDsComm(t, dsComm1)

	t.Log("verify ds block 1 successful")

	txblock1, _ := p.GetTxBlockVerbose("1")
	err2 := verifier.VerifyTxBlock(core.NewTxBlockFromTxBlockT(txblock1), dsComm1)
	if err2 != nil {
		t.Error(err2)
		t.FailNow()
	}

	t.Log("verify tx block 1 successful")


	currentDsBlockNum := uint64(1)
	currentTxBlockNum := uint64(1)
	preDsBlockHash := util.EncodeHex(dsBlock.BlockHash[:])

	for {
		latestTxBlock,_ := p.GetLatestTxBlock()
		latest,_ := strconv.ParseUint(latestTxBlock.Header.BlockNum,10,64)
		if latest > currentTxBlockNum {
			currentTxBlockNum++
			// before handle tx block, check ds block first
			txblockT,_ := p.GetTxBlockVerbose(strconv.FormatUint(currentTxBlockNum,10))
			dsBlockNum,_ := strconv.ParseUint(txblockT.Header.DSBlockNum,10,64)
			if dsBlockNum > currentDsBlockNum {
				currentDsBlockNum++
				dsBlockT,_ := p.GetDsBlockVerbose(strconv.FormatUint(dsBlockNum,10))
				dsBlock := core.NewDsBlockFromDsBlockT(dsBlockT)
				if strings.ToUpper(dsBlock.PrevDSHash) != strings.ToUpper(preDsBlockHash) {
					fmt.Println(dsBlock.PrevDSHash)
					fmt.Println(preDsBlockHash)
					t.Logf("verify ds block %d failed, pre hash wrong\n", dsBlockNum)
					t.FailNow()
				}
				preDsBlockHash = util.EncodeHex(dsBlock.BlockHash[:])
				newDsComm, err := verifier.VerifyDsBlock(dsBlock, dsComm)
				if err == nil {
					t.Logf("verify ds block %d succeed\n", dsBlockNum)
				} else {
					t.Logf("verify ds block %d failed\n", dsBlockNum)
					t.FailNow()
				}
				dsComm = newDsComm
			}

			err := verifier.VerifyTxBlock(core.NewTxBlockFromTxBlockT(txblockT), dsComm)
			if err == nil {
				t.Logf("verify tx block %d succeed\n", currentTxBlockNum)
			} else {
				t.Logf("verify tx block %d failed\n", currentTxBlockNum)
				t.FailNow()
			}
		} else {
			time.Sleep(time.Second)
		}
	}

}

func printDsComm(t *testing.T, dsComm *list.List) {
	cursor := dsComm.Front()
	for cursor != nil {
		t.Log(cursor.Value)
		cursor = cursor.Next()
	}
}
