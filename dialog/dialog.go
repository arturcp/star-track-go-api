package dialog

type Dialog struct {
	CharacterType string   `json:"character_type"`
	NpcID         int      `json:npc_id`
	Paragraphs    []string `json:paragraphs`
}
