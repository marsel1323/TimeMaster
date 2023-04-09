package model

import "time"

type Task struct {
	ID            int    `json:"id"`
	GoalID        int    `json:"goal_id"`
	Description   string `json:"description"`
	IsComplete    bool
	DueDate       time.Time `json:"due_date"`
	DaysOfWeek    []time.Weekday
	ScheduledTime time.Time
	RepeatTimes   int
	Interval      time.Duration
}
