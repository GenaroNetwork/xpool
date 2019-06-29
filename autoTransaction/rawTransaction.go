package autoTransaction

import (
"fmt"
"math/big"
"context"
"io/ioutil"

"github.com/GenaroNetwork/GenaroCore/accounts/keystore"
"github.com/GenaroNetwork/GenaroCore/cmd/utils"
"github.com/GenaroNetwork/GenaroCore/node"
"github.com/GenaroNetwork/GenaroCore/params"
"github.com/GenaroNetwork/GenaroCore/common"
"github.com/GenaroNetwork/GenaroCore/core/types"
"strings"
)

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = "eth"
	cfg.Version = params.VersionWithCommit("")
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

func makeConfigNode(ctx context.Context) *node.Node {
	nodeConfig := defaultNodeConfig()
	stack, err := node.New(&nodeConfig)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	return stack
}


func RawTransaction(keyDir,Password string, Value *big.Int,Nonce uint64,ToAccount string) string{
	keyJson, err := ioutil.ReadFile(keyDir)
	if err != nil {
		fmt.Println("read file error: ", err)
	}

	ctx := context.Background()
	stack := makeConfigNode(ctx)

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	acct, err := ks.Import(keyJson, Password, Password)
	if err != nil {
		utils.Fatalf("%v", err)
	}
	var To common.Address = common.HexToAddress(ToAccount)
	var Gas uint64 = 90000
	chain := big.NewInt(200)
	GasPrice := new(big.Int).SetUint64(18000000000)

	tx := types.NewTransaction(Nonce, To, Value, Gas, GasPrice, []byte("information"))

	signTx, err := ks.SignTxWithPassphrase(acct, Password, tx, chain)
	if err != nil {
		fmt.Printf("sign err: %s \n", err)
		return ""
	}else {
		fmt.Printf("%s\n",signTx.String())
	}

	arr := strings.Split(signTx.String(),":")
	Hex := arr[len(arr)-1]
	HexStr := fmt.Sprintf("0x%s", strings.TrimSpace(Hex))
	return HexStr
}


