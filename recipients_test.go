package contacts

import (
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/ratelimit"
)

var recipientClient *RecipientClient

func init() {
	recipientClient = New(os.Getenv("SENDGRID_APIKEY")).Recipients()
	recipientClient.HTTPClient = &http.Client{
		Transport: &RateLimitingTransport{},
	}
}

type RateLimitingTransport struct{}

var rateLimiter = ratelimit.New(1)

func (x *RateLimitingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	rateLimiter.Take()

	return http.DefaultTransport.RoundTrip(r)
}

func TestRecipientClient_Add(t *testing.T) {
	t.Run("One recipient", func(t *testing.T) {
		resp, err := recipientClient.Add(&Recipient{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"})

		if err != nil {
			t.Error(err)
		}

		if resp == nil || resp.ErrorCount > 0 {
			spew.Dump(resp)
			t.Fail()
		}

		/*t.Run("With custom fields", func(t *testing.T) {
			resp, err := recipientClient.Add(Recipient{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", CustomFields: []CustomField{
				{
					Name:  "favourite_beer",
					Value: "Budweiser",
				},
			}})

			if err != nil {
				t.Error(err)
			}

			if resp == nil || resp.ErrorCount > 0 {
				spew.Dump(resp)
				t.Fail()
			}
		})*/
	})

	t.Run("Multiple recipients", func(t *testing.T) {
		resp, err := recipientClient.Add(
			&Recipient{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
			&Recipient{FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"},
			&Recipient{FirstName: "Tim", LastName: "Smith", Email: "tim.smith@example.me"},
			&Recipient{FirstName: "Sally", LastName: "Davis", Email: "sally.davis@something.com"},
		)

		if err != nil {
			t.Error(err)
		}

		if resp == nil || resp.ErrorCount > 0 {
			spew.Dump(resp)
			t.Fail()
		}
	})
}

func TestRecipientClient_Update(t *testing.T) {
	r := &Recipient{FirstName: "Update", LastName: "Test", Email: "jimmy.smith@example.com"}

	_, err := recipientClient.Add(r)

	if err != nil {
		t.Error(err)
	}

	r.Email = "jimothy.smith@example.com"

	resp, err := recipientClient.Update(r)

	if err != nil {
		t.Error(err)
	}

	if resp == nil || resp.ErrorCount > 0 || resp.NewCount > 0 {
		t.Fail()
	}
}

func TestRecipientClient_Delete(t *testing.T) {
	r := &Recipient{FirstName: "Delete", LastName: "Test", Email: "delete.test@example.com"}

	_, err := recipientClient.Add(r)

	if err != nil {
		t.Error(err)
	}

	err = recipientClient.Delete([]string{r.ID})

	if err != nil {
		t.Error(err)
	}
}

func TestRecipientClient_List(t *testing.T) {
	recipients, err := recipientClient.List(1, 100)

	if err != nil {
		t.Error(err)
	}

	if len(recipients) <= 0 {
		t.Fail()
	}
}

func TestRecipientClient_Get(t *testing.T) {
	r := &Recipient{FirstName: "Get", LastName: "Test", Email: "get.test@example.com"}

	_, err := recipientClient.Add(r)

	if err != nil {
		t.Error(err)
	}

	got, err := recipientClient.Get(r.ID)

	if err != nil {
		t.Error(err)
	}

	if got == nil || got.Email != r.Email {
		t.Fail()
	}

	err = recipientClient.Delete([]string{r.ID})

	if err != nil {
		t.Error(err)
	}
}

func TestRecipientClient_ListsForRecipient(t *testing.T) {
	// @TODO.
}

func TestRecipientClient_BillableCount(t *testing.T) {
	billableCount, err := recipientClient.BillableCount()

	if err != nil {
		t.Error(err)
	}

	if billableCount < 0 {
		t.Fail()
	}
}

func TestRecipientClient_Count(t *testing.T) {
	count, err := recipientClient.Count()

	if err != nil {
		t.Error(err)
	}

	if count < 0 {
		t.Fail()
	}
}

func TestRecipientClient_SearchListWithConditions(t *testing.T) {
	// @TODO
}

func TestRecipientClient_Search(t *testing.T) {
	recipients, err := recipientClient.Search(SearchTerm{FieldName: "email", FieldValue: "example.com"})

	if err != nil {
		t.Error(err)
	}

	if len(recipients) == 0 {
		t.Fail()
	}

	spew.Dump(recipients)
}
