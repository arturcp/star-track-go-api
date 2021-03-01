package level

type Level struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Stages []Stage `json:"stages"`
}
