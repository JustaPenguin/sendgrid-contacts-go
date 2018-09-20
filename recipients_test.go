package contacts

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestToRecipientID(t *testing.T) {
	out := ToRecipientID("foo@example.com")

	if out != "Zm9vQGV4YW1wbGUuY29t" {
		t.Fail()
	}
}

func TestRecipientClient_Add(t *testing.T) {
	t.Run("One recipient", func(t *testing.T) {
		resp, err := client.Recipients().Add(&Recipient{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"})

		if err != nil {
			t.Error(err)
		}

		if resp == nil || resp.ErrorCount > 0 {
			spew.Dump(resp)
			t.Fail()
		}

		/*t.Run("With custom fields", func(t *testing.T) {
			resp, err := client.Recipients().Add(Recipient{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", CustomFields: []CustomField{
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
		resp, err := client.Recipients().Add(
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

	_, err := client.Recipients().Add(r)

	if err != nil {
		t.Error(err)
	}

	r.Email = "jimothy.smith@example.com"

	resp, err := client.Recipients().Update(r)

	if err != nil {
		t.Error(err)
	}

	if resp == nil || resp.ErrorCount > 0 || resp.NewCount > 0 {
		t.Fail()
	}
}

func TestRecipientClient_Delete(t *testing.T) {
	r := &Recipient{FirstName: "Delete", LastName: "Test", Email: "delete.test@example.com"}

	_, err := client.Recipients().Add(r)

	if err != nil {
		t.Error(err)
	}

	err = client.Recipients().Delete([]string{r.ID})

	if err != nil {
		t.Error(err)
	}
}

func TestRecipientClient_List(t *testing.T) {
	recipients, err := client.Recipients().List(1, 100)

	if err != nil {
		t.Error(err)
	}

	if len(recipients) <= 0 {
		t.Fail()
	}
}

func TestRecipientClient_Get(t *testing.T) {
	r := &Recipient{FirstName: "Get", LastName: "Test", Email: "get.test@example.com"}

	_, err := client.Recipients().Add(r)

	if err != nil {
		t.Error(err)
	}

	got, err := client.Recipients().Get(r.ID)

	if err != nil {
		t.Error(err)
	}

	if got == nil || got.Email != r.Email {
		t.Fail()
	}

	err = client.Recipients().Delete([]string{r.ID})

	if err != nil {
		t.Error(err)
	}
}

func TestRecipientClient_ListsForRecipient(t *testing.T) {
	// @TODO.
}

func TestRecipientClient_BillableCount(t *testing.T) {
	billableCount, err := client.Recipients().BillableCount()

	if err != nil {
		t.Error(err)
	}

	if billableCount < 0 {
		t.Fail()
	}
}

func TestRecipientClient_Count(t *testing.T) {
	count, err := client.Recipients().Count()

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
	recipients, err := client.Recipients().Search(SearchTerm{FieldName: "email", FieldValue: "example.com"})

	if err != nil {
		t.Error(err)
	}

	if len(recipients) == 0 {
		t.Fail()
	}
}
