# sendgrid-contacts-go
Golang bindings for SendGrid's [Contacts Marketing API](https://sendgrid.com/docs/API_Reference/Web_API_v3/Marketing_Campaigns/contactdb.html)

## Example

```go

import "github.com/justapenguin/sendgrid-contacts-go"

// ...

client := contacts.New("SENDGRID_APIKEY")

resp, err := client.Recipients().Add(
    &Recipient{
      FirstName: "John", 
      LastName:  "Doe", 
      Email:     "john.doe@example.com",
    },
)

if err != nil {
    // ...
}

// do something with resp...

```
