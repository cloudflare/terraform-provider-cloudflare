package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceCloudflareLogpushOwnershipChallengeModel represents the source
// cloudflare_logpush_ownership_challenge state structure from v4.x SDKv2 provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_logpush_ownership_challenge (unchanged in v5)
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_logpush_ownership_challenge.go
type SourceCloudflareLogpushOwnershipChallengeModel struct {
	AccountID                  types.String `tfsdk:"account_id"`
	ZoneID                     types.String `tfsdk:"zone_id"`
	DestinationConf            types.String `tfsdk:"destination_conf"`
	OwnershipChallengeFilename types.String `tfsdk:"ownership_challenge_filename"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Plugin Framework)
// ============================================================================

// TargetLogpushOwnershipChallengeModel represents the target
// cloudflare_logpush_ownership_challenge state structure (v500).
// This matches the model in the parent package (internal/services/logpush_ownership_challenge/model.go).
//
// Key changes from v4:
//   - ownership_challenge_filename (computed) removed
//   - filename (computed) added - replaces ownership_challenge_filename
//   - message (computed) added - new field
//   - valid (computed bool) added - new field
type TargetLogpushOwnershipChallengeModel struct {
	AccountID       types.String `tfsdk:"account_id"`
	ZoneID          types.String `tfsdk:"zone_id"`
	DestinationConf types.String `tfsdk:"destination_conf"`
	Filename        types.String `tfsdk:"filename"`
	Message         types.String `tfsdk:"message"`
	Valid           types.Bool   `tfsdk:"valid"`
}
