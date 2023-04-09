package model

type Goal struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	Specific    string `json:"specific"`
	Measurable  string `json:"measurable"`
	Achievable  string `json:"achievable"`
	Relevant    string `json:"relevant"`
	TimeBound   string `json:"time_bound"`
}
