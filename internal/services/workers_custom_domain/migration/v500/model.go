package v500

import "github.com/hashicorp/terraform-plugin-framework/types"

// SourceV4WorkersCustomDomainModel represents v4 cloudflare_worker_domain state.
type SourceV4WorkersCustomDomainModel struct {
	ID          types.String `tfsdk:"id"`
	AccountID   types.String `tfsdk:"account_id"`
	Hostname    types.String `tfsdk:"hostname"`
	Service     types.String `tfsdk:"service"`
	ZoneID      types.String `tfsdk:"zone_id"`
	Environment types.String `tfsdk:"environment"`
	ZoneName    types.String `tfsdk:"zone_name"`
	CertId      types.String `tfsdk:"cert_id"`
}

// TargetWorkersCustomDomainModel represents v5 cloudflare_workers_custom_domain state.
type TargetWorkersCustomDomainModel struct {
	ID          types.String `tfsdk:"id"`
	AccountID   types.String `tfsdk:"account_id"`
	Hostname    types.String `tfsdk:"hostname"`
	Service     types.String `tfsdk:"service"`
	ZoneID      types.String `tfsdk:"zone_id"`
	Environment types.String `tfsdk:"environment"`
	ZoneName    types.String `tfsdk:"zone_name"`
	CertId      types.String `tfsdk:"cert_id"`
}
