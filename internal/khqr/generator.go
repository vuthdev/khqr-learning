package khqr

import (
	"errors"
	"fmt"
)

type IndividualInfo struct {
	BakongAccountID string  // e.g. "kimhak@dev"
	MerchantName    string  // display name, e.g. "Kimhak"
	MerchantCity    string  // e.g. "Phnom Penh"
	Amount          float64 // 0 for a static/no-fixed-amount QR
	Currency        string  // "USD" or "KHR"
	BillNumber      string  // optional, your own reference number
}

func GenerateIndividual(info IndividualInfo) (string, error) {
	newQr, err := Builder().
		Individual(
			info.BakongAccountID,
		).
		MerchantName(info.MerchantName).
		MerchantCity(info.MerchantCity).
		Amount(info.Amount).
		Currency(info.Currency).
		// BillNumber(info.BillNumber)
		Build()


	if err != nil {
		return "", fmt.Errorf("Error generating KHQR")
	}

	fmt.Printf("Generated QR string: %v\n", newQr)

	return newQr, nil
}

type MerchantInfo struct {
	AccountID string
	MerchantID    string
	AcquiringBank string
	MerchantName    string  // display name, e.g. "Kimhak"
	MerchantCity    string  // e.g. "Phnom Penh"
	Amount          float64 // 0 for a static/no-fixed-amount QR
	Currency        string  // "USD" or "KHR"
	BillNumber      string
}

func GenerateMerchant(info MerchantInfo) (string, error) {
	newQr, err := Builder().
		Merchant(
			info.AccountID,
			info.MerchantID,
			info.AcquiringBank,
		).
		MerchantName(info.MerchantName).
		MerchantCity(info.MerchantCity).
		Amount(info.Amount).
		Currency(info.Currency).
		// BillNumber(info.BillNumber).
		Build()


	if err != nil {
		return "", fmt.Errorf("Error generating KHQR")
	}

	fmt.Printf("Generated QR string: %v\n", newQr)

	return newQr, nil
}

func currencyString(currency string) (string, error) {
	var trxn_currency string
	
	switch currency {
		case "USD":
			trxn_currency = fmt.Sprintf("%s", USD)
		
		case "KHR":
			trxn_currency = fmt.Sprintf("%s", KHR)
		default: return "", errors.New("Wrong currency")
	}

	return trxn_currency, nil
}
