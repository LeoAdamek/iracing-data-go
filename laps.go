package iracing

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type Lap struct {
	GroupID         int64        `json:"group_id"`
	Name            string       `json:"name"`
	CustomerID      int64        `json:"cust_id"`
	DisplayName     string       `json:"display_name"`
	LapNumber       uint16       `json:"lap_number"`
	Flags           int          `json:"flags"`
	Incident        bool         `json:"incident"`
	SessionTime     uint64       `json:"session_time"`
	LapTime         int64        `json:"lap_time"`
	TeamFastestLap  bool         `json:"team_fastest_lap"`
	PersonalBestLap bool         `json:"personal_best_lap"`
	LicenseGroup    LicenseGroup `json:"license_level"`
	CarNumber       string       `json:"car_number"`
	LapPosition     int8         `json:"lap_position"`
	FastestLap      bool         `json:"fastest_lap"`
	AI              bool         `json:"ai"`
	IntervalValue   *int64       `json:"interval"`
	IntervalUnits   *string      `json:"interval_units"`
}

func (l Lap) Interval() *time.Duration {
	if l.IntervalValue == nil || l.IntervalUnits == nil {
		return nil
	}

	interval := time.Duration(*l.IntervalValue) * time.Millisecond
	return &interval
}

func (l Lap) Time() time.Duration {
	return time.Duration(l.LapTime) * time.Millisecond
}

type LapChartData struct {
	Success   bool `json:"success"`
	ChunkInfo struct {
		Size            uint64   `json:"chunk_size"`
		Count           uint64   `json:"num_chunk"`
		Rows            uint64   `json:"rows"`
		BaseDownLoadURL string   `json:"base_download_url"`
		Names           []string `json:"chunk_file_names"`
	} `json:"chunk_info"`
}

func (c *Client) GetLaps(sessionID uint64, sessionNumber int) ([]Lap, error) {
	link := CacheLink{}
	chartData := LapChartData{}
	laps := []Lap{}

	loc := Host + "/data/results/lap_chart_data?subsession_id=" + strconv.FormatUint(sessionID, 10) +
		"&simsession_number=" + strconv.FormatInt(int64(sessionNumber), 10)

	if c.Verbose {
		log.Println(loc)
	}

	if err := c.json(http.MethodGet, loc, nil, &link); err != nil {
		return nil, err
	}

	if c.Verbose {
		log.Println(link.URL)
	}

	if err := c.json(http.MethodGet, link.URL, nil, &chartData); err != nil {
		return nil, err
	}

	for _, name := range chartData.ChunkInfo.Names {
		loc = chartData.ChunkInfo.BaseDownLoadURL + name
		chunk := []Lap{}

		if c.Verbose {
			log.Println(loc)
		}

		if err := c.json(http.MethodGet, loc, nil, &chunk); err != nil {
			return nil, err
		}

		laps = append(laps, chunk...)
	}

	return laps, nil
}
