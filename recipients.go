package contacts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Recipient is a Contact added to the SendGrid API.
type Recipient struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt int    `json:"created_at"`
	// @TODO there are missing fields here

	CustomFields []CustomField `json:"-"` // @TODO custom unmarshal json for this
}

func (r *Recipient) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(*r)

	if err != nil {
		return nil, err
	}

	var fields map[string]interface{}

	err = json.Unmarshal(b, &fields)

	if err != nil {
		return nil, err
	}

	for _, f := range r.CustomFields {
		fields[f.Name] = f.Value
	}

	return json.Marshal(fields)
}

// RecipientClient defines methods for interacting with Recipients
type RecipientClient struct {
	*Client
}

// RecipientResponse is a response from operations dealing with Recipients
type RecipientResponse struct {
	ErrorCount          int      `json:"error_count"`
	ErrorIndices        []int    `json:"error_indices"`
	UnmodifiedIndices   []int    `json:"unmodified_indices"`
	NewCount            int      `json:"new_count"`
	PersistedRecipients []string `json:"persisted_recipients"`
	UpdatedCount        int      `json:"updated_count"`
	Errors              []struct {
		Message      string `json:"message"`
		ErrorIndices []int  `json:"error_indices"`
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (c *RecipientClient) attachIDs(resp *RecipientResponse, recipients []*Recipient) {
	if resp == nil {
		return
	}

	count := 0

	for index, recipient := range recipients {
		if contains(resp.ErrorIndices, index) || contains(resp.UnmodifiedIndices, index) {
			continue
		}

		recipient.ID = resp.PersistedRecipients[count]

		count++
	}
}

// Add multiple Recipients. Recipient IDs are attached to recipients upon success
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Add-Multiple-Recipients-POST
func (c *RecipientClient) Add(recipients ...*Recipient) (resp *RecipientResponse, err error) {
	err = c.makeRequest(http.MethodPost, "/contactdb/recipients", recipients, &resp)

	c.attachIDs(resp, recipients)

	return resp, err
}

// Update a Recipient.
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Update-Recipient-PATCH
func (c *RecipientClient) Update(recipients ...*Recipient) (resp *RecipientResponse, err error) {
	err = c.makeRequest(http.MethodPatch, "/contactdb/recipients", recipients, &resp)

	c.attachIDs(resp, recipients)

	return resp, err
}

// Delete one or more Recipients
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-Recipient-DELETE
func (c *RecipientClient) Delete(recipientIDs []string) error {
	return c.makeRequest(http.MethodDelete, "/contactdb/recipients", recipientIDs, nil)
}

type recipientListResponse struct {
	Recipients []*Recipient `json:"recipients"`
}

// List Recipients
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-Recipients-GET
func (c *RecipientClient) List(page int, pageSize int) ([]*Recipient, error) {
	var recipients recipientListResponse

	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/recipients?page=%d&page_size=%d", page, pageSize), nil, &recipients)

	if err != nil {
		return nil, err
	}

	return recipients.Recipients, nil
}

// Get (Retrieve) a Recipient
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Retrieve-a-Recipient-GET
func (c *RecipientClient) Get(recipientID string) (*Recipient, error) {
	var recipient *Recipient

	err := c.makeRequest(http.MethodGet, "/contactdb/recipients/"+recipientID, nil, &recipient)

	if err != nil {
		return nil, err
	}

	return recipient, nil
}

// ListsForRecipient gets the Lists that a Recipient is on.
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Get-the-Lists-the-Recipient-Is-On-GET
func (c *RecipientClient) ListsForRecipient(recipientID string) ([]List, error) {
	var lists []List

	err := c.makeRequest(http.MethodGet, "/contactdb/recipients/"+recipientID+"/lists", nil, &lists)

	if err != nil {
		return nil, err
	}

	return lists, err
}

type recipientCountResponse struct {
	RecipientCount int `json:"recipient_count"`
}

// BillableCount gets the count of billable Recipients
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Get-a-Count-of-Billable-Recipients-GET
func (c *RecipientClient) BillableCount() (int, error) {
	var recipientCount recipientCountResponse

	err := c.makeRequest(http.MethodGet, "/contactdb/recipients/billable_count", nil, &recipientCount)

	if err != nil {
		return -1, err
	}

	return recipientCount.RecipientCount, nil
}

// Count gets the count of recipients
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Get-a-Count-of-Recipients-GET
func (c *RecipientClient) Count() (int, error) {
	var recipientCount recipientCountResponse

	err := c.makeRequest(http.MethodGet, "/contactdb/recipients/count", nil, &recipientCount)

	if err != nil {
		return -1, err
	}

	return recipientCount.RecipientCount, nil
}

type recipientSearch struct {
	ListID     int         `json:"list_id"`
	Conditions []Condition `json:"conditions"`
}

// SearchWithConditions searches with Conditions
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Search-with-conditions-POST
func (c *RecipientClient) SearchListWithConditions(listID int, conditions ...Condition) ([]*Recipient, error) {
	var recipients recipientListResponse

	err := c.makeRequest(http.MethodGet, "/contactdb/recipients/search", recipientSearch{ListID: listID, Conditions: conditions}, &recipients)

	if err != nil {
		return nil, err
	}

	return recipients.Recipients, nil
}

type SearchTerm struct {
	FieldName  string
	FieldValue string
}

// Search for recipients matching Search criteria
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Get-Recipients-Matching-Search-Criteria-GET
func (c *RecipientClient) Search(criteria ...SearchTerm) ([]*Recipient, error) {
	u, err := url.Parse("/contactdb/recipients/search")

	if err != nil {
		return nil, err
	}

	q := u.Query()

	for _, term := range criteria {
		q.Add(term.FieldName, url.QueryEscape(term.FieldValue))
	}

	u.RawQuery = q.Encode()

	var recipients recipientListResponse

	err = c.makeRequest(http.MethodGet, u.String(), nil, &recipients)

	if err != nil {
		return nil, err
	}

	return recipients.Recipients, err
}
