package contacts

import "testing"

func TestCustomFieldsClient_Create(t *testing.T) {
	cf := &CustomField{
		Name: "pet",
		Type: "text",
	}

	err := client.CustomFields().Create(cf)

	if err != nil {
		t.Error(err)
	}
}

func TestCustomFieldsClient_Get(t *testing.T) {
	cf := &CustomField{
		Name: "favourite_beer",
		Type: "text",
	}

	err := client.CustomFields().Create(cf)

	if err != nil {
		t.Error(err)
	}

	got, err := client.CustomFields().Get(cf.ID)

	if err != nil {
		t.Error(err)
	}

	if got.ID != cf.ID || got.Name != cf.Name || got.Type != cf.Type {
		t.Fail()
	}
}

func TestCustomFieldsClient_List(t *testing.T) {
	cf := &CustomField{
		Name: "favourite_book",
		Type: "text",
	}

	err := client.CustomFields().Create(cf)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fields, err := client.CustomFields().List()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(fields) == 0 {
		t.Fail()
	}
}

func TestCustomFieldsClient_Delete(t *testing.T) {
	cf := &CustomField{
		Name: "delete_field",
		Type: "text",
	}

	err := client.CustomFields().Create(cf)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = client.CustomFields().Delete(cf.ID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestCustomFieldsClient_ReservedFields(t *testing.T) {
	f, err := client.CustomFields().ReservedFields()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(f) == 0 {
		t.Fail()
	}
}
