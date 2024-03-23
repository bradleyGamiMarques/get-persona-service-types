package types

type GetP3RPersonaByNameResponse struct {
	Agility      int    `json:"agility"`
	Arcana       string `json:"arcana"`
	Background   string `json:"background"`
	Endurance    int    `json:"endurance"`
	Luck         int    `json:"luck"`
	Magic        int    `json:"magic"`
	PersonaName  string `json:"personaname"`
	PersonaLevel int    `json:"personalevel"`
	Strength     int    `json:"strength"`
}
