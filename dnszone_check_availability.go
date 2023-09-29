package bunnynet

import (
	"context"
)

type CheckAvailabilityGetOpts struct {
	Name *string `json:"Name,omitempty"`
}

type Availability struct {
	Available *bool `json:"Available,omitempty"`
}

func (s *DNSZoneService) CheckAvailability(ctx context.Context, opts *CheckAvailabilityGetOpts) (*Availability, error) {
	return resourcePostWithResponse[Availability](ctx, s.client, "/dnszone/checkavailability", opts)
}
