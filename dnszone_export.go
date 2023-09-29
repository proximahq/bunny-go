package bunnynet

import (
	"context"
	"fmt"
)

// Update changes the configuration the DNS Zone with the given ID.
// The updated DNS Zone is returned.
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_update
func (s *DNSZoneService) Export(ctx context.Context, id int64) (string, error) {
	path := fmt.Sprintf("dnszone/%d/export", id)

	return resourcePostWithResponseString(
		ctx,
		s.client,
		path,
		nil,
	)
}
