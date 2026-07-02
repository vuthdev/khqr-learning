package khqr

type KHQRTag string

const (
	TagPayloadFormatIndicator KHQRTag = "00"
	TagPointofInitiationMethod KHQRTag = "01"
	TagIndividual KHQRTag = "29"
	TagMerchantAccountInfo KHQRTag = "30"
	TagMerchantCategoryCode KHQRTag = "52"
	TagTransactionCurrency KHQRTag = "53"
	TagTransactionAmount KHQRTag = "54"
	TagCountryCode KHQRTag = "58"
	TagMerchantName KHQRTag = "59"
	TagMerchantCity KHQRTag = "60"
	TagAdditionalData KHQRTag = "62"
	TagCRCchecksum KHQRTag = "63"

	// Account Info Tags
	Tag2930BakongAccountId KHQRTag = "00"
	Tag2930MerchantId KHQRTag = "01"
	Tag2930AcquiringBankName KHQRTag = "02"

	// Additional Data Tags
	Tag62BillNumber KHQRTag = "01"
	Tag62MobileNumber KHQRTag = "02"
	Tag62StoreLabel KHQRTag = "03"
	Tag62TerminalLabel KHQRTag = "07"
)