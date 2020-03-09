package test

import (
	"fmt"
	"github.com/truechain/truechain-engineering-code/crypto"
	"math/big"
	"testing"

	"github.com/truechain/truechain-engineering-code/core"
	"github.com/truechain/truechain-engineering-code/core/state"
	"github.com/truechain/truechain-engineering-code/core/types"
)

///////////////////////////////////////////////////////////////////////
func TestOnlyDeposit(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(number uint64, gen *core.BlockGen, fastChain *core.BlockChain, header *types.Header, statedb *state.StateDB) {
		sendTranction(number, gen, statedb, mAccount, saddr1, big.NewInt(6000000000000000000), priKey, signer, nil, header)

		sendDepositTransaction(number, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-61, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight, gen, saddr1, big.NewInt(3000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.GetEpochFromID(2).BeginHeight-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2), gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.MinCalcRedeemHeight(2)-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)

	}
	skey, _ := crypto.HexToECDSA("c6c559a2791634e48e001f2376b61702d6a0d7be04a8ef179e9e066976f5091d")
	saddr := crypto.PubkeyToAddress(skey.PublicKey)
	skey2, _ := crypto.HexToECDSA("aea5a4adeaef3ad87891e68884d961c2a4daeb8659235b94a1a1daa5c5dab233")
	saddr2 := crypto.PubkeyToAddress(skey2.PublicKey)
	skey3, _ := crypto.HexToECDSA("7aa55374ab8e81516b1f00e02f8a8a58b99e98de95f776710979aa931a676bc6")
	saddr3 := crypto.PubkeyToAddress(skey3.PublicKey)
	skey26, _ := crypto.HexToECDSA("6d9e8bd95ce048ce1b778b5d03967982f070b59fcb8cb494f8d0757b798aaf6b")
	saddr26 := crypto.PubkeyToAddress(skey26.PublicKey)

	fmt.Println("saddr", saddr.String(), "saddr2", saddr2.String(), "saddr3", saddr3.String(), "saddr26 ", saddr26.String())

	manager := newTestPOSManager(101, executable)
	fmt.Println(" saddr1 ", manager.GetBalance(saddr1), " StakingAddress ", manager.GetBalance(types.StakingAddress), " ", types.ToTrue(manager.GetBalance(types.StakingAddress)))
	fmt.Println("epoch ", types.GetEpochFromID(1), " ", types.GetEpochFromID(2), " ", types.GetEpochFromID(3), " ", types.GetEpochFromID(4), " ", types.GetEpochFromID(5))
	fmt.Println("epoch ", types.GetEpochFromID(2), " ", types.MinCalcRedeemHeight(2))
	//epoch  [id:1,begin:1,end:2000]   [id:2,begin:2001,end:4000]   [id:3,begin:4001,end:6000]
	//epoch  [id:2,begin:2001,end:4000]   5002
}

func TestCancelMoreDeposit(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(number uint64, gen *core.BlockGen, fastChain *core.BlockChain, header *types.Header, statedb *state.StateDB) {
		sendTranction(number, gen, statedb, mAccount, saddr1, big.NewInt(6000000000000000000), priKey, signer, nil, header)

		sendDepositTransaction(number, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight, gen, saddr1, big.NewInt(2000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight-60, gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight-120, gen, saddr1, big.NewInt(3000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2), gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
	}
	manager := newTestPOSManager(101, executable)
	fmt.Println(" saddr1 ", manager.GetBalance(saddr1), " StakingAddress ", manager.GetBalance(types.StakingAddress), " ", types.ToTrue(manager.GetBalance(types.StakingAddress)))
	fmt.Println("epoch ", types.GetEpochFromID(1), " ", types.GetEpochFromID(2), " ", types.GetEpochFromID(3))
	fmt.Println("epoch ", types.GetEpochFromID(2), " ", types.MinCalcRedeemHeight(2))
}

func TestWithdrawMoreDeposit(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(number uint64, gen *core.BlockGen, fastChain *core.BlockChain, header *types.Header, statedb *state.StateDB) {
		sendTranction(number, gen, statedb, mAccount, saddr1, big.NewInt(6000000000000000000), priKey, signer, nil, header)

		sendDepositTransaction(number, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight, gen, saddr1, big.NewInt(3000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2), gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2)-10, gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2)-20, gen, saddr1, big.NewInt(2000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
	}
	manager := newTestPOSManager(101, executable)
	fmt.Println(" saddr1 ", manager.GetBalance(saddr1), " StakingAddress ", manager.GetBalance(types.StakingAddress), " ", types.ToTrue(manager.GetBalance(types.StakingAddress)))
	fmt.Println("epoch ", types.GetEpochFromID(1), " ", types.GetEpochFromID(2), " ", types.GetEpochFromID(3))
	fmt.Println("epoch ", types.GetEpochFromID(2), " ", types.MinCalcRedeemHeight(2))
}

func TestWithdrawAll(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(number uint64, gen *core.BlockGen, fastChain *core.BlockChain, header *types.Header, statedb *state.StateDB) {
		sendTranction(number, gen, statedb, mAccount, saddr1, big.NewInt(6000000000000000000), priKey, signer, nil, header)

		sendDepositTransaction(number, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-61, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.GetEpochFromID(2).BeginHeight-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2), gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.MinCalcRedeemHeight(2)-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)

	}

	manager := newTestPOSManager(101, executable)
	fmt.Println(" saddr1 ", manager.GetBalance(saddr1), " StakingAddress ", manager.GetBalance(types.StakingAddress), " ", types.ToTrue(manager.GetBalance(types.StakingAddress)))
	fmt.Println("epoch ", types.GetEpochFromID(1), " ", types.GetEpochFromID(2), " ", types.GetEpochFromID(3), " ", types.GetEpochFromID(4), " ", types.GetEpochFromID(5))
	fmt.Println("epoch ", types.GetEpochFromID(2), " ", types.MinCalcRedeemHeight(2))
}

///////////////////////////////////////////////////////////////////////
func TestDepositAppend(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(number uint64, gen *core.BlockGen, fastChain *core.BlockChain, header *types.Header, statedb *state.StateDB) {
		sendTranction(number, gen, statedb, mAccount, saddr1, big.NewInt(6000000000000000000), priKey, signer, nil, header)

		sendDepositTransaction(number, gen, saddr1, big.NewInt(4000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-31, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendDepositAppendTransaction(number, gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-41, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendCancelTransaction(number-types.GetEpochFromID(2).BeginHeight, gen, saddr1, big.NewInt(3000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.GetEpochFromID(2).BeginHeight-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
		sendWithdrawTransaction(number-types.MinCalcRedeemHeight(2), gen, saddr1, big.NewInt(1000000000000000000), skey1, signer, statedb, fastChain, abiStaking, nil)
		sendGetDepositTransaction(number-types.MinCalcRedeemHeight(2)-11, gen, saddr1, skey1, signer, statedb, fastChain, abiStaking, nil)
	}

	manager := newTestPOSManager(101, executable)
	fmt.Println(" saddr1 ", manager.GetBalance(saddr1), " StakingAddress ", manager.GetBalance(types.StakingAddress), " ", types.ToTrue(manager.GetBalance(types.StakingAddress)))
	fmt.Println("epoch ", types.GetEpochFromID(1), " ", types.GetEpochFromID(2), " ", types.GetEpochFromID(3), " ", types.GetEpochFromID(4), " ", types.GetEpochFromID(5))
	fmt.Println("epoch ", types.GetEpochFromID(2), " ", types.MinCalcRedeemHeight(2))
}

func TestGetAddress(t *testing.T) {
	// Create a helper to check if a gas allowance results in an executable transaction
	skey, _ := crypto.HexToECDSA("308d34e60db54866100b395831f8a8b3c493f8fe733c53d3571da7025955bb54")
	saddr := crypto.PubkeyToAddress(skey.PublicKey)
	skey2, _ := crypto.HexToECDSA("48552cb89a19028d116c7853c460f0c76d50cddaf2d7c217ac611b696e4680c6")
	saddr2 := crypto.PubkeyToAddress(skey2.PublicKey)
	skey3, _ := crypto.HexToECDSA("6e7595dcb8eda2213c1a0940c4920ce7db89d8f805fc4e85567fd1355c83cff2")
	saddr3 := crypto.PubkeyToAddress(skey3.PublicKey)
	skey4, _ := crypto.HexToECDSA("62eab9d8657c25330f587c2fbb292a559ed27ddb391b9efe6014f920e67d2f1a")
	saddr4 := crypto.PubkeyToAddress(skey4.PublicKey)

	fmt.Println("saddr", saddr.String(), "saddr2", saddr2.String(), "saddr3", saddr3.String(), "saddr26 ", saddr4.String())

	skey5, _ := crypto.HexToECDSA("dbb0d9954bef0db91d7d15c44855cb0d0e662d01ac2a15d31d38724236802fbd")
	saddr5 := crypto.PubkeyToAddress(skey5.PublicKey)
	skey6, _ := crypto.HexToECDSA("2801712bcc44a58f4b2d0e74df50b9875747d60f0b8a133ea591276ca004ad3e")
	saddr6 := crypto.PubkeyToAddress(skey6.PublicKey)
	skey7, _ := crypto.HexToECDSA("5e6ea3e3ba8a3d8940088247eda01a0909320f729ae3afcdc5747b2ced1ac460")
	saddr7 := crypto.PubkeyToAddress(skey7.PublicKey)

	fmt.Println("saddr5", saddr5.String(), "saddr6", saddr6.String(), "saddr7", saddr7.String())
}
