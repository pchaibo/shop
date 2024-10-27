package tronApi

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func (t *TronApiEngine) Trconaddress() (address, privateKey, addhex string) {
	//tron 地址生成
	address, privateKey, addhex = GenerateKeyPair()
	//fmt.Printf(" address: %s\n privateKey:%s\n hexadd:%s\n", address, privateKey, addhex)
	//add = fmt.Printf(" address: %s", address)
	//resstr := Gettronhex(addhex)
	//fmt.Println("resstraddress:", resstr)
	//私钥转地址
	//PriKeytoaddress(privateKey)
	return address, privateKey, addhex
}

func GenerateKeyPair() (b5, pk, addhex string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	address = "41" + address[2:]
	addb, _ := hex.DecodeString(address)
	addhex = address
	firstHash := sha256.Sum256(addb)
	secondHash := sha256.Sum256(firstHash[:])
	secret := secondHash[:4]
	addb = append(addb, secret...)
	return base58.Encode(addb), hexutil.Encode(privateKeyBytes)[2:], addhex
}

// 创建私钥转地址
func PriKeytoaddress(priKeyHash string) {
	//priKeyHash := "796c823671b118258b53ef6056fd1f9fc96d125600f348f75f397b2000267fe8"
	priKey, err := crypto.HexToECDSA(priKeyHash)
	if err != nil {
		panic(err)
	}
	priKeyBytes := crypto.FromECDSA(priKey)
	fmt.Printf("私钥为: %s\n", hex.EncodeToString(priKeyBytes))

	pubKey := priKey.Public().(*ecdsa.PublicKey)
	// 获取公钥并去除头部0x04
	pubKeyBytes := crypto.FromECDSAPub(pubKey)[1:]
	fmt.Printf("公钥为: %s\n", hex.EncodeToString(pubKeyBytes))

	// 获取地址
	addr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("地址为: %s\n", strings.Replace(addr.Hex(), "0x", "41", 1))

}

// 41 转换base58adddrss
func Gettronhex(str string) (to string) {
	addb, _ := hex.DecodeString(str)
	firstHash := sha256.Sum256(addb)
	secondHash := sha256.Sum256(firstHash[:])
	secret := secondHash[:4]
	addb = append(addb, secret...)
	to = base58.Encode(addb)
	return
}

// 普通 eth
func Tadd() {
	// 1. 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 2. 从私钥生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ECDSA")
		return
	}

	// 3. 从公钥生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Tron Wallet Address:", address.Hex())
}

//
