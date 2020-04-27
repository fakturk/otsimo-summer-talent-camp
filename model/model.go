package model

import "time"

type Candidate struct {
	ID string `bson:"_id" json:"id"`
	First_Name string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Email string `json:"email"`
	Department string `json:"department"`
	University string `json:"university"`
	Experience bool `json:"experience"`
	Status string `json:"status"`
	Meeting_Count int `json:"meeting_count"`
	Next_Meeting *time.Time `json:"next_meeting"`
	Assignee string `json:"assignee"`
}

type Assignee struct {
	ID string  `json:"_id"`
	Name string `json:"name"`
	Department string `json:"department"`
}

type MeetingTime struct {
	MeetingTime *time.Time `json:"meeting_time"`
}
