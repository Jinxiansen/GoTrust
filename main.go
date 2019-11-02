package main

// #cgo CFLAGS: -I${SRCDIR}/TrustWalletCore
// #cgo LDFLAGS: -L${SRCDIR}/TrustWalletCore -lTrezorCrypto -lTrustWalletCore -lprotobuf -lc++
// #include "TWHDWallet.h"
// #include "TWString.h"
// #include "TWAeternityAddress.h"
// #include "TWBitcoinAddress.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print("I'm log.")
	test()
}

var mnemonic = "ripple scissors kick mammal hire column oak again sun offer wealth tomorrow wagon turn fatal"
var boyMnemonic = "boy video model trap novel debris clip father art era edit salute"

func test() {
	fmt.Println("--------------------")

	wallet := C.TWHDWalletCreateWithMnemonic(convertToTWString(boyMnemonic), convertToTWString(""))

	privateKey := C.TWHDWalletGetExtendedPrivateKey(wallet, C.TWPurposeBIP44, C.TWCoinTypeQtum, C.TWHDVersionXPRV)
	pubKey := C.TWHDWalletGetExtendedPublicKey(wallet, C.TWPurposeBIP44, C.TWCoinTypeQtum, C.TWHDVersionXPUB)

	path := C.CString("m/88'/0'/0'")

	hdPubKey := C.TWHDWalletGetPublicKeyFromExtended(pubKey, C.TWStringCreateWithUTF8Bytes(path))
	address := C.TWBitcoinAddressCreateWithPublicKey(hdPubKey, C.TWCoinTypeP2pkhPrefix(C.TWCoinTypeQtum))
	addressStr := C.TWBitcoinAddressDescription(address)

	fmt.Println("PrivateKey = ", convertGoString(privateKey))
	fmt.Println("pubKey = ", convertGoString(pubKey))
	fmt.Println("Address: ", convertGoString(addressStr))

}

func createAEWallet() {

	mnemonic := "boy video model trap novel debris clip father art era edit salute"
	pwd := ""

	btMnemonic := convertToTWString(mnemonic)
	btPwd := convertToTWString(pwd)

	wallet := C.TWHDWalletCreateWithMnemonic(btMnemonic, btPwd)

	privateKey := C.TWHDWalletGetKey(wallet, C.TWCoinTypeDerivationPath(C.TWCoinTypeAeternity))
	privateKeyStr := C.TWPrivateKeyDescription(privateKey)

	pubKey := C.TWPrivateKeyGetPublicKeyEd25519(privateKey)
	pubKeyStr := C.TWPublicKeyDescription(pubKey)

	address := C.TWAeternityAddressCreateWithPublicKey(pubKey)
	addressStr := C.TWAeternityAddressDescription(address)

	fmt.Println("PrivateKey: ", privateKeyStr)
	fmt.Println("PubKey: ", convertGoString(pubKeyStr))
	fmt.Println("addressStr ", convertGoString(addressStr))

	// -------------------------

	privData := C.TWDataCreateWithHexString(privateKeyStr)
	impPriv := C.TWPrivateKeyCreateWithData(privData)

	impPub := C.TWPrivateKeyGetPublicKeyEd25519(impPriv)
	impPubStr := C.TWPublicKeyDescription(impPub)

	impAddr := C.TWAeternityAddressCreateWithPublicKey(impPub)
	impAddrStr := C.TWAeternityAddressDescription(impAddr)

	fmt.Println("Wif 公钥：", convertGoString(impPubStr))
	fmt.Println("Wif 地址：", convertGoString(impAddrStr))

}

func convertToTWString(str string) unsafe.Pointer {
	csStr := C.CString(str)
	defer C.free(unsafe.Pointer(csStr))
	twStr := C.TWStringCreateWithUTF8Bytes(csStr)
	return twStr
}

func convertGoString(twStr unsafe.Pointer) string {
	return C.GoString(C.TWStringUTF8Bytes(twStr))
}
