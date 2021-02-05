package ranking

type Ranking struct {
	ID         int    `json:"id"`
	PlayerName string `json:"player_name"`
	Points     int    `json:"points"`
	Character  string `json:"character"`
}
