package dto

type Subscription struct {
	ID            string `json:"id"`
	Interval      string `json:"interval"`
	IntervalCount int    `json:"interval_count"`
}
