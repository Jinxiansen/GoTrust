package main

// #cgo CFLAGS: -I${SRCDIR}/TrustWalletCore
// #cgo LDFLAGS: -L${SRCDIR}/TrustWalletCore -lTrezorCrypto -lTrustWalletCore -lprotobuf -lc++
// #include "TWHDWallet.h"
// #include "TWString.h"
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

func test() {
	fmt.Println("--------------------")

	csMnemonic := C.CString(mnemonic)
	csPwd := C.CString("")
	defer C.free(unsafe.Pointer(csMnemonic))
	defer C.free(unsafe.Pointer(csPwd))

	mne := C.TWStringCreateWithUTF8Bytes(csMnemonic)
	pwd := C.TWStringCreateWithUTF8Bytes(csPwd)

	wallet := C.TWHDWalletCreateWithMnemonic(mne, pwd)

	privateKey := C.TWHDWalletGetExtendedPrivateKey(wallet, C.TWPurposeBIP44, C.TWCoinTypeQtum, C.TWHDVersionXPRV)
	pubKey := C.TWHDWalletGetExtendedPublicKey(wallet, C.TWPurposeBIP44, C.TWCoinTypeQtum, C.TWHDVersionXPUB)

	fmt.Println("PrivateKey = ", C.GoString(C.TWStringUTF8Bytes(privateKey)))
	fmt.Println("pubKey = ", C.GoString(C.TWStringUTF8Bytes(pubKey)))

}

// func createWallet() {

// 	mnemonic := "b"
// 	pwd := ""

// 	csMnemonic := C.CString(mnemonic)
// 	csPwd := C.CString(pwd)
// 	defer C.free(unsafe.Pointer(csMnemonic))
// 	defer C.free(unsafe.Pointer(csPwd))

// 	btMnemonic := C.TWStringCreateWithUTF8Bytes(csMnemonic)
// 	btPwd := C.TWStringCreateWithUTF8Bytes(csPwd)

// 	wallet := C.TWHDWalletCreateWithMnemonic(btMnemonic, btPwd)

// 	privateKey := C.TWHDWalletGetKey(wallet, C.TWCoinTypeDerivationPath(TWCoinTypeAeternity))

// }

// TEST(TWAeternityAddress, HDWallet) {
//     auto mnemonic = "shoot island position soft burden budget tooth cruel issue economy destroy above";
//     auto passphrase = "";

//     auto wallet = WRAP(TWHDWallet, TWHDWalletCreateWithMnemonic(
//             STRING(mnemonic).get(),
//             STRING(passphrase).get()
//     ));

//     auto privateKey = TWHDWalletGetKey(wallet.get(), TWCoinTypeDerivationPath(TWCoinTypeAeternity));
//     auto publicKey = TWPrivateKeyGetPublicKeyEd25519(privateKey);
//     auto address = TWAeternityAddressCreateWithPublicKey(publicKey);
//     auto addressStr = WRAPS(TWAeternityAddressDescription(address));

//     assertStringsEqual(addressStr, "ak_QDHJSfvHG9sDHBobaWt2TAGhuhipYjEqZEH34bWugpJfJc3GN");
// }
