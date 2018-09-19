package contacts

import (
	"testing"
	"time"
)

func TestListsClient_Create(t *testing.T) {
	list, err := client.Lists().Create("test_list")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if list == nil || list.Name != "test_list" {
		t.Fail()
	}
}

func TestListsClient_Delete(t *testing.T) {
	list, err := client.Lists().Create("delete_list")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list2, err := client.Lists().Create("delete_list2")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = client.Lists().Delete(list.ID, list2.ID)

	if err != nil {
		t.Fail()
	}
}

func TestListsClient_Get(t *testing.T) {
	list, err := client.Lists().Create("get_list")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	got, err := client.Lists().Get(list.ID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if got == nil || got.ID != list.ID || got.Name != list.Name {
		t.Fail()
	}
}

func TestListsClient_Update(t *testing.T) {
	list, err := client.Lists().Create("update_list")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Name = "new_name_list"

	err = client.Lists().Update(list)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	got, err := client.Lists().Get(list.ID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if got == nil || got.Name != "new_name_list" {
		t.Fail()
	}
}

func TestListsClient_List(t *testing.T) {
	_, err := client.Lists().Create("list_list")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	lists, err := client.Lists().List()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(lists) == 0 {
		t.Fail()
	}

	found := false

	for _, list := range lists {
		if list.Name == "list_list" {
			found = true
		}
	}

	if !found {
		t.Fail()
	}
}

func TestListsClient_Recipients(t *testing.T) {
	r := &Recipient{Email: "foo@bar.com"}

	_, err := client.Recipients().Add(r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list, err := client.Lists().Create("with_recipients")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = client.Lists().AddRecipients(list.ID, r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// adding recipients to lists takes some time it seems...
	time.Sleep(time.Second * 5)

	recipients, err := client.Lists().ListRecipients(list.ID, 1, 100)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	found := false

	for _, recipient := range recipients {
		if recipient.ID == r.ID {
			found = true
		}
	}

	if !found {
		t.Fail()
	}

	err = client.Lists().DeleteRecipient(list.ID, r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	recipients, err = client.Lists().ListRecipients(list.ID, 1, 100)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, recipient := range recipients {
		if recipient.ID == r.ID {
			t.Fail()
		}
	}
}
