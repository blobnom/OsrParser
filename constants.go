package OsrParser

// ALL OSU PLAYMODES
const (
	OSU   = 0
	TAIKO = 1
	CTB   = 2
	MANIA = 3
)

// CLICKSTATE
const (
	LEFTCLICK  = 2 << 0
	RIGHTCLICK = 2 << 1
	KEY1       = 2 << 2
	KEY2       = 2 << 3
	SMOKE      = 2 << 4
)

// Keys
var (
	APIKEY = "" // Enter your osu apikey, optional if you want to parse the beatmap
)
