package khqr

import (
	"fmt"
	"strings"
	"time"
)

type KHQRBuilder struct {
	isMerchant bool
	accountId string
	merchantName string
	merchantCity string
	merchantID string // for merchant only
	acquiringBank string // for merchant only
	amount float64
	currency string
	billNumber string
}

func Builder() *KHQRBuilder {
	return &KHQRBuilder{}
}

func (b *KHQRBuilder) Individual(accountId string) *KHQRBuilder {
	b.isMerchant = false
	b.accountId = accountId
	return b
}

func (b *KHQRBuilder) Merchant(accountId string, merchantId string, acquiringBank string) *KHQRBuilder {
	b.isMerchant = true
	b.accountId = accountId
	b.merchantID = merchantId
	b.acquiringBank = acquiringBank
	return b
}

func (b *KHQRBuilder) MerchantName(merchantName string) *KHQRBuilder {
	b.merchantName = merchantName
	return b
}

func (b *KHQRBuilder) MerchantCity(merchantCity string) *KHQRBuilder {
	b.merchantCity = merchantCity
	return b
}

func (b *KHQRBuilder) Amount(amount float64) *KHQRBuilder {
	b.amount = amount
	return b
}

func (b *KHQRBuilder) Currency(currency string) *KHQRBuilder {
	b.currency = currency
	return b
}

// func (b *KHQRBuilder) BillNumber(billNumber string) *KHQRBuilder {
// 	b.billNumber = billNumber
// 	return b
// }

func (b *KHQRBuilder) Build() (string, error) {
	var sb strings.Builder

	sb.WriteString(tlv(TagPayloadFormatIndicator, "01"))

	sb.WriteString(tlv(TagPointofInitiationMethod, "12"))

	if b.isMerchant {
		sb.WriteString(
			tlv(
				TagMerchantAccountInfo,
				tlv(Tag2930BakongAccountId, b.accountId) +
				tlv(Tag2930MerchantId, b.merchantID) +
				tlv(Tag2930AcquiringBankName, b.acquiringBank),
			),
		)
	} else {
		sb.WriteString(tlv(TagIndividual, tlv(Tag2930BakongAccountId, b.accountId)))
	}

	sb.WriteString(tlv(TagMerchantCategoryCode, "5999"))

	trxn_ccy, err := currencyString(b.currency)
	if err != nil {
		fmt.Println("Error: wrong currency")
	}
	sb.WriteString(tlv(TagTransactionCurrency, trxn_ccy))

	amountString := fmt.Sprintf("%.2f", b.amount)
	sb.WriteString(tlv(TagTransactionAmount, amountString))

	sb.WriteString(tlv(TagCountryCode, "KH"))

	sb.WriteString(tlv(TagMerchantName, b.merchantName))
	sb.WriteString(tlv(TagMerchantCity, b.merchantCity))

	// sb.WriteString(
	// 	tlv(
	// 		TagAdditionalData,
	// 		tlv(Tag62BillNumber, b.billNumber),
	// 	),
	// )

	now := time.Now()
	expires := now.AddDate(0, 0, 1)

	sb.WriteString(
		tlv(
			TagKhqrSpecific, 
			tlv(TagKhqrSpecific99_00, fmt.Sprint(now.UnixMilli())) +
			tlv(TagKhqrSpecific99_01, fmt.Sprint(expires.UnixMilli())),
		),
	)

	finalPayload := sb.String()
	payloadForCRC := finalPayload + checksumPlaceholder()

	data := []byte(payloadForCRC)
	hexCode := crc16Hex(data)  // turn it into "3449" style text
	fullQR := string(finalPayload) + tlv(TagCRCchecksum, hexCode)

	return fullQR, nil
}