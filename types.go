package types

type GetP3RPersonaByNameResponse struct {
	Agility      int    `json:"agility"`
	Arcana       string `json:"arcana"`
	Background   string `json:"background"`
	Endurance    int    `json:"endurance"`
	Luck         int    `json:"luck"`
	Magic        int    `json:"magic"`
	PersonaName  string `json:"personaName"`
	PersonaLevel int    `json:"personaLevel"`
	Strength     int    `json:"strength"`
}

type GetP3RPersonasByArcanaResponse []struct {
	Agility      int    `json:"agility"`
	Arcana       string `json:"arcana"`
	Background   string `json:"background"`
	Endurance    int    `json:"endurance"`
	Luck         int    `json:"luck"`
	Magic        int    `json:"magic"`
	PersonaName  string `json:"personaName"`
	PersonaLevel int    `json:"personaLevel"`
	Strength     int    `json:"strength"`
}
