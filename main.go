package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"khqr-learn/internal/bakong"
	"khqr-learn/internal/khqr"

	"github.com/joho/godotenv"
	"github.com/skip2/go-qrcode"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	outputFolder := "qrcodes"
	fileName := "my_khqr_invoice.png"

	err = os.Mkdir(outputFolder, os.ModePerm) 
	if err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}

	filePath := filepath.Join(outputFolder, fileName)

	// merchantQr, err := khqr.GenerateMerchant(
	// 	khqr.MerchantInfo{
	// 		AccountID: "sereyvuth_duong@bkrt",
	// 		MerchantName:    "Sereyvuth Duong",
	// 		MerchantCity:    "Phnom Penh",
	// 		Amount:          1.00,
	// 		Currency:        "USD",
	// 		BillNumber:      "TEST001",
	// 		MerchantID: "TESTMERCHANTID1",
	// 		AcquiringBank: "aba",
	// 	})

	// fmt.Println("QR:", merchantQr)

	// hashMerc := khqr.HashMD5(merchantQr)
	
	// fmt.Println("MD5 for Merchant:", hashMerc)

	var token string = os.Getenv("BAKONG_TOKEN")

	// example data
	var accountID string = "sereyvuth_duong@test"
	var merchantName string = "Sereyvuth Duong"
	var merchantCity string = "Phnom Penh"
	var amount float64 = 1.00
	var currency string = "USD"

	client := bakong.NewClient(token)
	check_account, err := client.CheckAccountByID(accountID)
	if err != nil {
		fmt.Printf("Error trying to check account: %v\n", err)
		return
	}
	if check_account {
		qr, err := khqr.GenerateIndividual(khqr.IndividualInfo{
			BakongAccountID: accountID,
			MerchantName:    merchantName,
			MerchantCity:    merchantCity,
			Amount:          amount,
			Currency:        currency,
			//BillNumber:      "TEST001",
		})
		if err != nil {
			fmt.Println("generate error:", err)
			return
		}

		fmt.Println("KHQR string: ", qr)
		hashIndi := khqr.HashMD5(qr)
		fmt.Println("MD5 for Individual:", hashIndi)
		fmt.Printf("Account Exist")

		fmt.Println("generating QR code...")
		pngBytes, _ := qrcode.Encode(qr, qrcode.Medium, 256)

		err = os.WriteFile(filePath, pngBytes, 0644)
		if err != nil {
			log.Fatalf("Failed to save the PNG file: %v", err)
		}

		fmt.Printf("🎉 Success! Your KHQR image has been saved to: %s\n", filePath)
	} else {
		fmt.Println("Account not exist")
	}
}
