package khqr

import (
	"errors"
	"fmt"
	"strings"
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
	var sb strings.Builder

	sb.WriteString(tlv(TagPayloadFormatIndicator, "01"))
	sb.WriteString(tlv(TagPointofInitiationMethod, "12"))

	IndividualInfoString := fmt.Sprintf("00%02d%s", len(info.BakongAccountID), info.BakongAccountID)	

	sb.WriteString(tlv(TagIndividual, IndividualInfoString))

	sb.WriteString(tlv(TagMerchantCategoryCode, "5999"))

	trxn_ccy, err := currencyString(info.Currency)
	if err != nil {
		fmt.Println("Error: wrong currency")
	}
	sb.WriteString(tlv(TagTransactionCurrency, trxn_ccy))

	amountString := fmt.Sprintf("%.2f", info.Amount)
	sb.WriteString(tlv(TagTransactionAmount, amountString))

	sb.WriteString(tlv(TagCountryCode, "KH"))

	sb.WriteString(tlv(TagMerchantName, info.MerchantName))
	sb.WriteString(tlv(TagMerchantCity, info.MerchantCity))

	additional_data := fmt.Sprintf("01%02d%s", len(info.BillNumber), info.BillNumber)
	sb.WriteString(tlv(TagAdditionalData, additional_data))

	finalPayload := sb.String()
	payloadForCRC := finalPayload + checksumPlaceholder()

	data := []byte(payloadForCRC)
	hexCode := fmt.Sprintf("%04X", crc16Hex(data))  // turn it into "3449" style text
	fullQR := string(finalPayload) + tlv(TagCRCchecksum, hexCode)

	return fullQR, nil
}


type MerchantInfo struct {
	IndividualInfo
	MerchantID    string
	AcquiringBank string
}

func GenerateMerchant(info MerchantInfo) (string, error) {
	var sb strings.Builder

	sb.WriteString(tlv(TagPayloadFormatIndicator, "01"))
	sb.WriteString(tlv(TagPointofInitiationMethod, "12"))

	sb.WriteString(
		tlv(
			TagMerchantAccountInfo,
			tlv(Tag2930BakongAccountId, info.BakongAccountID) +
			tlv(Tag2930MerchantId, info.MerchantID) +
			tlv(Tag2930AcquiringBankName, info.AcquiringBank),
		),
	)

	sb.WriteString(tlv(TagMerchantCategoryCode, "5999"))


	trxn_ccy, err := currencyString(info.Currency)
	if err != nil {
		fmt.Println("Error: wrong currency")
	}
	sb.WriteString(tlv(TagTransactionCurrency, trxn_ccy))

	amountString := fmt.Sprintf("%.2f", info.Amount)
	sb.WriteString(tlv(TagTransactionAmount, amountString))

	sb.WriteString(tlv(TagCountryCode, "KH"))

	sb.WriteString(tlv(TagMerchantName, info.MerchantName))
	sb.WriteString(tlv(TagMerchantCity, info.MerchantCity))

	sb.WriteString(
		tlv(
			TagAdditionalData,
			tlv(Tag62BillNumber, info.BillNumber),
		),
	)

	finalPayload := sb.String()
	payloadForCRC := finalPayload + checksumPlaceholder()

	data := []byte(payloadForCRC)
	hexCode := fmt.Sprintf("%04X", crc16Hex(data))  // turn it into "3449" style text
	fullQR := string(finalPayload) + tlv(TagCRCchecksum, hexCode)

	return fullQR, nil
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
