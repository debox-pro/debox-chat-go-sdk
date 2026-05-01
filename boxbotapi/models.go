package boxbotapi

type SightMsg struct {
	Name     string `json:"name"`
	SightURI string `json:"sightUrl"`
	Content  string `json:"content"`
	Duration string `json:"duration"`
	Size     string `json:"size"`
}
