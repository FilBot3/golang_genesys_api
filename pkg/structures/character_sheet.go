package structures

// CharacterSheet is the overall structure of how the Character will look.
type CharacterSheet struct {
	Name             string             `json:"name"`
	Archetype        string             `json:"archetype"`
	Career           string             `json:"career"`
	Player           string             `json:"player"`
	Money            string             `json:"money"`
	Portrait         string             `json:"portrait"`
	Description      Description        `json:"description"`
	Characteristics  Characteristics    `json:"characteristics"`
	FiguredStats     FiguredStats       `json:"figured_stats"`
	Experience       Experience         `json:"experience"`
	Skills           Skills             `json:"skills"`
	Talents          Talents            `json:"talents"`
	SpecialAbilities []SpecialAbilities `json:"special_abilities"`
	Weapons          []Weapons          `json:"weapons"`
	Items            []Items            `json:"items"`
	Motivations      Motivations        `json:"motivations"`
	CriticalInjuries []CriticalInjuries `json:"critical_injuries"`
	Notes            string             `json:"notes"`
}

// Description is the metadata section of the character.
type Description struct {
	Gender          string `json:"gender"`
	Age             string `json:"age"`
	Height          string `json:"height"`
	Build           string `json:"build"`
	Hair            string `json:"hair"`
	Eyes            string `json:"eyes"`
	NotableFeatures string `json:"notable_features"`
}

// Characteristics are the raw stats of the character.
type Characteristics struct {
	Brawn     string `json:"brawn"`
	Agility   string `json:"agility"`
	Intellect string `json:"intellect"`
	Cunning   string `json:"cunning"`
	Willpower string `json:"willpower"`
	Presence  string `json:"presence"`
}

// Defense is the number of setback dice that are added to each attack action
// against this character.
type Defense struct {
	Ranged string `json:"ranged"`
	Melee  string `json:"melee"`
}

// FiguredStats take the characteristics, plus some information about the
// character's archetype.
type FiguredStats struct {
	Soak    string  `json:"soak"`
	Wounds  string  `json:"wounds"`
	Strain  string  `json:"strain"`
	Defense Defense `json:"defense"`
}

// Experience is the number of experience points the character as earned and
// saved.
type Experience struct {
	Total     string `json:"total"`
	Available string `json:"available"`
}

// Skill is the rough outline of what each skill is.
type Skill struct {
	Name   string `json:"name"`
	Career string `json:"career"`
	Rank   string `json:"rank"`
}

// Skills are the largest part of a character next to their talents and
// motivations.
type Skills struct {
	General   []Skill `json:"general"`
	Magic     []Skill `json:"magic"`
	Combat    []Skill `json:"combat"`
	Social    []Skill `json:"social"`
	Knowledge []Skill `json:"knowledge"`
}

// Talent is the next largest thing a character has. These change the game.
type Talent struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Type        string `json:"type"`
	Activation  string `json:"activation"`
	Description string `json:"description"`
}

// Talents is a structure of Talent structures that list the abilities the
// Character can do.
type Talents struct {
	Tier1 []Talent `json:"tier_1"`
	Tier2 []Talent `json:"tier_2"`
	Tier3 []Talent `json:"tier_3"`
	Tier4 []Talent `json:"tier_4"`
	Tier5 []Talent `json:"tier_5"`
}

// SpecialAbilities are similar to Talents, but they're more specialized and
// only certain archetypes have them.
type SpecialAbilities struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
}

// Weapons are the primary device for overcoming something, next to skills.
type Weapons struct {
	Name     string `json:"name"`
	Skill    string `json:"skill"`
	Damage   string `json:"damage"`
	Critical string `json:"critical"`
	Range    string `json:"range"`
	Special  string `json:"special"`
}

// Items are things that give the character the ability to do other things???
type Items struct {
	Name         string `json:"name"`
	Source       string `json:"source"`
	Price        string `json:"price"`
	Encumberance string `json:"encumberance"`
	Rarity       string `json:"rarity"`
	Description  string `json:"description"`
}

// Motivations are a role-player's queue and direction for how to play their
// character.
type Motivations struct {
	Strength string `json:"strength"`
	Flaw     string `json:"flaw"`
	Desire   string `json:"desire"`
	Fear     string `json:"fear"`
}

// CriticalInjuries are things that your character suffers when a weapon's crit
// rating is triggered.
type CriticalInjuries struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
