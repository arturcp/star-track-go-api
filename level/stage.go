package level

type Stage struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	GalaxyType string `json:"galaxy_type"`
}
