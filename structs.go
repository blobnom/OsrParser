package OsrParser

import (
	"time"
)

// Replay is the Parsed replay.
type Replay struct {
	PlayMode      int8
	OsuVersion    int32
	BeatmapMD5    string
	BeatmapParsed bool
	Beatmap       *BeatmapData
	Username      string
	ReplayMD5     string
	Count300      uint16
	Count100      uint16
	Count50       uint16
	CountGeki     uint16
	CountKatu     uint16
	CountMiss     uint16
	Score         int32
	MaxCombo      uint16
	Fullcombo     bool
	Mods          []string
	LifebarGraph  []LifeBarGraph
	Timestamp     time.Time
	ReplayData    []*ReplayData
	OnlineScoreID int64
}

// ReplayData is the Parsed Compressed Replay data.
type ReplayData struct {
	Time       int64
	MosueX     float32
	MouseY     float32
	KeyPressed *KeyPressed
}

// KeyPressed is the Parsed Compressed KeyPressed.
type KeyPressed struct {
	LeftClick  bool
	RightClick bool
	Key1       bool
	Key2       bool
	Smoke      bool
}

// LifeBarGraph is the Bar under the Score stuff.
type LifeBarGraph struct {
	Time int32
	HP   float32
}

// BeatmapData from osu api
type BeatmapData []struct {
	BeatmapsetID     string `json:"beatmapset_id"`
	BeatmapID        string `json:"beatmap_id"`
	Approved         string `json:"approved"`
	TotalLength      string `json:"total_length"`
	HitLength        string `json:"hit_length"`
	Version          string `json:"version"`
	FileMd5          string `json:"file_md5"`
	DiffSize         string `json:"diff_size"`
	DiffOverall      string `json:"diff_overall"`
	DiffApproach     string `json:"diff_approach"`
	DiffDrain        string `json:"diff_drain"`
	Mode             string `json:"mode"`
	ApprovedDate     string `json:"approved_date"`
	LastUpdate       string `json:"last_update"`
	Artist           string `json:"artist"`
	Title            string `json:"title"`
	Creator          string `json:"creator"`
	CreatorID        string `json:"creator_id"`
	Bpm              string `json:"bpm"`
	Source           string `json:"source"`
	Tags             string `json:"tags"`
	GenreID          string `json:"genre_id"`
	LanguageID       string `json:"language_id"`
	FavouriteCount   string `json:"favourite_count"`
	Playcount        string `json:"playcount"`
	Passcount        string `json:"passcount"`
	MaxCombo         string `json:"max_combo"`
	Difficultyrating string `json:"difficultyrating"`
}
