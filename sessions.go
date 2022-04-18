package iracing

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type EventType uint8

const (
	// Race Event
	EventTypeRace EventType = 5
)

type TrackCategory uint8

const (
	TrackCategoryRoad TrackCategory = 2
)

type Session struct {
	SubsessionID       uint64    `json:"subsession_id"`
	SeasonID           uint64    `json:"season_id"`
	SeasonName         string    `json:"season_name"`
	SeasonShortName    string    `json:"season_short_name"`
	SeasonYear         uint16    `json:"season_year"`
	SeasonQuarter      uint8     `json:"season_quarter"`
	SeriesID           uint64    `json:"series_id"`
	SeriesName         string    `json:"series_name"`
	SeriesShortName    string    `json:"series_short_name"`
	SeriesLogo         string    `json:"series_logo"`
	RaceWeek           uint8     `json:"race_week_num"`
	SessionID          uint64    `json:"session_id"`
	LicenseCategoryID  uint8     `json:"license_category_id"`
	Start              time.Time `json:"start_time"`
	End                time.Time `json:"end_time"`
	CornersPerLap      uint8     `json:"corners_per_lap"`
	CautionType        uint8     `json:"caution_type"`
	EventType          EventType `json:"event_type"`
	EventTypeName      string    `json:"event_type_name"`
	DriverChanges      bool      `json:"driver_changes"`
	MinimumTeamDrivers uint8     `json:"min_team_drivers"`
	MaximumTeamDrivers uint8     `json:"max_team_drivers"`
	DriverChangeRule   uint8     `json:"driver_change_rule"`
	PontsType          string    `json:"points_type"`
	StengthOfField     int16     `json:"event_strength_of_field"`
	AverageLap         int64     `json:"event_average_lap"`
	LapsComplete       int       `json:"laps_complete"`
	Cautions           int8      `json:"num_cautions"`
	CautionLaps        int       `json:"num_caution_laps"`
	LeadChanges        int       `json:"num_lead_changes"`
	Official           bool      `json:"official_session"`
	CanProtest         bool      `json:"can_protest"`
	Track              Track     `json:"track"`
}

type Track struct {
	ID           uint64 `json:"track_id"`
	Name         string `json:"track_name"`
	Layout       string `json:"config_name"`
	CategoryID   uint8  `json:"category_id"`
	CategoryName string `json:"category"`
}

type EventLog struct {
}

// Get data for a session
func (c *Client) GetSession(ctx context.Context, id uint64) (*Session, error) {
	link := &CacheLink{}
	session := &Session{}

	sessionUrl := Host + "/data/results/get?subsession_id=" + strconv.FormatUint(id, 10)

	if err := c.json(ctx, http.MethodGet, sessionUrl, nil, link); err != nil {
		return nil, err
	}

	if err := c.json(ctx, http.MethodGet, link.URL, nil, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (c *Client) GetSessionEventLog(ctx context.Context, sessionID uint64, sessionNumber int) (*EventLog, error) {
	link := CacheLink{}
	uri, _ := url.Parse(Host + "/data/results/event_log")

	if err := c.json(ctx, http.MethodGet, uri.String(), nil, &link); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s Session) SeriesLogoURL() string {
	return ImageHost + "/img/logos/series/" + s.SeriesLogo
}
