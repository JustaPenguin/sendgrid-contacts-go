package contacts

type List struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	RecipientCount int    `json:"recipient_count"`
}

type ListsClient struct {
	*Client
}
