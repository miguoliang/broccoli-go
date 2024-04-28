package dto

import "time"

type Subscription struct {
	ID            string `json:"id"`
	Interval      string `json:"interval"`
	IntervalCount int    `json:"interval_count"`
}

type GetUsageRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required,gt|gtfield=StartDate" time_format:"2006-01-02"`
}
