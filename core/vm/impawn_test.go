package vm

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/truechain/truechain-engineering-code/common"
	"github.com/truechain/truechain-engineering-code/core/types"
	"github.com/truechain/truechain-engineering-code/crypto"
	"github.com/truechain/truechain-engineering-code/rlp"
	"math/big"
	"testing"
)

/////////////////////////////////////////////////////////////////////

func TestImpawnImpl(t *testing.T) {
	types.PreselectionPeriod = 2
	types.EpochLength = 5
	types.MaxRedeemHeight = 0
	fmt.Println(" epoch 1 ", types.GetEpochFromID(1))
	fmt.Println(" epoch 2 ", types.GetEpochFromID(2))
	impl := NewImpawnImpl()
	for i := uint64(0); i < types.PreselectionPeriod+types.EpochLength+1; i++ {
		value := big.NewInt(100)
		priKey, _ := crypto.GenerateKey()
		from := crypto.PubkeyToAddress(priKey.PublicKey)
		pub := crypto.FromECDSAPub(&priKey.PublicKey)
		fmt.Println(" ", from.String())
		impl.InsertSAccount2(20+i, from, pub, new(big.Int).Sub(value, big.NewInt(int64(10*i))), big.NewInt(0), true)
		priKeyDA, _ := crypto.GenerateKey()
		daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
		impl.InsertDAccount2(20+i, daAddress, from, new(big.Int).Sub(value, big.NewInt(int64(10*i))))
	}

	for i := uint64(0); i < types.EpochLength; i++ {
		value := big.NewInt(100)
		priKey, _ := crypto.GenerateKey()
		from := crypto.PubkeyToAddress(priKey.PublicKey)
		pub := crypto.FromECDSAPub(&priKey.PublicKey)
		impl.InsertSAccount2(28+i, from, pub, value, big.NewInt(0), true)
		priKeyDA, _ := crypto.GenerateKey()
		daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
		impl.InsertDAccount2(28+i, daAddress, from, value)
	}

	impl.DoElections(2, 28)
	fmt.Println(impl.getCurrentEpochInfo())
	fmt.Println(impl.getElections(2))
	fmt.Println(impl.getCurrentEpoch())

	//impl.Shift()
	//impl.Reward()
	//impl.CancelDAccount()
	//impl.CancelSAccount()
	//impl.RedeemSAccount()
	//impl.RedeemDAccount()

	for i := uint64(0); i < types.EpochLength; i++ {
		value := big.NewInt(100)
		priKey, _ := crypto.GenerateKey()
		from := crypto.PubkeyToAddress(priKey.PublicKey)
		pub := crypto.FromECDSAPub(&priKey.PublicKey)
		impl.InsertSAccount2(21+i, from, pub, value, big.NewInt(0), true)
		priKeyDA, _ := crypto.GenerateKey()
		daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
		impl.InsertDAccount2(21+i, daAddress, from, value)
		if i == types.EpochLength-3 {
			impl.DoElections(3, 28)
		}
	}

	fmt.Println(impl.getCurrentEpochInfo())
	fmt.Println(impl.getElections(2))
	fmt.Println(impl.getCurrentEpoch())

}

func TestEpoch(t *testing.T) {
	types.DposForkPoint = 20
	fmt.Println(" first  ", types.GetFirstEpoch())
	fmt.Println(" epoch 2 ", types.GetEpochFromID(1))
	fmt.Println(" epoch 2 ", types.GetEpochFromID(2))
	fmt.Println(" epoch 2 ", types.GetEpochFromID(2))
	fmt.Println(" epoch 3 ", types.GetEpochFromID(3))
	fmt.Println(" epoch 4 ", types.GetEpochFromID(4))
	fmt.Println(types.GetEpochFromHeight(0))
	fmt.Println(types.GetEpochFromHeight(12000))
	fmt.Println(types.GetEpochFromHeight(12021))
	fmt.Println(types.GetEpochFromHeight(22021))
	fmt.Println(types.GetEpochFromHeight(32020))
	fmt.Println(types.GetEpochFromHeight(32021))
	fmt.Println(types.GetEpochFromHeight(42020))
	fmt.Println("GetEpochFromRange ", types.GetEpochFromRange(21, 12020))
	fmt.Println("GetEpochFromRange ", types.GetEpochFromRange(21, 12021))
	fmt.Println("GetEpochFromRange ", types.GetEpochFromRange(21, 22020))
	fmt.Println("GetEpochFromRange ", types.GetEpochFromRange(21, 22021))
}

// Underlying data structure
/////////////////////////////////////////////////////////////////////
func TestImpawnUnit(t *testing.T) {
	types.DposForkPoint = 1
	types.PreselectionPeriod = 0
	types.EpochLength = 50
	types.MaxRedeemHeight = 0
	fmt.Println(" epoch 1 ", types.GetEpochFromID(1))
	fmt.Println(" epoch 2 ", types.GetEpochFromID(2))
	priKey, _ := crypto.GenerateKey()
	fmt.Printf("%x \n", crypto.FromECDSAPub(&priKey.PublicKey))
	coinbase := crypto.PubkeyToAddress(priKey.PublicKey)
	value := make([]*PairstakingValue, 0)
	Reward := make([]*RewardItem, 0)
	for i := 3; i > 0; i-- {
		pv := &PairstakingValue{
			amount: new(big.Int).SetUint64(10*uint64(i) + 10),
			height: new(big.Int).SetUint64(10*uint64(i) + 60),
			state:  types.StateStakingAuto,
		}
		value = append(value, pv)
		ri := &RewardItem{
			Amount: new(big.Int).SetUint64(10*uint64(i) + 10),
			Height: new(big.Int).SetUint64(10*uint64(i) + 10),
		}
		Reward = append(Reward, ri)
	}
	iMunit := &impawnUnit{
		address: coinbase,
		value:   value,
		redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
			Amount:  new(big.Int).SetUint64(15),
			EpochID: 1,
			State:   types.StateStakingAuto,
		}),
	}

	iunit := &impawnUnit{
		address: coinbase,
		value: append(make([]*PairstakingValue, 0), &PairstakingValue{
			amount: new(big.Int).SetUint64(15),
			height: new(big.Int).SetUint64(65),
			state:  types.StateStakingAuto,
		}),
		redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
			Amount:  new(big.Int).SetUint64(15),
			EpochID: 1,
			State:   types.StateStakingAuto,
		}),
	}

	iMunit.update(iunit, true)
	for i, value := range iMunit.value {
		fmt.Printf("%d %d %d %d \n", i, value.height, value.amount, value.state)
	}

	fmt.Println(iMunit.getAllStaking(90))
	fmt.Println(iMunit.GetRewardAddress().String())

	pv := &PairstakingValue{
		amount: new(big.Int).SetUint64(25),
		height: new(big.Int).SetUint64(60),
		state:  types.StateStakingAuto,
	}
	iMunit.value = append(iMunit.value, pv)
	iMunit.sort()

	for i, value := range iMunit.value {
		fmt.Printf("%d %d %d %d \n", i, value.height, value.amount, value.state)
	}

	iMunit.stopStakingInfo(new(big.Int).SetInt64(30), new(big.Int).SetInt64(35))

	iMunit.merge(1, 100)

	_, value1, _ := iMunit.redeeming(90, new(big.Int).SetInt64(60))

	for i, value := range iMunit.value {
		fmt.Printf("%d %d %d %d \n", i, value.height, value.amount, value.state)
	}
	fmt.Printf("insertRedeemInfo redeemInof %d %v %d \n", iMunit.redeemInof[0].Amount, iMunit.redeemInof, value1)

	iMunit.finishRedeemed()

	for i, value := range iMunit.value {
		fmt.Printf("merge %d %d %d  %d \n", i, value.height, value.amount, value.state)
	}
}

func initialStakingAccount(n int, m int, stride int, SAaddress common.Address, DAaddress common.Address, priKey *ecdsa.PrivateKey, priKeyDA *ecdsa.PrivateKey) *StakingAccount {
	var das []*DelegationAccount
	for k := 0; k < n; k++ {
		da := &DelegationAccount{
			deleAddress: SAaddress,
			unit:        initialImpawnUnit(m, stride, DAaddress),
		}
		das = append(das, da)
	}

	da := &DelegationAccount{
		deleAddress: SAaddress,
		unit:        initialImpawnUnit(m, stride-5, DAaddress),
	}
	das = append(das, da)

	saccount := &StakingAccount{
		unit:       initialImpawnUnit(m, stride, SAaddress),
		votepubkey: crypto.FromECDSAPub(&priKey.PublicKey),
		fee:        new(big.Int).SetUint64(1000),
		committee:  true,
		delegation: das,
		modify: &AlterableInfo{
			fee:        new(big.Int).SetUint64(1000),
			votePubkey: crypto.FromECDSAPub(&priKeyDA.PublicKey),
		},
	}
	return saccount
}

func initialImpawnUnit(n int, stride int, address common.Address) *impawnUnit {
	value := make([]*PairstakingValue, 0)
	Reward := make([]*RewardItem, 0)
	for i := 0; i < n; i++ {
		pv := &PairstakingValue{
			amount: new(big.Int).SetUint64(uint64(stride*i + stride)),
			height: new(big.Int).SetUint64(uint64(stride*i + stride)),
			state:  types.StateStakingAuto,
		}
		value = append(value, pv)
		ri := &RewardItem{
			Amount: new(big.Int).SetUint64(uint64(stride*i + stride)),
			Height: new(big.Int).SetUint64(uint64(stride*i + stride)),
		}
		Reward = append(Reward, ri)
	}
	iMunit := &impawnUnit{
		address: address,
		value:   value,
		redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
			Amount:  new(big.Int).SetUint64(1000),
			EpochID: 2,
			State:   types.StateRedeem,
		}),
	}
	return iMunit
}

func TestDelegationAccount(t *testing.T) {

	priKey, _ := crypto.GenerateKey()
	saAddress := crypto.PubkeyToAddress(priKey.PublicKey)
	priKeyDA, _ := crypto.GenerateKey()
	daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
	da := &DelegationAccount{
		deleAddress: saAddress,
		unit:        initialImpawnUnit(3, 5, daAddress),
	}

	priKey1, _ := crypto.GenerateKey()
	saAddress1 := crypto.PubkeyToAddress(priKey1.PublicKey)
	priKeyDA1, _ := crypto.GenerateKey()
	daAddress1 := crypto.PubkeyToAddress(priKeyDA1.PublicKey)
	da1 := &DelegationAccount{
		deleAddress: saAddress1,
		unit:        initialImpawnUnit(3, 4, daAddress1),
	}
	da.update(da1, false)
	da.stopStakingInfo(new(big.Int).SetInt64(30), new(big.Int).SetInt64(15))

	da.redeeming(90, new(big.Int).SetInt64(60))

	da.finishRedeemed()

	da.merge(1, 100)
}

func TestStakingAccount(t *testing.T) {

	priKey, _ := crypto.GenerateKey()
	saAddress := crypto.PubkeyToAddress(priKey.PublicKey)
	priKeyDA, _ := crypto.GenerateKey()
	daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
	saccount := initialStakingAccount(3, 1, 100, saAddress, daAddress, priKey, priKeyDA)

	fmt.Println(saccount.getAllStaking(300), " ", saccount.getDA(daAddress).getAllStaking(300), " ", saccount.unit.getAllStaking(300))

	priKey1, _ := crypto.GenerateKey()
	saAddress1 := crypto.PubkeyToAddress(priKey1.PublicKey)
	priKeyDA1, _ := crypto.GenerateKey()
	daAddress1 := crypto.PubkeyToAddress(priKeyDA1.PublicKey)
	saccount1 := initialStakingAccount(3, 1, 50, saAddress1, daAddress1, priKey1, priKeyDA1)

	saccount.update(saccount1, 300, false, false)
	fmt.Println("Committee ", saccount.isInCommittee())
	saccount.stopStakingInfo(new(big.Int).SetInt64(1000), new(big.Int).SetInt64(300))

	saccount.redeeming(90, new(big.Int).SetInt64(60))

	saccount.finishRedeemed()

	saccount.merge(1, 300)
}

func TestSAImpawns(t *testing.T) {
	var sas []*StakingAccount

	sa := common.Address{}

	for i := 0; i < 3; i++ {
		priKey, _ := crypto.GenerateKey()
		saAddress := crypto.PubkeyToAddress(priKey.PublicKey)
		priKeyDA, _ := crypto.GenerateKey()
		daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
		saccount := initialStakingAccount(1, 1, 10, saAddress, daAddress, priKey, priKeyDA)
		sas = append(sas, saccount)
		sa = saAddress
	}
	SAIs := SAImpawns(sas)
	fmt.Println(" sa ", SAIs.getValidStaking(15), " all ", SAIs.getAllStaking(15), " da ", SAIs.getSA(sa))
	priKey, _ := crypto.GenerateKey()
	saAddress := crypto.PubkeyToAddress(priKey.PublicKey)
	priKeyDA, _ := crypto.GenerateKey()
	daAddress := crypto.PubkeyToAddress(priKeyDA.PublicKey)
	saccount := initialStakingAccount(1, 1, 9, saAddress, daAddress, priKey, priKeyDA)
	SAIs.update(saccount, 30, false, false)
	saccount = initialStakingAccount(1, 1, 15, saAddress, daAddress, priKey, priKeyDA)
	SAIs.update(saccount, 30, false, false)

	SAIs.sort(15, false)
}

// RLP
/////////////////////////////////////////////////////////////////////
func TestRlpImpawnImpl(t *testing.T) {
	impl := makeImpawnImpl()
	bzs, err := rlp.EncodeToBytes(impl)
	if err != nil {
		fmt.Println(err.Error())
	}

	var tmp ImpawnImpl
	err = rlp.DecodeBytes(bzs, &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}

	for m, sas := range tmp.accounts {
		for n, account := range sas {
			fmt.Printf("account %d %d %v %d \n", m, n, account, account.fee)
		}
	}
	fmt.Printf("%v \n", tmp.curEpochID)
}

func makeImpawnImpl() *ImpawnImpl {
	accounts := make(map[uint64]SAImpawns)

	priKey, _ := crypto.GenerateKey()
	coinbase := crypto.PubkeyToAddress(priKey.PublicKey)

	for i := 0; i < types.MixEpochCount; i++ {

		accounts[uint64(i)] = make(SAImpawns, 3)
		var sas []*StakingAccount
		for j := 0; j < 3; j++ {
			unit := &impawnUnit{
				address: coinbase,
				value: append(make([]*PairstakingValue, 0), &PairstakingValue{
					amount: new(big.Int).SetUint64(1000),
					height: new(big.Int).SetUint64(1000),
					state:  types.StateRedeem,
				}),
				redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
					Amount:  new(big.Int).SetUint64(1000),
					EpochID: 2,
					State:   types.StateRedeem,
				}),
			}
			das := []*DelegationAccount{}
			for k := 0; k < 3; k++ {
				da := &DelegationAccount{
					deleAddress: coinbase,
					unit:        unit,
				}
				das = append(das, da)
			}

			saccount := &StakingAccount{
				unit:       unit,
				votepubkey: crypto.FromECDSAPub(&priKey.PublicKey),
				fee:        new(big.Int).SetUint64(1000),
				committee:  true,
				delegation: das,
				modify: &AlterableInfo{
					fee:        new(big.Int).SetUint64(1000),
					votePubkey: crypto.FromECDSAPub(&priKey.PublicKey),
				},
			}
			sas = append(sas, saccount)
		}
		copy(accounts[uint64(i)], sas)
	}

	impl := &ImpawnImpl{
		curEpochID: 1000,
		accounts:   accounts,
	}
	return impl
}

func TestRlpDelegationAccount(t *testing.T) {

	mint := makeDelegationAccount()
	bzs, err := rlp.EncodeToBytes(mint)
	if err != nil {
		fmt.Println(err.Error())
	}

	var tmp DelegationAccount
	err = rlp.DecodeBytes(bzs, &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func makeDelegationAccount() *DelegationAccount {
	priKey, _ := crypto.GenerateKey()
	coinbase := crypto.PubkeyToAddress(priKey.PublicKey)
	iMunit := &impawnUnit{
		address: coinbase,
		value: append(make([]*PairstakingValue, 0), &PairstakingValue{
			amount: new(big.Int).SetUint64(1000),
			height: new(big.Int).SetUint64(101),
			state:  types.StateRedeem,
		}),
		redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
			Amount:  new(big.Int).SetUint64(1000),
			EpochID: 2,
			State:   types.StateRedeem,
		}),
	}
	da := &DelegationAccount{
		deleAddress: coinbase,
		unit:        iMunit,
	}
	return da
}

type SAccount struct {
	Fee    *big.Float
	Amount *big.Int
}

type DAccount struct {
	fee    *big.Float
	Amount *big.Int
}

type sAccount struct {
	Fee    *big.Float
	Amount *big.Int
}

func TestRlpAccount(t *testing.T) {
	SA := &SAccount{
		Fee:    new(big.Float).SetFloat64(101.1),
		Amount: new(big.Int).SetUint64(1000),
	}
	bzs, err := rlp.EncodeToBytes(SA)
	if err != nil {
		fmt.Println(err.Error())
	}

	var tmp SAccount
	err = rlp.DecodeBytes(bzs, &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("11 ", tmp.Fee, " ", tmp.Amount)

	DA := &DAccount{
		fee:    new(big.Float).SetFloat64(101.1),
		Amount: new(big.Int).SetUint64(1000),
	}
	dzs, err := rlp.EncodeToBytes(DA)
	if err != nil {
		fmt.Println(err.Error())
	}

	var dmp DAccount
	err = rlp.DecodeBytes(dzs, &dmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("22 ", dmp.fee, " ", dmp.Amount)

	sA := &sAccount{
		Fee:    new(big.Float).SetFloat64(101.1),
		Amount: new(big.Int).SetUint64(1000),
	}
	szs, err := rlp.EncodeToBytes(sA)
	if err != nil {
		fmt.Println(err.Error())
	}

	var smp sAccount
	err = rlp.DecodeBytes(szs, &smp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("33 ", tmp.Fee, " ", smp.Amount)
}

func TestRlpimpawnUnit(t *testing.T) {

	mint := makeUnit()
	bzs, err := rlp.EncodeToBytes(mint)
	if err != nil {
		fmt.Println(err.Error())
	}

	var tmp impawnUnit
	err = rlp.DecodeBytes(bzs, &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, ri := range tmp.value {
		fmt.Printf("account %d %v \n", i, ri)
	}
}

func makeUnit() *impawnUnit {
	priKey, _ := crypto.GenerateKey()
	coinbase := crypto.PubkeyToAddress(priKey.PublicKey)
	iMunit := &impawnUnit{
		address: coinbase,
		value: append(make([]*PairstakingValue, 0), &PairstakingValue{
			amount: new(big.Int).SetUint64(1000),
			height: new(big.Int).SetUint64(101),
			state:  types.StateRedeem,
		}),
		redeemInof: append(make([]*RedeemItem, 0), &RedeemItem{
			Amount:  new(big.Int).SetUint64(1000),
			EpochID: 2,
			State:   types.StateRedeem,
		}),
	}
	return iMunit
}

func TestRlpAlterableInfo(t *testing.T) {
	impl := makeAlterableInfo()
	bzs, err := rlp.EncodeToBytes(impl)
	if err != nil {
		fmt.Println(err.Error())
	}

	var tmp AlterableInfo
	err = rlp.DecodeBytes(bzs, &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("account %d %d \n", tmp.fee, tmp.votePubkey)

}

func makeAlterableInfo() *AlterableInfo {
	priKey, _ := crypto.GenerateKey()

	modify := &AlterableInfo{
		fee:        nil,
		votePubkey: crypto.FromECDSAPub(&priKey.PublicKey),
	}
	return modify
}

/////////////////////////////////////////////////////////////////////