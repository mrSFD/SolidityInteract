package main
import (
	"fmt"
	"log"
	"context"
	"math/big"
	"crypto/ecdsa"
	
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"Todo"
)
var (
	URL = "wss://rinkeby.infura.io/ws/v3/14a9f178d7564fecac3afa7e62d095ff"
)
func main() {
	//------------------------------------------------------
	PrivateKey, err := crypto.HexToECDSA("381c1e5cbe5a6dc7e34e58f65dc55bf957859b7331296c9d6d679b58666aae37")
    if err != nil {
        log.Fatal(err)
    }
	//------------------------------------------------------
   	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------	
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------
	// contract address	
	cAdd := common.HexToAddress("0x097333Cf08A5b21fb5Ff1d98250E162872378B93")
	//------------------------------------------------------
	t, err := Todo.NewTodo(cAdd,client)	
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------------
	tx, err := bind.NewKeyedTransactorWithChainID(PrivateKey,chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = uint64(3000000)
	tx.GasPrice = gasPrice

	//----------- ADD Message ------------------------------
	// name, message, id
	// comment: "crypto/ecdsa" and "github.com/ethereum/go-ethereum/crypto"
//	trs, err := t.Add(tx,"test1", "ccc",big.NewInt(0))
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(trs.Hash())
	
	//------------ GET Message -----------------------------
	// using: "crypto/ecdsa" and "github.com/ethereum/go-ethereum/crypto"

	publicKey := PrivateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		log.Fatal("error casting public key to ECDSA")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	tasks, err := t.Get(&bind.CallOpts{From:fromAddress,},big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
	
	//------------ LIST Message ----------------------------
	// using: "crypto/ecdsa" and "github.com/ethereum/go-ethereum/crypto"

//	publicKey := PrivateKey.Public()
//    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//    if !ok {
//  	    log.Fatal("error casting public key to ECDSA")
//    }
//    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

//	tasks, err := t.List(&bind.CallOpts{From:fromAddress,})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(tasks)
}