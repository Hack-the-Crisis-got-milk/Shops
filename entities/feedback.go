package entities

type FeedbackType string

const (
	BusynessFeedbackType     FeedbackType = "busyness"
	AvailabilityFeedbackType              = "availability"
)

type Feedback struct {
	ShopID      string       `json:"shop_id"`
	ItemGroupID string       `json:"item_group_id"`
	Type        FeedbackType `json:"type"`
	Value       string       `json:"value"`
}

