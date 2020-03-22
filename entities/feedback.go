package entities

type FeedbackType string
type AvailabilityValue string
type BusynessValue string

const (
	BusynessFeedbackType     FeedbackType = "busyness"
	AvailabilityFeedbackType              = "availability"

	Available   AvailabilityValue = "available"
	Unavailable                   = "unavailable"

	Empty   BusynessValue = "empty"
	Average               = "average"
	Busy                  = "busy"
)

var busynessValues = map[BusynessValue]int{
	Empty:   0,
	Average: 1,
	Busy:    2,
}

var availabilityValues = map[AvailabilityValue]int{
	Unavailable: 1,
	Available:   1,
}

type Feedback struct {
	ShopID      string       `json:"shop_id"`
	ItemGroupID string       `json:"item_group_id"`
	Type        FeedbackType `json:"type"`
	Value       string       `json:"value"`
}

func (f Feedback) LessThan(filter Filter) bool {
	switch {
	case f.Type == BusynessFeedbackType && filter.Type == BusynessFilter:
		return busynessValues[BusynessValue(f.Value)] > busynessValues[BusynessValue(filter.Value)]
	case f.Type == AvailabilityFeedbackType && filter.Type == AvailableFilter:
		return f.ItemGroupID == filter.Value && f.Value == Unavailable
	}
	return false
}

func (f Feedback) IsForFilter(filter Filter) bool {
	if f.Type == BusynessFeedbackType && filter.Type == BusynessFilter {
		return true
	}
	if f.Type == AvailabilityFeedbackType && filter.Type == AvailableFilter {
		return true
	}
	return false
}
