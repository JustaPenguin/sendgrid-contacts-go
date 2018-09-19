package contacts

import (
	"fmt"
	"net/http"
)

type Segment struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	ListID     int         `json:"list_id"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
	AndOr    string `json:"and_or"`
}

type SegmentsClient struct {
	*Client
}

// Create a Segment
//
// The response of the initial Create will return a recipient_count of 0 because it takes some time to populate.
// Follow up with a Get to verify segment size.
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Create-a-Segment-POST
func (c *SegmentsClient) Create(segment *Segment) error {
	return c.makeRequest(http.MethodPost, "/contactdb/segments", segment, &segment)
}

type listSegmentsResponse struct {
	Segments []*Segment `json:"segments"`
}

// List all Segments
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-All-Segments-GET
func (c *SegmentsClient) List() ([]*Segment, error) {
	var resp *listSegmentsResponse

	err := c.makeRequest(http.MethodGet, "/contactdb/segments", nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.Segments, nil
}

// Get (Retrieve) a Segment
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Retrieve-a-Segment-GET
func (c *SegmentsClient) Get(segmentID uint) (*Segment, error) {
	var segment *Segment

	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/segments/%d", segmentID), nil, &segment)

	if err != nil {
		return nil, err
	}

	return segment, nil
}

// Update a Segment
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Update-a-Segment-PATCH
func (c *SegmentsClient) Update(segment *Segment) error {
	return c.makeRequest(http.MethodPatch, fmt.Sprintf("/contactdb/segments/%d", segment.ID), segment, &segment)
}

// Delete a Segment
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#Delete-a-Segment-DELETE
func (c *SegmentsClient) Delete(segmentID uint) error {
	return c.makeRequest(http.MethodDelete, fmt.Sprintf("/contactdb/segments/%d", segmentID), nil, nil)
}

// ListRecipients on a Segment
//
// https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html#List-Recipients-On-a-Segment-GET
func (c *SegmentsClient) ListRecipients(segmentID, pageSize, page uint) ([]*Recipient, error) {
	var resp *listRecipientsResponse

	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/contactdb/segments/%d/recipients?page_size=%d&page=%d", segmentID, pageSize, page), nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp.Recipients, nil
}
