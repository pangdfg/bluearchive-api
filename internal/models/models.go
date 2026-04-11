package models

// Generic response wrapper

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Upstream root response

type UpstreamRoot struct {
	Status    int      `json:"status"`
	Uptime    int64    `json:"uptime"`
	Endpoints []string `json:"endpoints"`
}

// Character models

type CharacterListResponse struct {
	Status int         `json:"status"`
	Data   []Character `json:"data"`
}

type CharacterResponse struct {
	Status int       `json:"status"`
	Data   Character `json:"data"`
}

type Character struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	ProfileImage string       `json:"profile_image,omitempty"`
	Rarity       string       `json:"rarity,omitempty"`
	School       string       `json:"school,omitempty"`
	Club         string       `json:"club,omitempty"`
	Role         string       `json:"role,omitempty"`
	ArmorType    string       `json:"armor_type,omitempty"`
	BulletType   string       `json:"bullet_type,omitempty"`
	Position     string       `json:"position,omitempty"`
	Weapon       string       `json:"weapon,omitempty"`
	Skills       []Skill      `json:"skills,omitempty"`
	Stats        *CharStats   `json:"stats,omitempty"`
	Info         *CharInfo    `json:"info,omitempty"`
}

type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon,omitempty"`
	Type        string `json:"type,omitempty"`
}

type CharStats struct {
	Attack       int `json:"attack,omitempty"`
	Defense      int `json:"defense,omitempty"`
	HP           int `json:"hp,omitempty"`
	HealPower    int `json:"heal_power,omitempty"`
	AccuracyPoint int `json:"accuracy_point,omitempty"`
	DodgePoint   int `json:"dodge_point,omitempty"`
	CriticalPoint int `json:"critical_point,omitempty"`
	StabilityPoint int `json:"stability_point,omitempty"`
	Range        int `json:"range,omitempty"`
	MoveSpeed    int `json:"move_speed,omitempty"`
	SightPoint   int `json:"sight_point,omitempty"`
}

type CharInfo struct {
	Age          string `json:"age,omitempty"`
	Birthday     string `json:"birthday,omitempty"`
	Height       string `json:"height,omitempty"`
	Hobbies      string `json:"hobbies,omitempty"`
	Artist       string `json:"artist,omitempty"`
	VoiceActor   string `json:"voice_actor,omitempty"`
}

// Raid models

type RaidResponse struct {
	Status int  `json:"status"`
	Data   Raid `json:"data"`
}

type Raid struct {
	Current  *RaidEvent `json:"current,omitempty"`
	Upcoming *RaidEvent `json:"upcoming,omitempty"`
}

type RaidEvent struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartAt   string `json:"start_at,omitempty"`
	EndAt     string `json:"end_at,omitempty"`
	Type      string `json:"type,omitempty"`
	Terrain   string `json:"terrain,omitempty"`
	BossName  string `json:"boss_name,omitempty"`
	BossImage string `json:"boss_image,omitempty"`
}

// Banner models

type BannerResponse struct {
	Status int      `json:"status"`
	Data   []Banner `json:"data"`
}

type Banner struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	StartAt    string   `json:"start_at,omitempty"`
	EndAt      string   `json:"end_at,omitempty"`
	Type       string   `json:"type,omitempty"`
	Characters []string `json:"characters,omitempty"`
	Image      string   `json:"image,omitempty"`
}

// Health / info

type HealthResponse struct {
	Status    string `json:"status"`
	Uptime    int64  `json:"uptime_seconds"`
	Upstream  string `json:"upstream"`
	CacheHits int64  `json:"cache_hits"`
}