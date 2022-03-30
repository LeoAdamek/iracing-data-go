package iracing

import (
	"net/http"
	"time"
)

type Car struct {
	ID                      int64     `json:"car_id"`
	AIEnabled               bool      `json:"ai_enabled"`
	AllowNumberColors       bool      `json:"allow_number_colors"`
	AllowNumberFont         bool      `json:"allow_number_font"`
	AllowSponsor1           bool      `json:"allow_sponsor1"`
	AllowSponsor2           bool      `json:"allow_sponsor2"`
	AllowWheelColor         bool      `json:"allow_wheel_color"`
	AwardExcempt            bool      `json:"award_excempt"`
	Make                    string    `json:"car_make"`
	Model                   string    `json:"car_model"`
	Name                    string    `json:"car_name"`
	NameAbbreviated         string    `json:"car_name_abbreviated"`
	Weight                  int       `json:"car_weight"`
	Created                 time.Time `json:"created"`
	FreeWithSubscription    bool      `json:"free_with_subscription"`
	Headlights              bool      `json:"has_headlights"`
	HasMultipleDryTyreTypes bool      `json:"has_multiple_dry_tire_types"`
	Power                   int       `json:"hp"`
	MinPowerAdjust          int       `json:"min_power_adjust_pct"`
	MaxPowerAdjust          int       `json:"max_power_adjust_pct"`
	MaxWeightPenalty        int       `json:"max_weight_penalty_kg"`
	PackageID               int       `json:"package_id"`
	Patterns                int       `json:"patterns"`
	Retired                 bool      `json:"retired"`
	SKU                     int       `json:"sku"`
	CarTypes                []CarType `json:"car_types"`
	Categories              []string  `json:"categories"`
}

type CarType struct {
	CarType string `json:"car_type"`
}

func (c *Client) GetCars() ([]Car, error) {
	cars := []Car{}

	link := &CacheLink{}

	if err := c.json(http.MethodGet, Host+"/data/car/get", nil, link); err != nil {
		return nil, err
	}

	if err := c.json(http.MethodGet, link.URL, nil, &cars); err != nil {
		return nil, err
	}

	return cars, nil
}
