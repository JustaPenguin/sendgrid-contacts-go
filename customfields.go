package contacts

import (
	"fmt"
	"net/http"
)

// CustomField is a field which can be added to a Recipient
type CustomField struct {
	ID    uint        `json:"id,omitempty"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"-"`
}

// CustomFieldsClient provides methods for managing CustomFields.
type CustomFieldsClient struct {
	client *Client
}

// Create a Custom Field
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Create-a-Custom-Field-POST
func (c *CustomFieldsClient) Create(field *CustomField) error {
	return c.client.makeRequest(http.MethodPost, "/contactdb/custom_fields", field, &field)
}

type customFieldListResponse struct {
	CustomFields []*CustomField `json:"custom_fields"`
}

// List all Custom Fields
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-All-Custom-Fields-GET
func (c *CustomFieldsClient) List() ([]*CustomField, error) {
	var resp *customFieldListResponse

	err := c.client.makeRequest(http.MethodGet, "/contactdb/custom_fields", nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.CustomFields, nil
}

// Get (Retrieve) a Custom Field
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Retrieve-a-Custom-Field-GET
func (c *CustomFieldsClient) Get(customFieldID uint) (*CustomField, error) {
	var field *CustomField

	err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/custom_fields/%d", customFieldID), nil, &field)

	if err != nil {
		return nil, err
	}

	return field, nil
}

// Delete a Custom Field
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-a-Custom-Field-DELETE
func (c *CustomFieldsClient) Delete(customFieldID uint) error {
	return c.client.makeRequest(http.MethodDelete, fmt.Sprintf("/contactdb/custom_fields/%d", customFieldID), nil, nil)
}

type reservedFieldsResponse struct {
	ReservedFields []*CustomField `json:"reserved_fields"`
}

// ReservedFields returns all fields which are reserved
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#-Reserved-Fields
func (c *CustomFieldsClient) ReservedFields() ([]*CustomField, error) {
	var resp *reservedFieldsResponse

	err := c.client.makeRequest(http.MethodGet, "/contactdb/reserved_fields", nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.ReservedFields, nil
}
