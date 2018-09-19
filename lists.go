package contacts

import (
	"fmt"
	"net/http"
)

// A List is a group of Recipients
type List struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name"`
	RecipientCount int    `json:"recipient_count,omitempty"`
}

// ListsClient provides methods for interacting with Lists.
type ListsClient struct {
	client *Client
}

// Create a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Create-a-List-POST
func (c *ListsClient) Create(name string) (*List, error) {
	list := &List{Name: name}

	err := c.client.makeRequest(http.MethodPost, "/contactdb/lists", list, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

type listListsResponse struct {
	Lists []*List `json:"lists"`
}

// List all Lists
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-All-Lists-GET
func (c *ListsClient) List() ([]*List, error) {
	var resp *listListsResponse

	err := c.client.makeRequest(http.MethodGet, "/contactdb/lists", nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.Lists, nil
}

// Delete multiple Lists
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-Multiple-lists-DELETE
func (c *ListsClient) Delete(listIDs []uint) error {
	return c.client.makeRequest(http.MethodDelete, "/contactdb/lists", listIDs, nil)
}

// Get (Retrieve) a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Retrieve-a-List-GET
func (c *ListsClient) Get(listID uint) (*List, error) {
	var list *List

	err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/lists/%d", listID), nil, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Update a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Update-a-List-PATCH
func (c *ListsClient) Update(list *List) error {
	return c.client.makeRequest(http.MethodPatch, fmt.Sprintf("/contactdb/lists/%d", list.ID), list, nil)
}

// ListRecipients on a given List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-Recipients-on-a-List-GET
func (c *ListsClient) ListRecipients(listID, pageSize, pageNum uint) ([]*Recipient, error) {
	var resp *listRecipientsResponse

	err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/lists/%d/recipients?page_size=%d&page=%d", listID, pageSize, pageNum), nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.Recipients, nil
}

// AddRecipients to a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Add-Multiple-Recipients-to-a-List-POST
func (c *ListsClient) AddRecipients(listID uint, recipients []*Recipient) error {
	var recipientIDs []string

	for _, recipient := range recipients {
		recipientIDs = append(recipientIDs, recipient.ID)
	}

	return c.AddRecipientsByIDs(listID, recipientIDs)
}

// AddRecipientsByIDs to a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Add-Multiple-Recipients-to-a-List-POST
func (c *ListsClient) AddRecipientsByIDs(listID uint, recipientIDs []string) error {
	return c.client.makeRequest(http.MethodPost, fmt.Sprintf("/contactdb/lists/%d/recipients", listID), recipientIDs, nil)
}

// DeleteRecipient from a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-a-Single-Recipient-from-a-Single-List-DELETE
func (c *ListsClient) DeleteRecipient(listID uint, recipient *Recipient) error {
	return c.DeleteRecipientByID(listID, recipient.ID)
}

// DeleteRecipientByID from a List
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-a-Single-Recipient-from-a-Single-List-DELETE
func (c *ListsClient) DeleteRecipientByID(listID uint, recipientID string) error {
	return c.client.makeRequest(http.MethodDelete, fmt.Sprintf("/contactdb/lists/%d/recipients/%s", listID, recipientID), nil, nil)
}
