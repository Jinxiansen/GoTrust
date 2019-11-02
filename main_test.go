package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	fmt.Println("Print: Test_main")

	test()
}

func Test_AE_Wallet(t *testing.T) {

	createAEWallet()
}

func Test_ConvertTWString(t *testing.T) {

	s := convertToTWString("hello")
	fmt.Println("S = ", s)
}
