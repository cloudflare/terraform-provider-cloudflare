package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareByoIPPrefixModel represents the source cloudflare_byo_ip_prefix state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_byo_ip_prefix.go
type SourceCloudflareByoIPPrefixModel struct {
	ID            types.String `tfsdk:"id"`
	AccountID     types.String `tfsdk:"account_id"`
	PrefixID      types.String `tfsdk:"prefix_id"`
	Description   types.String `tfsdk:"description"`
	Advertisement types.String `tfsdk:"advertisement"`
}

// TargetByoIPPrefixModel represents the target cloudflare_byo_ip_prefix state structure (v500).
// Must match byo_ip_prefix.ByoIPPrefixModel exactly.
type TargetByoIPPrefixModel struct {
	ID                       types.String      `tfsdk:"id"`
	AccountID                types.String      `tfsdk:"account_id"`
	ASN                      types.Int64       `tfsdk:"asn"`
	CIDR                     types.String      `tfsdk:"cidr"`
	LOADocumentID            types.String      `tfsdk:"loa_document_id"`
	DelegateLOACreation      types.Bool        `tfsdk:"delegate_loa_creation"`
	Description              types.String      `tfsdk:"description"`
	Advertised               types.Bool        `tfsdk:"advertised"`
	AdvertisedModifiedAt     timetypes.RFC3339 `tfsdk:"advertised_modified_at"`
	Approved                 types.String      `tfsdk:"approved"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	IrrValidationState       types.String      `tfsdk:"irr_validation_state"`
	ModifiedAt               timetypes.RFC3339 `tfsdk:"modified_at"`
	OnDemandEnabled          types.Bool        `tfsdk:"on_demand_enabled"`
	OnDemandLocked           types.Bool        `tfsdk:"on_demand_locked"`
	OwnershipValidationState types.String      `tfsdk:"ownership_validation_state"`
	OwnershipValidationToken types.String      `tfsdk:"ownership_validation_token"`
	RPKIValidationState      types.String      `tfsdk:"rpki_validation_state"`
}
