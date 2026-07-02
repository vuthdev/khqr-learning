package bakong

import "time"

// PollUntilPaidOrTimeout repeatedly checks a transaction's status until
// it's PAID or the timeout is reached. This mirrors the pattern from the
// Python bakong-khqr SDK docs: don't hammer the API every second for the
// full duration, back off over time.
//
// TODO(you): implement a polling loop that:
//  1. Tracks elapsed time since start.
//  2. Increases the sleep interval the longer it's been waiting
//     (e.g. 1s for the first 10s, then 3s, then 5s+ — design your own
//     backoff matrix, there's no single "correct" answer here).
//  3. Stops and returns "TIMEOUT" once timeout is exceeded.
//  4. Stops and returns "PAID" as soon as status flips.
//
// This is good practice for later: your Favorite Schedule Quartz jobs
// have similar "don't hammer a resource" concerns.
func (c *Client) PollUntilPaidOrTimeout(md5Hash string, timeout time.Duration) (string, error) {
	panic("not implemented")
}
