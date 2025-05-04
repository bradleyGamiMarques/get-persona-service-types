package types

type P3RPersona struct {
	PersonaName  string `json:"personaName"`
	Arcana       string `json:"arcana"`
	PersonaLevel int    `json:"personaLevel"`
	Strength     int    `json:"strength"`
	Magic        int    `json:"magic"`
	Endurance    int    `json:"endurance"`
	Agility      int    `json:"agility"`
	Luck         int    `json:"luck"`
}
type P3RPersonaListItem struct {
	PersonaName  string `json:"personaName"`
	Arcana       string `json:"arcana"`
	PersonaLevel int    `json:"personaLevel"`
	IsDLC        bool   `json:"isDLC"`
}
type GetP3RPersonaByNameResponse P3RPersona
type GetP3RPersonasByArcanaResponse []P3RPersona
type GetP3RPersonasResponse []P3RPersonaListItem
type GetP3RPersonasRequest struct {
	Arcanas []string `json:"arcanas"`
}
