package authenticated_origin_pulls

import (
	"fmt"
)

// AuthenticatedOriginPullsArrayResultEnvelope handles the array response from the batch PUT endpoint.
// The API endpoint PUT /zones/{zone_id}/origin_tls_client_auth/hostnames returns an array
// of hostname associations, but Terraform models this as individual resources.
type AuthenticatedOriginPullsArrayResultEnvelope struct {
	Result []AuthenticatedOriginPullsModel `json:"result"`
}

// FindByHostname extracts the matching hostname from the array response.
// Returns an error if the hostname is not found in the response.
func (e *AuthenticatedOriginPullsArrayResultEnvelope) FindByHostname(hostname string) (*AuthenticatedOriginPullsModel, error) {
	for i := range e.Result {
		if e.Result[i].Hostname.ValueString() == hostname {
			return &e.Result[i], nil
		}
	}
	return nil, fmt.Errorf("hostname %q not found in API response", hostname)
}
