package tronApi

import (
	"math/big"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/shopspring/decimal"
)

func TestTrc20Balance(t *testing.T) {
	amount, err := apiEngine.Trc20Balance("", "")
	if err != nil {
		t.Log(err)
	}
	t.Log(amount)

}

func TestHexTo10(t *testing.T) {
	data := "000000000000000000000000000000000000000000000006aaf7c8516f4123ea"
	if common.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) == 64 {
		t.Log(data)

		var n big.Int
		m, ok := n.SetString(data, 16)
		if ok {
			t.Log(n.Int64())
		}
		t.Log(m.String())

		balanceDecimal := decimal.NewFromBigInt(m, 0).Div(decimal.New(1, 18))
		t.Log(balanceDecimal.String())
		t.Log(balanceDecimal.IntPart())

	}
}

func TestEncodeBalanceParams(t *testing.T) {
	base58address := ""
	t.Log(apiEngine.encodeBalanceParams(base58address))
}

func TestParseTRC20NumericProperty(t *testing.T) {
	data := "0000000000000000000000000000000000000000000000000000000b9dbab820"
	t.Log(apiEngine.parseTRC20NumericProperty(data))
}

func TestTrc20Info(t *testing.T) {
	trc20Info, err := apiEngine.Trc20Info("")
	if err != nil {
		t.Log(err)
	}
	t.Log(trc20Info)
}

//Trc20Gas

func TestTrc20Gas(t *testing.T) {
	from := ""
	to := ""
	contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" //TNuoKL1ni8aoshfFL1ASca1Gou9RXwAzfn
	amount := big.NewInt(2123456)
	// tx参数
	req := &ReqTriggerSmartContract{
		OwnerAddress:     from,
		ContractAddress:  contract,
		FunctionSelector: "transfer(address,uint256)",
		Parameter:        "",
		CallValue:        0,
		FeeLimit:         20000000,
		Visible:          true,
	}
	rep, err := apiEngine.Trc20TriggerSmartContract(req, to, amount)
	if err != nil {
		t.Log(err.Error())
	}
	//t.Log(fmt.Sprintf("%+v", rep))
	t.Log("广播前hash：", rep.Transaction.TxID)
	rep.Transaction.Signature, err = apiEngine.signHashTx(rep.Transaction.TxID,
		"")

	gas, err := apiEngine.Trc20Gas(
		from,
		to,
		contract,
		amount,
	)
	if err != nil {
		t.Log(err)
	}
	t.Log("-------------------->gas:", gas)

	hash, err := apiEngine.BroadcastTx(rep.Transaction)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log("广播后hash：", hash)
}
