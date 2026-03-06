package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// TransformV4toV500 converts v4 cloudflare_worker_domain state to
// v5 cloudflare_workers_custom_domain state.
func TransformV4toV500(_ context.Context, source SourceV4WorkersCustomDomainModel) (*TargetWorkersCustomDomainModel, diag.Diagnostics) {
	return &TargetWorkersCustomDomainModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Hostname:    source.Hostname,
		Service:     source.Service,
		ZoneID:      source.ZoneID,
		Environment: source.Environment,
		ZoneName:    source.ZoneName,
	}, nil
}
