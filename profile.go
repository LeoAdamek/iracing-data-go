package iracing

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type LicenseCategory uint8

const (
	LicenseCategoryUnknown LicenseCategory = iota
	LicenseCategoryOval
	LicenseCategoryRoad
	LicenseCategoryDirtOval
	LicenseCategoryDirtRoad
)

type LicenseGroup uint8

const (
	LicenseGroupInvalid LicenseGroup = iota
	LicenseGroupRookie
	LicenseGroupD
	LicenseGroupC
	LicenseGroupB
	LicenseGroupA
	LicenseGroupPro
	LicenseGroupProWC
)

type Profile struct {
	ID                   uint64    `json:"cust_id"`
	DisplayName          string    `json:"display_name"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	OnCarName            string    `json:"on_car_name"`
	MemberSince          time.Time `json:"member_since"`
	LastSeason           int64     `json:"last_season"`
	Flags                int       `json:"flags"`
	ClubID               int       `json:"club_id"`
	ClubName             string    `json:"club_name"`
	ConnectionType       string    `json:"connection_type"`
	DownloadServer       string    `json:"download_server"`
	LastLogin            time.Time `json:"last_login"`
	ReadCompetitionRules time.Time `json:"read_comp_rules"`
	Licenses             Licenses  `json:"licenses"`
}

type Licenses struct {
	Oval     License `json:"oval"`
	Road     License `json:"road"`
	DirtOval License `json:"dirt_oval"`
	DirtRoad License `json:"dirt_road"`
}

type License struct {
	Category           LicenseCategory `json:"category_id"`
	CategoryName       string          `json:"category"`
	LicenseLevel       uint            `json:"license_level"`
	SafetyRating       float64         `json:"safety_rating"`
	CornersPerIncident float64         `json:"cpi"`
	IRating            uint16          `json:"irating"`
	TTRating           uint16          `json:"tt_rating"`
	MPRRaces           uint            `json:"mpr_num_races"`
	Color              string          `json:"color"`
	Group              LicenseGroup    `json:"group_id"`
	GroupName          string          `json:"group_name"`
}

func (c *Client) GetSelf(ctx context.Context) (*Profile, error) {
	link := CacheLink{}
	profile := &Profile{}

	if err := c.json(ctx, http.MethodGet, Host+"/data/member/info", nil, &link); err != nil {
		return nil, err
	}

	if err := c.json(ctx, http.MethodGet, link.URL, nil, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (c *Client) GetProfiles(ctx context.Context, userIDs []uint64, includeLicenses bool) ([]Profile, error) {
	link := CacheLink{}
	profiles := []Profile{}

	loc, _ := url.Parse(Host + "/data/member/get")

	if includeLicenses {
		loc.Query().Add("include_licenses", "true")
	}

	for _, id := range userIDs {
		loc.Query().Add("cust_ids", strconv.FormatUint(id, 10))
	}

	if err := c.json(ctx, http.MethodGet, loc.String(), nil, &link); err != nil {
		return nil, err
	}

	if err := c.json(ctx, http.MethodGet, link.URL, nil, &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}
