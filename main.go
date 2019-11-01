package main

// #cgo CFLAGS: -I${SRCDIR}/TrustWalletCore
//// #cgo LDFLAGS: ${SRCDIR}/TrustWalletCore/libTrustWalletCore.a -lstdc++
// #include "TWHDWallet.h"
import "C"

import "fmt"

func main() {
	fmt.Print("I'm log.")
}

func test() {
	fmt.Println("--------------------")

	a := C.TWHDWalletIsValid("L3fi366VNv7arRq5FSbiN6tcJJB4")
	fmt.Println("Res = ", a)
}
