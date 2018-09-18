package contacts

type Segment struct {
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
