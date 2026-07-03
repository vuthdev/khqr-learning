package bakong

import (
	"time"
)

func (c *Client) PollUntilPaidOrTimeout(md5Hash string, timeout time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	interval := 1 * time.Second
	
	for time.Now().Before(deadline) {
		status, err := c.CheckTransactionByMD5(md5Hash)
		if err != nil {
			return "", err
		}

		if status == "PAID" {
			return "PAID", nil
		}

		time.Sleep(interval)
		interval += 1 * time.Second
	}

	return "TIMEOUT", nil
}
