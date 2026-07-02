package bakong

// Client wraps calls to the Bakong Open API sandbox (SIT) environment.
// Base URL for SIT: https://sit-api-bakong.nbc.gov.kh
// Base URL for production: https://api-bakong.nbc.gov.kh
type Client struct {
	Token   string // your developer token from api-bakong.nbc.gov.kh/register/
	BaseURL string
}

// NewClient creates a client pointed at the sandbox by default.
// TODO(you): fill in BaseURL default.
func NewClient(token string) *Client {
	panic("not implemented")
}

// CheckTransactionByMD5 asks Bakong whether the transaction matching
// this MD5 hash has been paid.
//
// TODO(you):
//  1. Build a POST request to {BaseURL}/v1/check_transaction_by_md5
//  2. Set header: Authorization: Bearer {Token}
//  3. Body: {"md5": "<hash>"}
//  4. Parse the JSON response — look at the SDK docs/PDF for the exact
//     shape (responseCode, responseMessage, data.status or similar).
//
// Return a simple string status ("PAID", "UNPAID", "ERROR") to start —
// you can build a richer struct once you see the real response shape.
func (c *Client) CheckTransactionByMD5(md5Hash string) (string, error) {
	panic("not implemented")
}

// CheckAccountByID verifies a Bakong account ID exists/is valid.
// Useful sanity check before generating a QR for an account.
// Endpoint: {BaseURL}/v1/check_account_by_id
//
// TODO(you): implement similarly to CheckTransactionByMD5.
func (c *Client) CheckAccountByID(accountID string) (bool, error) {
	panic("not implemented")
}
