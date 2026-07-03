package main

import (
	"fmt"
	"log"
	"os"

	"khqr-learn/internal/bakong"
	"khqr-learn/internal/khqr"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	qr, err := khqr.GenerateIndividual(khqr.IndividualInfo{
		BakongAccountID: "Sereyvuth@dev",
		MerchantName:    "Sereyvuth Duong",
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
			AccountID: "Sereyvuth@aba",
			MerchantName:    "Test Name",
			MerchantCity:    "Phnom Penh",
			Amount:          1.00,
			Currency:        "USD",
			BillNumber:      "TEST001",
			MerchantID: "TESTMERCHANTID1",
			AcquiringBank: "aba",
		})

	fmt.Println("QR:", merchantQr)

	hashIndi := khqr.HashMD5(qr)
	hashMerc := khqr.HashMD5(merchantQr)
	fmt.Println("MD5 for Individual:", hashIndi)
	fmt.Println("MD5 for Merchant:", hashMerc)

	var token string = os.Getenv("BAKONG_TOKEN")

	client := bakong.NewClient(token)
	check_account, err := client.CheckAccountByID("testvuthvuth@devb")
	if err != nil {
		fmt.Printf("Error trying to check account: %v\n", err)
		return
	}
	if check_account {
		fmt.Printf("Account Id Exist")
	} else {
		fmt.Println("Account ID not exist")
	}
}
