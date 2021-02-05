package character

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Bio      string `json:"bio"`
	Pronouns string `json:"pronouns"`
}
