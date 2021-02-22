package verifier

import (
	"container/list"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"os"
	"testing"
)

// todo change this test for community testnet or mainnet in the future
func TestVerify(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	p := provider.NewProvider("https://kaus-exposed-apis6-api.dev.z7a.xyz/")
	verifier := &Verifier{RpcClient: p, NumOfDsGuard: 9}
	dsComm := list.New()

	// 9 guard nodes
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0213D5A7F74B28F3F588FF6520748DBB541986E98F75FA78D6334B2D0AAB4C1E57",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0239D4CAE39A7AC2F285796BABF7D28DC8EB7767E78409C70926D0929EA2941E36",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02D2D695D4A352412E0D32A8BDF6EA3A606D35FE2C2F850C54D68727D065894986",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "02E5E1BE6C924349F2C2B20CE05A2650B3E56C7722A2E5952EE27D12DEE7A4A6E6",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0300AB86B413FAA64A52FB61B5A28A6C361F87A5B0871C4F01C394D261415B0989",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "03019AF5B10FFE09FB0EE02B59195EF5E6F5BE51D17EAF5604EA452078CD465C4B",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0323086D473DF937B6297FB755FA8E57C0FB2760512AED7757748B597C48F797A0",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "032AEE20CFC59EAEB7838DAC2A9BAF96C8D69CF2C866FB4A3F1DFB02BCFCA356BB",
	})
	dsComm.PushBack(core.PairOfNode{
		PubKey: "033207325A3CC671034FEBA86EC8D0AA412DF60C7E8292044D510DF582787DCC05",
	})
	// one normal node
	dsComm.PushBack(core.PairOfNode{
		PubKey: "0334AA0F7CA2EAA56B6B752533F9C60777E96C6D1ABE84B463F60ADD89843794AE",
	})

	dst, _ := p.GetDsBlockVerbose("1")

	dsComm1, err := verifier.VerifyDsBlock(core.NewDsBlockFromDsBlockT(dst), dsComm)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	printDsComm(t, dsComm1)

	t.Log("verify ds block 1 successful")

	txblock1, _ := verifier.RpcClient.GetTxBlockVerbose("1")
	err2 := verifier.VerifyTxBlock(core.NewTxBlockFromTxBlockT(txblock1), dsComm1)
	if err2 != nil {
		t.Error(err2)
		t.FailNow()
	}

	t.Log("verify tx block 1 successful")

	dst, _ = p.GetDsBlockVerbose("2")
	dsComm2, err3 := verifier.VerifyDsBlock(core.NewDsBlockFromDsBlockT(dst), dsComm1)
	if err3 != nil {
		t.Error(err3)
		t.FailNow()
	}

	printDsComm(t, dsComm2)
	t.Log("verify ds block 2 successful")

	dst, _ = p.GetDsBlockVerbose("3")
	dsComm3, err4 := verifier.VerifyDsBlock(core.NewDsBlockFromDsBlockT(dst), dsComm1)
	if err4 != nil {
		t.Error(err4)
		t.FailNow()
	}

	printDsComm(t, dsComm3)
	t.Log("verify ds block 3 successful")

	txblock3, _ := verifier.RpcClient.GetTxBlockVerbose("3")
	err5 := verifier.VerifyTxBlock(core.NewTxBlockFromTxBlockT(txblock3), dsComm3)
	if err5 != nil {
		t.Error(err5)
		t.FailNow()
	}
	t.Log("verify tx block 3 successful")
}

func printDsComm(t *testing.T, dsComm *list.List) {
	cursor := dsComm.Front()
	for cursor != nil {
		t.Log(cursor.Value)
		cursor = cursor.Next()
	}
}
