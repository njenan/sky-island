package skyisland

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Data
type Data struct {
	Timestamp int64
	Data      interface{}
}

// Client
type Client struct {
	hc       *http.Client
	endpoint string
}

// NewClient creates a new client usable with the Sky Island API
func NewClient(url string, port int, timeout time.Duration) *Client {
	c := &Client{
		hc: &http.Client{
			Timeout: timeout,
		},
		endpoint: fmt.Sprintf("%s:%d/api/v1/function", url, port),
	}
	return c
}

// Function makes the call to the API
func (c *Client) Function(url, call string) (*Data, error) {
	d := fmt.Sprintf(`{"url": "%s", "call": "%s"}`, url, call)
	req, err := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewBuffer([]byte(d)))
	if err != nil {
		return nil, err
	}
	res, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data Data
	if err := json.NewDecoder(res.Body).Decode(res); err != nil {
		return nil, err
	}
	return &data, nil
}
