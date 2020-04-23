package model

type Candidate struct {
	ID string  `json:"_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Department string `json:"department"`
	University string `json:"university"`
	Experience bool `json:"experience"`
	Status string `json:"status"`
	MeetingCount int `json:"meeting_count"`
	NextMeeting string `json:"next_meeting"`
	Assignee string `json:"assignee"`
}

type Assignee struct {
	ID string  `json:"_id"`
	Name string `json:"name"`
	Department string `json:"department"`
}
