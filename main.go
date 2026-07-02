package main

import (
	"fmt"
	"khqr-learn/internal/khqr"
)

func main() {
	qr, err := khqr.GenerateIndividual(khqr.IndividualInfo{
		BakongAccountID: "Sereyvuth@dev", // replace with your sandbox test account
		MerchantName:    "Test Name",
		MerchantCity:    "Phnom Penh",
		Amount:          1.00,
		Currency:        "USD",
		BillNumber:      "TEST001",
	})
	if err != nil {
		fmt.Println("generate error:", err)
		return
	}

	merchantQr, err := khqr.GenerateMerchant(
		khqr.MerchantInfo{
			IndividualInfo: khqr.IndividualInfo{
				BakongAccountID: "Sereyvuth@aba", // replace with your sandbox test account
				MerchantName:    "Test Name",
				MerchantCity:    "Phnom Penh",
				Amount:          1.00,
				Currency:        "USD",
				BillNumber:      "TEST001",
			},
			MerchantID: "TESTMERCHANTID1",
			AcquiringBank: "aba",
		})

	fmt.Println("QR:", merchantQr)

	hash := khqr.HashMD5(qr)
	fmt.Println("MD5:", hash)

	// wire with bakong
}
