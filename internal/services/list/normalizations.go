package list

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// normalizeListItems nullifies empty nested objects on list items returned by the
// API. When a list is of kind "ip", the API returns each item with hostname and
// redirect as empty objects ({}) rather than null. Terraform's Set type uses
// structural equality for element identity, so a known object with all-null
// attributes is not equal to a null object. This causes "Provider produced
// inconsistent result after apply" errors because the planned set element (with
// null hostname/redirect) cannot be correlated with the actual set element (with
// known-but-empty hostname/redirect).
//
// This function normalizes the API response so that empty nested objects are
// stored as null in state, matching what the user wrote in their configuration.
func normalizeListItems(ctx context.Context, items []ListItemModel) {
	for i := range items {
		items[i].Hostname = customfield.NullifyEmptyObject(ctx, items[i].Hostname)
		items[i].Redirect = customfield.NullifyEmptyObject(ctx, items[i].Redirect)
	}
}
