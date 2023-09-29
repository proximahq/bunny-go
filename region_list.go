package bunnynet

import (
	"context"
)

type Region struct {
	ID *int64 `json:"Id,omitempty"`

	Name                *string  `json:"Name,omitempty"`
	PricePerGigabyte    *float64 `json:"PricePerGigabyte,omitempty"`
	RegionCode          *string  `json:"RegionCode,omitempty"`
	ContinentCode       *string  `json:"ContinentCode,omitempty"`
	CountryCode         *string  `json:"CountryCode,omitempty"`
	Latitude            *float64 `json:"Latitude,omitempty"`
	Longitude           *float64 `json:"Longitude,omitempty"`
	AllowLatencyRouting *bool    `json:"AllowLatencyRouting,omitempty"`
}

type Regions []Region

func (s *RegionService) Get(ctx context.Context, id int64) (*Regions, error) {
	return resourceGet[Regions](ctx, s.client, "region", nil)
}
