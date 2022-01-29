package OsrParser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/itchio/lzma"
)

// NewReplay returns an Empty Replay
func NewReplay() *Replay {
	return &Replay{}
}

// PrintBeatmap prints informations about the replay's beatmap at index
func (rep *Replay) PrintBeatmap(index int) (err error) {
	if rep.BeatmapParsed {
		if index >= len(*rep.Beatmap) {
			err = errors.New("Index is bigger than BeatmapDatas array")
			return
		}
		fmt.Printf("Beatmap Data at index %d\n Title: %s\n Version: %s\n Artist: %s\n "+
			"Creator: %s\n BPM: %s\n Beatmap Link: https://osu.ppy.sh/beatmapsets/"+
			"%s#osu/%s\n Creator Link: https://osu.ppy.sh/users/%s\n", index,
			(*rep.Beatmap)[index].Title, (*rep.Beatmap)[index].Version,
			(*rep.Beatmap)[index].Artist, (*rep.Beatmap)[index].Creator,
			(*rep.Beatmap)[index].Bpm, (*rep.Beatmap)[index].BeatmapsetID,
			(*rep.Beatmap)[index].BeatmapID, (*rep.Beatmap)[index].CreatorID)
	} else {
		fmt.Println("Beatmap not yet parsed or beatmap not found.")
	}
	return
}

// PrintAllBeatmap prints informations about the replay's beatmaps
func (rep *Replay) PrintAllBeatmap() {
	if !rep.BeatmapParsed {
		fmt.Println("Beatmap not yet parsed or beatmap not found.")
	} else {
		for index := 0; index < len(*rep.Beatmap); index++ {
			rep.PrintBeatmap(index)
		}
	}

}

// PrintReplay informations about the replays
func (rep *Replay) PrintReplay() {
	var playMode string
	var fullCombo string
	var mods string

	switch rep.PlayMode {
	case 0:
		playMode = "Osu!Standard"
	case 1:
		playMode = "Osu!Taiko"
	case 2:
		playMode = "Osu!CTB"
	case 3:
		playMode = "Osu!Mania"
	}

	switch rep.Fullcombo {
	case false:
		fullCombo = "No FC"
	case true:
		fullCombo = "FC"
	}

	if len(rep.Mods) == 0 {
		mods = "None"
	} else {
		for index := 0; index < len(rep.Mods); index++ {
			mods = mods + fmt.Sprintf("%s ", rep.Mods[index])
		}
	}
	if rep.BeatmapParsed {
		fmt.Printf("Replay of %s playing %s played using %s\n Score: %d\n "+
			"300s: %d\n 100s: %d\n 50s: %d\n Gekis: %d\n Katus: %d\n Misses: %d\n "+
			"FC Status: %s\n Mods used: %s\n",
			rep.Username, (*rep.Beatmap)[0].Title, playMode, rep.Score, rep.Count300,
			rep.Count100, rep.Count50, rep.CountGeki, rep.CountKatu, rep.CountMiss,
			fullCombo, mods)
	} else {
		fmt.Printf("Replay of %s played using %s\n Score: %d\n 300s: %d\n "+
			"100s: %d\n 50s: %d\n Gekis: %d\n Katus: %d\n Misses: %d\n "+
			"FC Status: %s\n Mods used: %s\n",
			rep.Username, playMode, rep.Score, rep.Count300, rep.Count100,
			rep.Count50, rep.CountGeki, rep.CountKatu, rep.CountMiss,
			fullCombo, mods)
	}

}

// HasBit checks if bit is set
func hasBit(num uint64, pos uint) (r bool) {
	val := num & (1 << pos)
	r = val > 0
	return
}

// ParseMods parses a uint32 for mods and returns an array of strings
func ParseMods(mods uint32) (r []string, err error) {
	if mods > 536870911 { // 536870911 is all of the mods turned on
		err = errors.New("uint32 exceeds the mod limit")
	} else {
		if mods > 536870911 { // 536870911 is all of the mods turned on
			err = errors.New("uint32 exceeds the mod limit")
		} else {
			if hasBit(uint64(mods), 0) {
				r = append(r, "NF")
			}
			if hasBit(uint64(mods), 1) {
				r = append(r, "EZ")
			}
			if hasBit(uint64(mods), 2) {
				r = append(r, "TD")
			}
			if hasBit(uint64(mods), 3) {
				r = append(r, "HD")
			}
			if hasBit(uint64(mods), 4) {
				r = append(r, "HR")
			}
			if hasBit(uint64(mods), 5) {
				r = append(r, "SD")
			}
			if hasBit(uint64(mods), 6) {
				r = append(r, "DT")
			}
			if hasBit(uint64(mods), 7) {
				r = append(r, "RX")
			}
			if hasBit(uint64(mods), 8) {
				r = append(r, "HT")
			}
			if hasBit(uint64(mods), 9) {
				r = append(r, "NC")
			}
			if hasBit(uint64(mods), 10) {
				r = append(r, "FL")
			}
			if hasBit(uint64(mods), 11) {
				r = append(r, "AP")
			}
			if hasBit(uint64(mods), 12) {
				r = append(r, "SO")
			}
			if hasBit(uint64(mods), 13) {
				r = append(r, "AP")
			}
			if hasBit(uint64(mods), 14) {
				r = append(r, "PF")
			}
			if hasBit(uint64(mods), 15) {
				r = append(r, "4K")
			}
			if hasBit(uint64(mods), 16) {
				r = append(r, "5K")
			}
			if hasBit(uint64(mods), 17) {
				r = append(r, "6K")
			}
			if hasBit(uint64(mods), 18) {
				r = append(r, "7K")
			}
			if hasBit(uint64(mods), 19) {
				r = append(r, "8K")
			}
			if hasBit(uint64(mods), 20) {
				r = append(r, "FI")
			}
			if hasBit(uint64(mods), 21) {
				r = append(r, "Random")
			}
			if hasBit(uint64(mods), 22) {
				r = append(r, "CN")
			}
			if hasBit(uint64(mods), 23) {
				r = append(r, "TP")
			}
			if hasBit(uint64(mods), 24) {
				r = append(r, "Key9")
			}
			if hasBit(uint64(mods), 25) {
				r = append(r, "Coop")
			}
			if hasBit(uint64(mods), 26) {
				r = append(r, "1K")
			}
			if hasBit(uint64(mods), 27) {
				r = append(r, "2K")
			}
			if hasBit(uint64(mods), 28) {
				r = append(r, "3K")
			}
		}
	}
	return
}

// NewBeatmap returns a new pointer to BeatmapData
func NewBeatmap() *BeatmapData {
	return &BeatmapData{}
}

// ParseBeatmap returns parsed BeatmapData
func ParseBeatmap(beatmapMD5 string) (r *BeatmapData, err error) {
	resp, err := http.Get("https://osu.ppy.sh/api/get_beatmaps?k=" + APIKEY +
		"&h=" + beatmapMD5)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r = NewBeatmap()
	json.Unmarshal(body, r)
	return
}

// ParseReplay parses a Replay and returns a *Replay
func ParseReplay(file []byte, parseBeatmap bool) (r *Replay, err error) {
	var LifeBarRaw string
	var ts int64
	var slength int32
	var compressedReplay []byte

	b := bytes.NewBuffer(file)
	r = NewReplay()
	r.PlayMode, err = rInt8(b)
	if err != nil {
		return
	}
	r.OsuVersion, err = rInt32(b)
	if err != nil {
		return
	}
	r.BeatmapMD5, err = rBString(b)
	if err != nil {
		return
	}
	if parseBeatmap {
		if APIKEY == "" {
			fmt.Println("APIKEY is not set, skip fetching beatmap data")
			r.BeatmapParsed = false
		} else {
			fmt.Println("Getting beatmap datas from Osu!Api...")
			beatmap, _ := ParseBeatmap(r.BeatmapMD5)
			if len(*beatmap) == 0 {
				fmt.Println("Beatmap not found.")
			} else {
				r.Beatmap = beatmap
				r.BeatmapParsed = true
			}
		}

	} else {
		r.BeatmapParsed = false
	}
	r.Username, err = rBString(b)
	if err != nil {
		return
	}
	r.ReplayMD5, err = rBString(b)
	if err != nil {
		return
	}
	r.Count300, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Count100, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Count50, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountGeki, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountKatu, err = rUInt16(b)
	if err != nil {
		return
	}
	r.CountMiss, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Score, err = rInt32(b)
	if err != nil {
		return
	}
	r.MaxCombo, err = rUInt16(b)
	if err != nil {
		return
	}
	r.Fullcombo, err = rBool(b)
	if err != nil {
		return
	}
	modsNum, err := rUInt32(b)
	if err != nil {
		return
	}
	r.Mods, err = ParseMods(modsNum)
	if err != nil {
		return
	}
	LifeBarRaw, err = rBString(b)
	if err != nil {
		return
	}
	r.LifebarGraph = parseLifebar(LifeBarRaw)
	ts, err = rInt64(b)
	if err != nil {
		return
	}
	r.Timestamp = timeFromTicks(ts)
	slength, err = rInt32(b)
	if err != nil {
		return
	}
	compressedReplay, err = rSlice(b, slength)
	if err != nil {
		return
	}
	r.ReplayData, err = ParseCompressed(compressedReplay)
	if err != nil {
		return
	}
	return
}

// https://stackoverflow.com/questions/33144967/what-is-the-c-sharp-datetimeoffset-equivalent-in-go/33161703#33161703

func timeFromTicks(ticks int64) time.Time {
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(ticks/10000000+base, ticks%10000000).UTC()
}

func parseLifebar(s string) []LifeBarGraph {
	var o []LifeBarGraph
	s = strings.Trim(s, ",")
	life := strings.Split(s, ",")
	for i := 0; i < len(life); i++ {
		y := strings.Split(life[i], "|")
		if len(y) < 2 {
			continue
		}
		f, err := strconv.ParseFloat(y[1], 32)
		if err != nil {
			continue
		}
		v, _ := strconv.Atoi(y[0])
		o = append(o, LifeBarGraph{Time: int32(v), HP: float32(f)})
	}
	return o
}

// ParseCompressed parses a compressed replay, (ReplayData)
func ParseCompressed(file []byte) (d []*ReplayData, err error) {
	b := bytes.NewBuffer(file)
	r := lzma.NewReader(b)
	defer r.Close()

	x, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	s := strings.Trim(string(x), ",")

	sa := strings.Split(s, ",")

	for i := 0; i < len(sa); i++ {
		rd := sa[i]
		xd := strings.Split(rd, "|")
		if len(xd) < 4 {
			continue
		}
		var Time int
		var MouseX float64
		var MouseY float64
		var KPA int
		Time, err = strconv.Atoi(xd[0])
		if err != nil {
			return
		}
		MouseX, err = strconv.ParseFloat(xd[1], 32)
		if err != nil {
			return
		}
		MouseY, err = strconv.ParseFloat(xd[2], 32)
		if err != nil {
			return
		}
		KPA, err = strconv.Atoi(xd[3])
		if err != nil {
			return
		}
		KP := KeyPressed{
			LeftClick:  KPA&LEFTCLICK > 0,
			RightClick: KPA&RIGHTCLICK > 0,
			Key1:       KPA&KEY1 > 0,
			Key2:       KPA&KEY2 > 0,
			Smoke:      KPA&SMOKE > 0,
		}
		rdata := ReplayData{
			Time:       int64(Time),
			MosueX:     float32(MouseX),
			MouseY:     float32(MouseY),
			KeyPressed: &KP,
		}
		d = append(d, &rdata)
	}

	return
}
