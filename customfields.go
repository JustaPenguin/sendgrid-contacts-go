package contacts

type CustomField struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value interface{}
}

type CustomFieldsClient struct {
	*Client
}
