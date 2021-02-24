package character

type Character struct {
	ID        int      `json:"id"`
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	ImageURL  string   `json:"image_url"`
	Bio       string   `json:"bio"`
	MainColor string   `json:"main_color"`
	Pronouns  string   `json:"pronouns"`
	History   []string `json:"history"`
}
