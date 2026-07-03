package bakong

import (
	"bytes"
	"fmt"
	"io"
	"khqr-learn/internal/constant"
	"khqr-learn/internal/constant/utils"
	"net/http"
	"time"
)

const (
	account_id_enpoints = "/v1/check_bakong_account"
	check_transaction_by_md5_endpoints = "/v1/check_transaction_by_md5"
)

type Client struct {
	Token   string
	BaseURL string
}

type BakongResponse struct {
	ResponseCode *int `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrorCode *int `json:"errorCode"`
	Data interface{} `json:"data"`
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
		BaseURL: constant.BaseURL,
	}
}

// {BaseURL}/v1/check_transaction_by_md5
type md5Payload struct {
	Md5 string `json:"md5"`
}

func (c *Client) CheckTransactionByMD5(md5Hash string) (string, error) {
	payload := md5Payload { Md5: md5Hash }

	json_data := utils.ToJson(payload)
	client := http.Client {
		Timeout: 10 * time.Second,
	}

	fmt.Printf("Json data: %v", json_data)

	req, err := http.NewRequest("POST", c.BaseURL + check_transaction_by_md5_endpoints, bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		return "ERROR", fmt.Errorf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer " + c.Token)
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Body: %v", req.Body)
	res, err := client.Do(req)
	if err != nil {
		return "ERROR", fmt.Errorf("Error sending request: %v\n", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "ERROR", fmt.Errorf("Error reading response: %v\n", err)
	}

	var apiResponse BakongResponse

	_, err = utils.JsonUnmarshal(body, &apiResponse)
	if err != nil {
		return "ERROR", fmt.Errorf("Error decoding response")
	}

	if apiResponse.ResponseCode == nil {
		return "ERROR", fmt.Errorf("response missing responseCode")
	}

	data := apiResponse.Data.(map[string]any)

	if *apiResponse.ResponseCode == 0 && data["hash"] != nil {
		return "PAID", nil
	} else {
		return "UNPAID", err
	}
}

// Endpoint: {BaseURL}/v1/check_account_by_id
type accountIDPayload struct {
	AccountID string `json:"accountId"`
}

func (c *Client) CheckAccountByID(accountID string) (bool, error) {
	payload := accountIDPayload{ AccountID: accountID }

	json_data := utils.ToJson(payload)
	client := http.Client {
		Timeout: 10 * time.Second,
	}

	fmt.Printf("Json data: %v\n", json_data)

	req, err := http.NewRequest("POST", c.BaseURL + account_id_enpoints, bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		return false, fmt.Errorf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer " + c.Token)
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Body: %v\n", req.Body)
	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("Error sending request: %v\n", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("Error reading response: %v\n", err)
	}

	var apiResponse BakongResponse

	_, err = utils.JsonUnmarshal(body, &apiResponse)
	if err != nil {
		return false, fmt.Errorf("Error decoding response")
	}

	if apiResponse.ResponseCode == nil {
		return false, fmt.Errorf("response missing responseCode")
	}

	if (*apiResponse.ResponseCode == 0 && apiResponse.ResponseMessage == "Account ID exists") {
		return true, nil
	} else {
		return false, err
	}
}
