package types

type P3RPersona struct {
	Agility      int    `json:"agility"`
	Arcana       string `json:"arcana"`
	Background   string `json:"background"`
	Endurance    int    `json:"endurance"`
	IsDLC        *bool  `json:"isDLC,omitempty"`
	Luck         int    `json:"luck"`
	Magic        int    `json:"magic"`
	PersonaName  string `json:"personaName"`
	PersonaLevel int    `json:"personaLevel"`
	Strength     int    `json:"strength"`
}
type GetP3RPersonaByNameResponse P3RPersona
type GetP3RPersonasByArcanaResponse []P3RPersona
