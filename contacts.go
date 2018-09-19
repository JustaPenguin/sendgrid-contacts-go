package contacts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const sendgridAPIv3Base = "https://api.sendgrid.com/v3"

func New(apikey string) *Client {
	if apikey == "" {
		panic(errors.New("contacts: apikey must be set"))
	}

	return &Client{
		APIKey:     apikey,
		HTTPClient: http.DefaultClient,
	}
}

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func (c *Client) makeRequest(method, url string, data, output interface{}) error {
	var body io.Reader

	if method != http.MethodGet && data != nil {
		var err error
		body, err = c.marshal(data)

		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, sendgridAPIv3Base+url, body)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if output != nil && resp.StatusCode < http.StatusBadRequest {
		err = c.unmarshal(resp.Body, output)

		return err
	} else if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("contacts: bad status code observed: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) marshal(data interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Client) unmarshal(r io.Reader, into interface{}) error {
	return json.NewDecoder(r).Decode(into)
}

func (c *Client) Recipients() *RecipientClient {
	return &RecipientClient{client: c}
}

func (c *Client) Lists() *ListsClient {
	return &ListsClient{client: c}
}

func (c *Client) Segments() *SegmentsClient {
	return &SegmentsClient{client: c}
}

func (c *Client) CustomFields() *CustomFieldsClient {
	return &CustomFieldsClient{client: c}
}
